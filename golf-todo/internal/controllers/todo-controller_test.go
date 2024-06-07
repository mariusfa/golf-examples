package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"todo/internal/services"

	"github.com/google/uuid"
)

func TestGetTodos(t *testing.T) {
	todoServiceFake := &todoServiceFake{Todos: []services.Todo{
		{Id: "1", Title: "Buy milk"},
	}}
	todoController := NewTodoController(todoServiceFake)

	router := http.NewServeMux()
	todoController.RegisterRoutes(router)

	r := httptest.NewRequest("GET", "/todo", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	expected := `[{"id":"1","title":"Buy milk"}]`
	if w.Body.String() != expected {
		t.Errorf("Expected body to be %s, got %s", expected, w.Body.String())
	}
}

func TestInsertTodos(t *testing.T) {
	todoServiceFake := &todoServiceFake{Todos: []services.Todo{}}
	todoController := NewTodoController(todoServiceFake)

	router := http.NewServeMux()
	todoController.RegisterRoutes(router)

	todoToInsert := `{"title":"Buy milk"}`
	r := httptest.NewRequest("POST", "/todo", strings.NewReader(todoToInsert))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, w.Code)
	}
	insertedTodos := todoServiceFake.Todos
	if len(insertedTodos) != 1 {
		t.Errorf("Expected 1 todo, got %d", len(insertedTodos))
	}
	if insertedTodos[0].Id == "" {
		t.Errorf("Expected todo ID to be set, got %s", insertedTodos[0].Id)
	}
}

type todoServiceFake struct {
	Todos []services.Todo
}

func (t *todoServiceFake) GetTodos() ([]services.Todo, error) {
	return t.Todos, nil
}

func (t *todoServiceFake) Insert(todo services.Todo) error {
	id := uuid.New().String()
	todo.Id = id
	t.Todos = append(t.Todos, todo)
	return nil
}
