package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var tmpl *template.Template

type Todo struct {
	Item string
	Done bool
}

type PageData struct {
	Title string
	Todos []Todo
}

func todo(w http.ResponseWriter, r *http.Request) {
	tmpl = template.Must(template.ParseFiles("home.html"))

	data := PageData{
		Title: "Todo list",
		Todos: []Todo{
			{Item: "a", Done: true},
			{Item: "b", Done: false},
			{Item: "c", Done: false},
		},
	}

	tmpl.Execute(w, data)
}

func main() {
	fmt.Printf("Starting server at port 8089\n")
	mux := http.NewServeMux()
	mux.HandleFunc("/todo", todo)
	log.Fatal(http.ListenAndServe(":8099", mux))

}
