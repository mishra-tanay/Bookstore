package main

import (
	"api/server/details"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

var books []details.Book

func initfun() {
	file, err := ioutil.ReadFile("../books.json")
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(file, &books)
}
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(books)
}
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for _, item := range books {
		if item.ISBN == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&details.Book{})
}
func UpdateBook(w http.ResponseWriter, r *http.Request) {

}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	var ind int = -1
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for i, item := range books {
		if item.ISBN == params["id"] {
			ind = i
			break
		}
	}
	if ind != -1 {
		tempBook := books[ind]
		books = append(books[:ind], books[ind+1:]...)
		json.NewEncoder(w).Encode(tempBook)
	} else {
		json.NewEncoder(w).Encode("No Book Found")
	}
}
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var book details.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func main() {
	initfun()
	router := mux.NewRouter()
	router.HandleFunc("/books", GetBooks).Methods("GET")
	router.HandleFunc("/getBook/{id}", GetBook).Methods("GET")
	router.HandleFunc("/addBook", CreateBook).Methods("POST")
	router.HandleFunc("/books", UpdateBook).Methods("PUT")
	router.HandleFunc("/deletebyid/{id}", DeleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":3030", router))
}
