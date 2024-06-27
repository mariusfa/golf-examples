package controllers

import (
	"net/http"

	"github.com/mariusfa/golf/logging/accesslog"
	"github.com/mariusfa/golf/logging/applog"
	"github.com/mariusfa/golf/middleware"
)

type Ping struct{}

func NewPing() *Ping {
	return &Ping{}
}

func (p *Ping) RegisterRoutes(router *http.ServeMux) {
	handler := p.GetPong()
	handler = middleware.AccessLogMiddleware(handler, accesslog.GetLogger())
	router.Handle("/ping", handler)
}

func (p *Ping) GetPong() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		applog.Info("Ping from app log")
		w.Write([]byte("pong"))
	})
}
