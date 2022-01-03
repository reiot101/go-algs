package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	stack := New()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	assert.Equal(t, stack.Len(), 3)

	assert.Equal(t, stack.Peek(), 3)

	assert.Equal(t, stack.Pop(), 3)
	assert.Equal(t, stack.Pop(), 2)
	assert.Equal(t, stack.Pop(), 1)

	assert.Equal(t, stack.IsEmpty(), true)
}
