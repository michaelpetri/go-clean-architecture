package model

import "com.michael-petri.todo/domain/value"

type Todo struct {
	Id          *value.TodoId
	Description string
}

// NewTodo creates a new Todo with the given description.
func NewTodo(
	description string,
) *Todo {
	return &Todo{
		Description: description,
	}
}
