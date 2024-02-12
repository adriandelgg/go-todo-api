package main

import (
	"go-todo/internal/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("GET /todos", handlers.GetTodos)
	http.HandleFunc("GET /todos/{id}", handlers.GetTodo)
	http.HandleFunc("POST /newTodo", handlers.CreateTodo)
	http.HandleFunc("DELETE /deleteTodo/{id}", handlers.DeleteTodo)
	http.HandleFunc("PUT /todos/markCompleted/{id}", handlers.MarkCompleted)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
