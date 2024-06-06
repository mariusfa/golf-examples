package services

type TodoService struct {
	todoRepository TodoRepository
}

func (t *TodoService) GetTodos() ([]Todo, error) {
	return []Todo{
		{Id: "1", Title: "Buy milk"},
	}, nil
}
