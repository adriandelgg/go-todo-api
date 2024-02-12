package handlers

import (
	"fmt"
	"go-todo/internal/todos"
	"net/http"
)

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	body, err := parseBodyToJSON[todos.Todo](r.Body)
	if err != nil {
		http.Error(w, "Error parsing JSON request body.", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	todo := todos.NewTodo(body.Description)

	_, err = fmt.Fprintf(w, "Successfully created todo: %+v", todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetTodos(w http.ResponseWriter, _ *http.Request) {
	writeToJSON(w, todos.Todos())
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	id := getIdParam(w, r)

	todo, err := todos.GetTodo(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	writeToJSON(w, todo)
}

func MarkCompleted(w http.ResponseWriter, r *http.Request) {
	id := getIdParam(w, r)
	todo, err := todos.GetTodo(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	todo.MarkCompleted()

	_, err = fmt.Fprint(w, "Successfully completed todo")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id := getIdParam(w, r)
	todos.DeleteTodo(id)

	_, err := fmt.Fprint(w, "Successfully deleted todo")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
