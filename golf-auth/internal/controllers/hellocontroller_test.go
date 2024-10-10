package controllers

import (
	"auth/internal/repositories"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mariusfa/golf/auth"
	"github.com/mariusfa/golf/logging/middlewarelog"
	"github.com/mariusfa/golf/middleware"
)

func TestPublicRoute(t *testing.T) {
	userRepo := repositories.NewUserRepository()
	authParams := middleware.NewAuthParams("secret", userRepo, middlewarelog.GetLogger())
	helloController := NewHelloController(userRepo)
	router := http.NewServeMux()

	helloController.RegisterRoutes(router, authParams)

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

func TestPrivateRoute(t *testing.T) {
	userRepo := repositories.NewUserRepository()
	adminUser := middleware.User{Id: "123", Name: "admin"}
	userRepo.AddUser(adminUser)

	secret := "secret"
	authParams := middleware.NewAuthParams(secret, userRepo, middlewarelog.GetLogger())

	helloController := NewHelloController(userRepo)
	router := http.NewServeMux()
	helloController.RegisterRoutes(router, authParams)

	token, err := auth.CreateToken(adminUser.Id, secret, nil)
	if err != nil {
		t.Errorf("error creating token %s", err)
	}

	req := httptest.NewRequest("GET", "/private", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
	exptectedBody := "Hello private world"
	if w.Body.String() != exptectedBody {
		t.Errorf("expected body %s, got %s", exptectedBody, w.Body.String())
	}
}
