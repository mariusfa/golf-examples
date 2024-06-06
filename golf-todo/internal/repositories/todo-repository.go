package repositories

import (
	"database/sql"
	"todo/internal/services"
)

type TodoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

func (r *TodoRepository) Insert(todo services.Todo) error {
	_, err := r.db.Exec("INSERT INTO todoschema.todos (id, title) VALUES ($1, $2)", todo.Title)
	return err
}
