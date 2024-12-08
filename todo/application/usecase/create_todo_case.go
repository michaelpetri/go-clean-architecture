package usecase

import (
	"com.michael-petri/todo/domain/model"
	"com.michael-petri/todo/domain/repository"
	"com.michael-petri/todo/domain/value"
)

type CreateTodoCase struct {
	todos repository.TodoRepository
}

func NewCreateTodoCase(todos repository.TodoRepository) *CreateTodoCase {
	return &CreateTodoCase{
		todos: todos,
	}
}

func (r *CreateTodoCase) Invoke(todo *model.Todo) (*value.TodoId, error) {
	return r.todos.Save(todo)
}
