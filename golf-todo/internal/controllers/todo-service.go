package controllers

import "todo/internal/services"

type TodoService interface {
	GetTodos() ([]services.Todo, error)
	Insert(todo services.Todo) error
}
