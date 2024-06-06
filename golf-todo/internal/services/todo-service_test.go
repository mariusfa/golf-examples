package services

import "testing"

func TestGetTodos(t *testing.T) {
	todoRepositoryFake := &todoRepositoryFake{}
	todoService := &TodoService{todoRepositoryFake}
	todos, err := todoService.GetTodos()
	if err != nil {
		t.Fatalf("GetTodos failed: %v", err)
	}

	if len(todos) != 1 {
		t.Fatalf("Expected 1 todo, got %d", len(todos))
	}

	if todos[0].Id != "1" {
		t.Errorf("Expected todo ID to be %s, got %s", "1", todos[0].Id)
	}

	if todos[0].Title != "Buy milk" {
		t.Errorf("Expected todo title to be %s, got %s", "Buy milk", todos[0].Title)
	}
}

type todoRepositoryFake struct{}

func (t *todoRepositoryFake) GetTodos() ([]Todo, error) {
	return []Todo{
		{Id: "1", Title: "Buy milk"},
	}, nil
}
