package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"todo/internal/services"
)

func TestGetTodos(t *testing.T) {
	todoServiceFake := &todoServiceFake{}
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

type todoServiceFake struct{}

func (t *todoServiceFake) GetTodos() ([]services.Todo, error) {
	return []services.Todo{
		{Id: "1", Title: "Buy milk"},
	}, nil
}
