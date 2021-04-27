package main

import (
	"encoding/json"
	"fmt"
	"goapi/models"
	"log"
	"net/http"
)

var Articles []models.Article

func main() {
	Articles = []models.Article{
		{Title: "Computers and principles", Content: "", Desc: ""},
		{Title: "Chemical Reactions", Content: "", Desc: ""},
	}
	handleRequests()
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "returning All articles")
	json.NewEncoder(w).Encode(Articles)
}
