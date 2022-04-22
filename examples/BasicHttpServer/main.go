package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Article struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Published bool   `json:"published"`
}

var articles []Article

func main() {
	// populate the articles variable
	articles = []Article{
		Article{Id: 1, Title: "Hello", Published: true},
		Article{Id: 2, Title: "Hi", Published: false},
	}

	handleRequests()
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/status", statusHandler).Methods("GET")
	router.HandleFunc("/articles", postArticleHandler).Methods("POST")
	router.HandleFunc("/articles", getArticlesHandler).Methods("GET")
	router.HandleFunc("/articles/{id}", deleteArticleHandler).Methods("DELETE")
	router.HandleFunc("/articles/{id}", getArticleHandler).Methods("GET")

	port := 8080
	log.Printf("Server running on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), router))
}

func getArticleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET on /articles/{id}")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// loop over all articles
	for _, article := range articles {
		if article.Id == id {
			json.NewEncoder(w).Encode(article)
			break
		}
	}
}

func postArticleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST on /articles")
	// get the request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	// unmarshal the body into an article struct
	json.Unmarshal(reqBody, &article)
	//append the new article to the slice
	articles = append(articles, article)
	json.NewEncoder(w).Encode(article)
}

func getArticlesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET on /articles")
	json.NewEncoder(w).Encode(articles)
}

func deleteArticleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DELETE on /articles/{id}")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for index, article := range articles {
		if article.Id == id {
			articles = append(articles[:index], articles[index+1:]...)
			break
		}
	}
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET on /status")
	fmt.Fprint(w, "Running")
}
