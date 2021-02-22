// To test the Endpoints, send request to the localhost:8081
// with the respective HTTP verb and this data in the body as RAW
// {
//     "Id": "27",
//     "Title": "Sonic",
//     "desc": "The Hedgehog",
//     "content": "my articles content"
// }

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	ID      string `json:ID`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

var articlesArray Articles

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	fmt.Println(vars)
	key := vars["id"]

	// Loop over all of our Articles
	// if the article.ID equals the key we pass in
	// return the article encoded as JSON
	// fmt.Fprintf(w, "Key: "+key)

	for _, article := range articlesArray {
		if article.ID == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func allArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articlesArray)
}

func testPostArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Enpoint hit with verb POST")
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Article struct
	// append this to our Articles array.
	reqBody, _ := ioutil.ReadAll(r.Body)

	// Shows the request body of the request
	// fmt.Fprintf(w, "%+v", string(reqBody))

	var article Article

	// Deserialize the JSON into a go compatible structure
	json.Unmarshal(reqBody, &article)

	// Appends the new article saved in memory, to the already created slice
	articlesArray = append(articlesArray, article)

	// Prints in STOUT the new article, sent by the POST request in the body
	json.NewEncoder(w).Encode(article)
	fmt.Println(articlesArray)

}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, article := range articlesArray {
		if article.ID == id {
			articlesArray = append(articlesArray[:index], articlesArray[index+1:]...)
		}
	}
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	// fetchs the ID from the var that is going to be updated
	vars := mux.Vars(r)
	id := vars["id"]

	// get the body of our POST request
	// unmarshal this into a new Article struct
	// append this to our Articles array.
	reqBody, _ := ioutil.ReadAll(r.Body)

	for index, article := range articlesArray {
		if article.ID == id {
			json.Unmarshal(reqBody, &article)
			articlesArray[index] = article
			fmt.Println(article)
			fmt.Println(articlesArray)
		}
	}

}

func handleRequest() {
	// Mux can handle verbs
	myRouter := mux.NewRouter() // .StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/articles", allArticles)
	myRouter.HandleFunc("/articles", testPostArticles).Methods("POST")
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle).Methods("GET")
	myRouter.HandleFunc("/article/{id}", updateArticle).Methods("PUT")
	http.Handle("/", myRouter)
	log.Fatal(http.ListenAndServe(":8081", nil))

	// Standar pkg server
	// http.HandleFunc("/", homePage)
	// http.HandleFunc("/articles", allArticles)
}

func main() {
	articlesArray = Articles{
		Article{ID: "1", Title: "Test title", Desc: "Test Description", Content: "Hello World"},
		Article{ID: "2", Title: "Test title 2", Desc: "Test Description 2", Content: "Hello World 2"},
	}
	handleRequest()
}
