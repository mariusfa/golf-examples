package repositories

import (
	"database/sql"
	"todo/internal/services"

	"github.com/google/uuid"
)

type TodoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

func (r *TodoRepository) Insert(todo services.Todo) error {
	newId := uuid.New().String()
	_, err := r.db.Exec("INSERT INTO todoschema.todos (id, title) VALUES ($1, $2)", newId, todo.Title)
	return err
}

func (r *TodoRepository) GetTodos() ([]services.Todo, error) {
	rows, err := r.db.Query("SELECT id, title FROM todoschema.todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := []services.Todo{}
	for rows.Next() {
		var todo services.Todo
		if err := rows.Scan(&todo.Id, &todo.Title); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}
