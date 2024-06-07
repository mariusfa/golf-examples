package services

import "testing"

func TestGetTodos(t *testing.T) {
	todoRepositoryFake := &todoRepositoryFake{Todos: []Todo{
		{Id: "1", Title: "Buy milk"},
	}}
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

func TestInsertTodo(t *testing.T) {
	todoRepositoryFake := &todoRepositoryFake{Todos: []Todo{}}
	todoService := &TodoService{todoRepositoryFake}
	todo := Todo{Title: "Buy milk"}
	err := todoService.Insert(todo)

	if err != nil {
		t.Fatalf("Insert failed: %v", err)
	}

	insertedTodos := todoRepositoryFake.Todos
	if len(insertedTodos) != 1 {
		t.Fatalf("Expected 1 todo, got %d", len(insertedTodos))
	}
	if insertedTodos[0].Title != "Buy milk" {
		t.Errorf("Expected todo title to be %s, got %s", "Buy milk", insertedTodos[0].Title)
	}
}

type todoRepositoryFake struct {
	Todos []Todo
}

func (t *todoRepositoryFake) GetTodos() ([]Todo, error) {
	return t.Todos, nil
}

func (t *todoRepositoryFake) Insert(todo Todo) error {
	t.Todos = append(t.Todos, todo)
	return nil
}
