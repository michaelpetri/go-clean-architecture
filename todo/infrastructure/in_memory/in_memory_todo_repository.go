package in_memory

import (
	"com.michael-petri/todo/domain/model"
	"com.michael-petri/todo/domain/repository"
	"com.michael-petri/todo/domain/value"
	"errors"
)

type InMemoryTodoRepository struct {
	todos *map[uint64]*model.Todo
}

func NewInMemoryTodoRepository(data *map[uint64]*model.Todo) repository.TodoRepository {
	return InMemoryTodoRepository{
		data,
	}
}

func (r InMemoryTodoRepository) Get(id *value.TodoId) (*model.Todo, error) {
	todo := (*r.todos)[id.Value]

	if todo != nil {
		return todo, nil
	}

	return nil, errors.New("failed to get todo")
}

func (r InMemoryTodoRepository) All() ([]*model.Todo, error) {
	todos := make([]*model.Todo, 0, len(*r.todos))

	for _, todo := range *r.todos {
		todos = append(todos, todo)
	}

	return todos, nil
}

func (r InMemoryTodoRepository) Save(todo *model.Todo) (*value.TodoId, error) {
	nextIdValue := len(*r.todos) + 1

	todo.Id = value.NewTodoId(
		uint64(nextIdValue),
	)

	(*r.todos)[todo.Id.Value] = todo

	return todo.Id, nil
}

func (r InMemoryTodoRepository) Delete(id *value.TodoId) error {
	delete(*r.todos, id.Value)

	return nil
}
