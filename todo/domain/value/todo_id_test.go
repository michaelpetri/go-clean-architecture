package value

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPositiveValue(t *testing.T) {
	id := NewTodoId(1)

	assert.Equal(t, uint64(1), id.Value)
}
