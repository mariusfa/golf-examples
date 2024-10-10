package controllers

import (
	"net/http"

	"github.com/mariusfa/golf/middleware"
)

type HelloController struct{}

func NewHelloController() *HelloController {
	return &HelloController{}
}

func (c *HelloController) RegisterRoutes(router *http.ServeMux) {
	var endpoint http.Handler = http.HandlerFunc(c.PublicRoute)
	endpoint = middleware.PublicRoute(endpoint)
	router.Handle("GET /public", endpoint)
}

func (c *HelloController) PublicRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}
