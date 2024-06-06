package services

type TodoRepository interface {
	GetTodos() ([]Todo, error)
}
