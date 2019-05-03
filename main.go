package main

import (
	"encoding/json"
	// "fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"reflect"
	"strconv"
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

	books = append(books,
		Book{ID: 1, Title: "Go lang start", Author: "Mr. sekiro", Year: "2011"},
		Book{ID: 2, Title: "Go lang second chapter", Author: "Mr. YanRon", Year: "2015"},
		Book{ID: 3, Title: "Go lang third chapter", Author: "Mr. Masaomi", Year: "2014"},
		Book{ID: 4, Title: "Go lang foutth chapter", Author: "Ms. Amane", Year: "2010"},
		Book{ID: 5, Title: "Go lang end", Author: "Mr. Rokaku", Year: "2018"},
	)

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Println(params)
	// display type of id
	log.Println(reflect.TypeOf(params["id"]))

	i, _ := strconv.Atoi(params["id"])
	log.Println(reflect.TypeOf(i))

	for _, book := range books {
		if book.ID == i {
			json.NewEncoder(w).Encode(&book)
		}
	}
}

func addBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	books = append(books, book)
	json.NewEncoder(w).Encode(books)

}

func updateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("updateBook is called")
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	log.Println("removeBook is called")
}
