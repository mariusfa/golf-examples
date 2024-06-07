package services

type TodoRepositoryFake struct {
	Todos []Todo
}

func (t *TodoRepositoryFake) GetTodos() ([]Todo, error) {
	return t.Todos, nil
}

func (t *TodoRepositoryFake) Insert(todo Todo) error {
	t.Todos = append(t.Todos, todo)
	return nil
}
