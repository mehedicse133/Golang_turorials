package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to home page")
}

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to contact page")
}

func main() {
	fmt.Printf("Starting server at port 8088\n")

	http.HandleFunc("/", hello)
	http.HandleFunc("/contact", contact)

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "welcome to about page")
	})

	if err := http.ListenAndServe(":8088", nil); err != nil {
		log.Fatal(err)
	}
}
