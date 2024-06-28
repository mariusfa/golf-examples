package controllers

import (
	"encoding/json"
	"net/http"

	"todo/internal/services"

	"github.com/mariusfa/golf/logging/accesslog"
	"github.com/mariusfa/golf/middleware"
)

type TodoController struct {
	todoService TodoService
}

func NewTodoController(todoService TodoService) *TodoController {
	return &TodoController{todoService}
}

func (t *TodoController) RegisterRoutes(router *http.ServeMux) {
	getTodosHandler := t.GetTodos()
	getTodosHandler = middleware.AccessLogMiddleware(getTodosHandler, accesslog.GetLogger())
	router.Handle("GET /todo", getTodosHandler)

	insertTodoHandler := t.InsertTodo()
	insertTodoHandler = middleware.AccessLogMiddleware(insertTodoHandler, accesslog.GetLogger())
	router.Handle("POST /todo", insertTodoHandler)
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

func (t *TodoController) InsertTodo() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var todo services.Todo
		if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := t.todoService.Insert(todo); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	})
}
