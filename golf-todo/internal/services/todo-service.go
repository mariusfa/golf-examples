package services

type TodoService struct {
	todoRepository TodoRepository
}

func NewTodoService(todoRepository TodoRepository) *TodoService {
	return &TodoService{todoRepository}
}

func (t *TodoService) GetTodos() ([]Todo, error) {
	return t.todoRepository.GetTodos()
}

func (t *TodoService) Insert(todo Todo) error {
	return t.todoRepository.Insert(todo)
}
