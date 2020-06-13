package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
)

type Article struct {
    Id      string `json:"id"`
    Title   string `json:"title"`
    Desc    string `json:"desc"`
    Content string `json:"content"`
}

var Articles []Article

func homePage (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: /")
}

func createArticle ( w http.ResponseWriter, r *http.Request ) {
    reqBody, _ := ioutil.ReadAll(r.Body)
    fmt.Fprintf(w, "%+v", string(reqBody));
}

func returnAllArticles ( w http.ResponseWriter, r *http.Request ) {
    json.NewEncoder(w).Encode(Articles)
    fmt.Println("Endpoint Hit: /articles")
}

func returnOneArticle ( w http.ResponseWriter, r *http.Request ) {
    vars := mux.Vars(r)
    key  := vars["id"]
    
    for i := 0; i < len(Articles); i++ {
        if ( Articles[i].Id == key ) {
            json.NewEncoder(w).Encode(Articles[i])
        }
    }
    fmt.Println("Endpoint Hit: articles/" + key)
}


func handleRequests () {
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

func main () {
    fmt.Println("Rest API v2.0 - Mux Routers")

    Articles = []Article {
        Article {
            Id: "1",
            Title: "Hello",
            Desc: "Article Description",
            Content: "Article Content",
        },
        Article {
            Id: "2",
            Title: "Hello again!",
            Desc: "Article Description",
            Content: "Article Content",
        },
    }

    handleRequests()
}
