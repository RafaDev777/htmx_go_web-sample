package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Todo struct {
	Message string
	Id      int
}

func main() {
	fmt.Println("Server is up and Running on Port 8000,")
	fmt.Println("Go to : localhost:8000")
	data := map[string][]Todo{
		"Todos": {
			{Id: 1, Message: "Test - 1"},
			{Id: 2, Message: "Test - 2"},
		},
	}

	todosHandler := func(w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseFiles("index.html"))

		templ.Execute(w, data)
	}

	addTodoHandler := func(w http.ResponseWriter, r *http.Request) {
		message := r.PostFormValue("message")
		templ := template.Must(template.ParseFiles("index.html"))
		todo := Todo{Id: len(data["Todos"]) + 1, Message: message}
		data["Todos"] = append(data["Todos"], todo)

		templ.ExecuteTemplate(w, "todo-list-element", todo)
	}

	http.HandleFunc("/", todosHandler)
	http.HandleFunc("/add-todo", addTodoHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
