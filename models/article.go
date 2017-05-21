package models

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func ReturnAllArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Hello World!", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello Gophers!", Desc: "Article Description", Content: "Article Content"},
	}
	fmt.Println("Endpoint Hit: ReturnAllArticles")

	json.NewEncoder(w).Encode(articles)
}

func ReturnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	fmt.Fprintf(w, "Key "+key)
}
