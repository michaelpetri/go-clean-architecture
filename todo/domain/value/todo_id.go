package value

type TodoId struct {
	Value uint64
}

func NewTodoId(value uint64) *TodoId {
	return &TodoId{
		Value: value,
	}
}
