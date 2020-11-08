package main 

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
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

func testPostArticles(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "test POST endpoint")
}
func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Homepage endpoint hit")
}

func handleRequests(){

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", testPostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main(){
	handleRequests()
}