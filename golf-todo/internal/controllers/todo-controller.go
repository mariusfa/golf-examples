package controllers

import (
	"encoding/json"
	"net/http"

	"todo/internal/services"

	accesslog "github.com/mariusfa/golf/logging/access-log"
	"github.com/mariusfa/golf/middleware"
)

type TodoController struct {
	todoService TodoService
}

func NewTodoController(todoService TodoService) *TodoController {
	return &TodoController{todoService}
}

func (t *TodoController) RegisterRoutes(router *http.ServeMux) {
	handler := t.GetTodos()
	handler = middleware.AccessLogMiddleware(handler, &accesslog.AccessLog)
	router.Handle("/todo", handler)
}

func (t *TodoController) GetTodos() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		todos := []services.Todo{
			{Id: "1", Title: "Buy milk"},
		}

		jsonTodos, err := json.Marshal(todos)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonTodos)
	})
}
