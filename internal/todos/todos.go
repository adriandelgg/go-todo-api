package todos

import (
	"fmt"
	"slices"
)

var nextId uint
var allTodos []Todo

type Todo struct {
	ID          uint   `json:"id"`
	Description string `json:"description"`
	IsCompleted bool   `json:"isCompleted"`
}

func (t *Todo) MarkCompleted() {
	t.IsCompleted = true
}

func (t *Todo) MarkUncompleted() {
	t.IsCompleted = false
}

func Todos() []Todo {
	return allTodos
}

func DeleteTodo(id uint) {
	for i, todo := range allTodos {
		if todo.ID == id {
			allTodos = slices.Delete(allTodos, i, i+1)
			return
		}
	}
}

func GetTodo(id uint) (*Todo, error) {
	index := slices.IndexFunc(allTodos, func(todo Todo) bool {
		return todo.ID == id
	})
	if index == -1 {
		return &Todo{}, fmt.Errorf("todo doesn't exist with given ID of %v", id)
	}
	return &allTodos[index], nil
}

func NewTodo(description string) Todo {
	nextId++
	todo := Todo{
		ID:          nextId,
		Description: description,
		IsCompleted: false,
	}
	allTodos = append(allTodos, todo)
	return todo
}
