package services

type TodoRepository interface {
	GetTodos() ([]Todo, error)
	Insert(todo Todo) error
}
