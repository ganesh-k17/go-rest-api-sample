package main

import (
	"encoding/json"
	"fmt"
	"goapi/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var Articles []models.Article

func main() {
	Articles = []models.Article{
		{Id: "1", Title: "Computers and principles", Content: "", Desc: ""},
		{Id: "2", Title: "Chemical Reactions", Content: "", Desc: ""},
	}
	handleRequests()
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/article/{id}", getArticle)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("returning All articles")
	json.NewEncoder(w).Encode(Articles)
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for i := 0; i < len(Articles); i++ {
		art := Articles[i]
		if art.Id == key {
			json.NewEncoder(w).Encode(art)
			break
		}
	}

	// using for range

	// for _, article := range Articles {
	// 	if article.Id == key {
	// 		json.NewEncoder(w).Encode(article)
	// 	}
	// }
}
