package usecase

import (
	"com.michael-petri/todo/domain/model"
	"com.michael-petri/todo/domain/repository"
)

type ListTodosCase struct {
	todos repository.TodoRepository
}

func NewListTodosCase(todos repository.TodoRepository) *ListTodosCase {
	return &ListTodosCase{
		todos: todos,
	}
}

func (r *ListTodosCase) Invoke() ([]*model.Todo, error) {
	return r.todos.All()
}
