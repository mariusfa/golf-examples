package controllers

import (
	"net/http"

	accesslog "github.com/mariusfa/golf/logging/access-log"
	"github.com/mariusfa/golf/middleware"
)

type Ping struct {}

func NewPing() *Ping {
	return &Ping{}
}

func (p *Ping) RegisterRoutes(router *http.ServeMux) {
	handler := p.GetPong()
	handler = middleware.AccessLogMiddleware(handler, &accesslog.AccessLog)
	router.Handle("/ping", handler)
}

func (p *Ping) GetPong() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
}