package repository

import (
	"com.michael-petri.todo/domain/model"
	"com.michael-petri.todo/domain/value"
)

type TodoRepository interface {
	Get(id *value.TodoId) (*model.Todo, error)
	Save(todo *model.Todo) (*value.TodoId, error)
	All() ([]*model.Todo, error)
	Delete(id *value.TodoId) error
}
