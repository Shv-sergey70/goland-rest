package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Article struct {
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Test title", Desc: "Test Desctiption", Content: "Hello world!"},
	}
	fmt.Println("All articles endpoint")

	json.NewEncoder(w).Encode(articles)
}

func allPostArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Test title post", Desc: "Test Desctiption post", Content: "Hello world!"},
	}
	fmt.Println("All post articles endpoint")

	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage endpoint")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage)
	router.HandleFunc("/articles", allArticles).Methods("GET")
	router.HandleFunc("/articles", allPostArticles).Methods("POST")

	log.Fatal(http.ListenAndServe(":8083", router))
}

func main() {
	handleRequests()
}