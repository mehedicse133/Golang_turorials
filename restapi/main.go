package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string  `json: "id"`
	Isbn   string  `json: "id"`
	Title  string  `json: "id"`
	Author *Author `json: "author"`
}

type Author struct {
	Firstname string `json: "firstname`
	Lastname  string `json: "lastname`
}

var books []Book

func getbooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getbook(w http.ResponseWriter, r *http.Request) {

}
func addbook(w http.ResponseWriter, r *http.Request) {

}
func updatebook(w http.ResponseWriter, r *http.Request) {

}
func deletebook(w http.ResponseWriter, r *http.Request) {

}

func main() {
	r := mux.NewRouter()

	books = append(books, Book{ID: "1", Isbn: "4444", Title: "Python", Author: &Author{Firstname: "Mehedi", Lastname: "Hossain"}})
	books = append(books, Book{ID: "2", Isbn: "5555", Title: "Go", Author: &Author{Firstname: "Tanvir", Lastname: "Hossain"}})

	r.HandleFunc("/api/books/", getbooks).Methods("GET")
	r.HandleFunc("/api/books/", getbook).Methods("GET")
	r.HandleFunc("/api/books/", addbook).Methods("POST")
	r.HandleFunc("/api/books/", updatebook).Methods("PUT")
	r.HandleFunc("/api/books/", deletebook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))

}
