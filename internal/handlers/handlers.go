package handlers

import (
	"fmt"
	"go-todo/internal/todos"
	"net/http"
)

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	body, err := decodeBodyToJSON[todos.Todo](r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer deferBodyClose(r.Body)

	todo := todos.NewTodo(body.Description)

	if _, err2 := fmt.Fprintf(w, "Successfully created todo: %+v", todo); err2 != nil {
		http.Error(w, err2.Error(), http.StatusInternalServerError)
	}
}

func GetTodos(w http.ResponseWriter, _ *http.Request) {
	if err := writeJSONResponse(w, todos.Todos()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	id, err1 := getIdParam(r)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
		return
	}

	todo, err2 := todos.GetTodo(id)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusNotFound)
		return
	}

	if err3 := writeJSONResponse(w, todo); err3 != nil {
		http.Error(w, err3.Error(), http.StatusInternalServerError)
	}
}

func MarkCompleted(w http.ResponseWriter, r *http.Request) {
	id, err1 := getIdParam(r)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
		return
	}

	todo, err2 := todos.GetTodo(id)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusInternalServerError)
		return
	}
	todo.MarkCompleted()

	if _, err := fmt.Fprint(w, "Successfully completed todo"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err1 := getIdParam(r)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
		return
	}
	todos.DeleteTodo(id)

	if _, err := fmt.Fprint(w, "Successfully deleted todo"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
