package main 

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type Article struct {
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request){
	articles := Articles{
		Article{Title:"Test Title", Desc:"Test description", Content:"test content"},
	}

	fmt.Println("Endpoint hit: all articles endpoint")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Homepage endpoint hit")
}

func handleRequests(){
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main(){
	handleRequests()
}