package main

import (
	mo "api-skeleton/models"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homepage")
}

func handleRequests() {
	baseRouter := mux.NewRouter().StrictSlash(true)
	baseRouter.HandleFunc("/", homePage)
	baseRouter.HandleFunc("/article/all", mo.ReturnAllArticles)
	baseRouter.HandleFunc("/article/{id}", mo.ReturnSingleArticle)
	log.Fatal(http.ListenAndServe(":10000", baseRouter))
}

func main() {
	handleRequests()
}
