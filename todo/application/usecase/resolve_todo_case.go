package usecase

import (
	"com.michael-petri/todo/domain/repository"
	"com.michael-petri/todo/domain/value"
)

type ResolveTodoCase struct {
	todos repository.TodoRepository
}

func NewResolveTodoCase(todos repository.TodoRepository) *ResolveTodoCase {
	return &ResolveTodoCase{
		todos: todos,
	}
}

func (r *ResolveTodoCase) Invoke(id *value.TodoId) error {
	return r.todos.Delete(id)
}
