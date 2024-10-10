package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPublicRoute(t *testing.T) {
	helloController := NewHelloController()
	router := http.NewServeMux()

	helloController.RegisterRoutes(router)

	req := httptest.NewRequest("GET", "/public", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	exptectedBody := "Hello world"
	if w.Body.String() != exptectedBody {
		t.Errorf("expected body %s, got %s", exptectedBody, w.Body.String())
	}
}
