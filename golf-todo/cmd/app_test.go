package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"todo/internal/services"
)

func TestAppStartup(t *testing.T) {
	todoRepositoryFake := &services.TodoRepositoryFake{Todos: []services.Todo{}}
	router := setupRouter(todoRepositoryFake)

	r := httptest.NewRequest("GET", "/todos", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}
