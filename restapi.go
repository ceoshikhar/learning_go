package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Article is a struct
type Article struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// Articles  array of type `Article`
var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: /")
}

func createArticle(w http.ResponseWriter, r *http.Request) {
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Articles)
	fmt.Println("Endpoint Hit: /articles")
}

func returnOneArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for i := 0; i < len(Articles); i++ {
		if Articles[i].ID == key {
			json.NewEncoder(w).Encode(Articles[i])
		}
	}
	fmt.Println("Endpoint Hit: articles/" + key)
}

func handleRequests() {
	// Initialize `mux` router
	myRouter := mux.NewRouter().StrictSlash(true)

	// Handle `root` route
	myRouter.HandleFunc("/", homePage)
	// Hangle articles route
	myRouter.HandleFunc("/articles", returnAllArticles)
	// Handle POST to create a new `article`
	myRouter.HandleFunc("/articles", createArticle).Methods("POST")
	// Handle single article route
	myRouter.HandleFunc("/articles/{id}", returnOneArticle)

	// Start the server
	log.Fatal(http.ListenAndServe(":6969", myRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")

	Articles = []Article{
		Article{
			ID:      "1",
			Title:   "Hello",
			Desc:    "Article Description",
			Content: "Article Content",
		},
		Article{
			ID:      "2",
			Title:   "Hello again!",
			Desc:    "Article Description",
			Content: "Article Content",
		},
	}

	handleRequests()
}
