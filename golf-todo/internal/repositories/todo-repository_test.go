package repositories

import (
	"testing"
	"todo/internal/services"

	"github.com/mariusfa/golf/database"
)

func TestInsertAndGetTodos(t *testing.T) {
	db, err := database.Setup(&dbConfig)
	if err != nil {
		t.Errorf("Error setting up database: %v", err)
	}
	defer db.Close()

	TodoRepository := NewTodoRepository(db)

	todoToInsert := services.Todo{
		Title: "Test todo",
	}

	if err := TodoRepository.Insert(todoToInsert); err != nil {
		t.Errorf("Error inserting todo: %v", err)
	}

	todos, err := TodoRepository.GetTodos()
	if err != nil {
		t.Errorf("Error getting todos: %v", err)
	}

	if len(todos) != 1 {
		t.Errorf("Expected 1 todo, got %v", len(todos))
	}

	if todos[0].Title != todoToInsert.Title {
		t.Errorf("Expected todo title to be %v, got %v", todoToInsert.Title, todos[0].Title)
	}

	if todos[0].Id == "" {
		t.Errorf("Expected todo ID to be set, got %v", todos[0].Id)
	}
}
