package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPong(t *testing.T) {
	pingController := NewPing()

	router := http.NewServeMux()
	pingController.RegisterRoutes(router)

	r := httptest.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	expected := "pong"
	if w.Body.String() != expected {
		t.Errorf("Expected body to be %s, got %s", expected, w.Body.String())
	}
}