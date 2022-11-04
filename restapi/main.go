package main

//https://www.youtube.com/watch?v=SonwZ6MF5BE
import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

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
	w.Header().Set("Content-type", "application/json")
	id := mux.Vars(r)
	for _, item := range books {
		if item.ID == id["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})

}
func addbook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)

}

func updatebook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/jsom")
	id := mux.Vars(r)
	for index, item := range books {
		if item.ID == id["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = id["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}
func deletebook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/jsom")
	id := mux.Vars(r)
	for index, item := range books {
		if item.ID == id["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	r := mux.NewRouter()

	books = append(books, Book{ID: "1", Isbn: "4444", Title: "Python", Author: &Author{Firstname: "Mehedi", Lastname: "Hossain"}})
	books = append(books, Book{ID: "2", Isbn: "5555", Title: "Go", Author: &Author{Firstname: "Tanvir", Lastname: "Hossain"}})

	r.HandleFunc("/api/books/", getbooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getbook).Methods("GET")
	r.HandleFunc("/api/books/", addbook).Methods("POST")
	r.HandleFunc("/api/books/", updatebook).Methods("PUT")
	r.HandleFunc("/api/books/", deletebook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))

}
