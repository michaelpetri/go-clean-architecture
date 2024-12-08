package in_memory

import (
	"com.michael-petri/todo/domain/model"
	"com.michael-petri/todo/domain/value"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSaveNewTodo(t *testing.T) {
	inMemoryState := make(map[int64]*model.Todo)
	todos := NewInMemoryTodoRepository(&inMemoryState)

	id, err := todos.Save(
		model.NewTodo("Implement database repository"),
	)

	if err != nil {
		t.Failed()
	}

	_, actual := first(inMemoryState)

	assert.Len(t, inMemoryState, 1)
	assert.Equal(t, id, actual.Id)
	assert.Equal(t, "Implement database repository", actual.Description)
}

func TestUpdateExistingTodo(t *testing.T) {
	inMemoryState := make(map[int64]*model.Todo)

	preExisting := model.NewTodo("Implement repository")
	preExisting.Id = &value.TodoId{Value: 1337}

	inMemoryState[preExisting.Id.Value] = preExisting

	todos := NewInMemoryTodoRepository(&inMemoryState)

	preExisting.Description = "Implement database repository"
	_, err := todos.Save(preExisting)

	if err != nil {
		t.Failed()
	}

	_, actual := first(inMemoryState)

	assert.Equal(t, preExisting.Id, actual.Id)
	assert.Equal(t, "Implement database repository", actual.Description)
}

func TestGetExisting(t *testing.T) {
	inMemoryState := make(map[int64]*model.Todo)

	preExisting := model.NewTodo("Implement database repository")
	preExisting.Id = &value.TodoId{Value: 1337}

	inMemoryState[preExisting.Id.Value] = preExisting

	todos := NewInMemoryTodoRepository(&inMemoryState)

	actual, err := todos.Get(preExisting.Id)

	if err != nil {
		t.Failed()
	}

	assert.Equal(t, preExisting.Id, actual.Id)
	assert.Equal(t, "Implement database repository", actual.Description)
}

func TestAll(t *testing.T) {
	inMemoryState := make(map[int64]*model.Todo)

	inMemoryState[0] = model.NewTodo("Implement database repository")
	inMemoryState[1] = model.NewTodo("Write more tests")

	todos := NewInMemoryTodoRepository(&inMemoryState)

	actual, err := todos.All()

	if err != nil {
		t.Failed()
	}

	assert.Equal(t, inMemoryState[0], actual[0])
	assert.Equal(t, inMemoryState[1], actual[1])
}

func TestResolve(t *testing.T) {
	inMemoryState := make(map[int64]*model.Todo)

	preExisting := model.NewTodo("Implement database repository")
	preExisting.Id = &value.TodoId{Value: 1337}

	inMemoryState[preExisting.Id.Value] = preExisting

	todos := NewInMemoryTodoRepository(&inMemoryState)

	err := todos.Delete(preExisting.Id)

	if err != nil {
		t.Failed()
	}

	assert.Len(t, inMemoryState, 0)
}

func first[K comparable, V any](it map[K]V) (K, V) {
	var firstKey K
	var firstValue V
	for currentKey, currentValue := range it {
		firstKey = currentKey
		firstValue = currentValue
		break
	}

	return firstKey, firstValue
}
