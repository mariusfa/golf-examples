package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAppStartup(t *testing.T) {
	router := appSetup()

	r := httptest.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", status)
	}
}

func TestHealthWiredUp(t *testing.T) {
	router := appSetup()

	r := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", status)
	}
}

func TestMetricsWiredUp(t *testing.T) {
	router := appSetup()

	r := httptest.NewRequest("GET", "/metrics", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", status)
	}
}