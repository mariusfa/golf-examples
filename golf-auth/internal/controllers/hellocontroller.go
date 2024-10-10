package controllers

import (
	"auth/internal/repositories"
	"net/http"

	"github.com/mariusfa/golf/middleware"
)

type HelloController struct {
	userRepo *repositories.UserRepository
}

func NewHelloController(userRepo *repositories.UserRepository) *HelloController {
	return &HelloController{userRepo: userRepo}
}

func (c *HelloController) RegisterRoutes(router *http.ServeMux, authParams middleware.AuthParams) {
	router.Handle("GET /public", middleware.PublicRoute(http.HandlerFunc(c.PublicRoute)))
	router.Handle("GET /private", middleware.PrivateRoute(http.HandlerFunc(c.PrivateRoute), authParams))
}

func (c *HelloController) PublicRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

func (c *HelloController) PrivateRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello private world"))
}
