package main

import (
	// "encoding/json"
	// "fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

var books []Book

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("Get all books is called")
}

func getBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Get single book is called")
}

func addBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Add book is called")
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("updateBook is called")
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	log.Println("removeBook is called")
}
