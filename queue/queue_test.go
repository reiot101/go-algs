package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	queue := New()

	queue.Push(1)
	queue.Push(2)
	queue.Push(3)

	assert.Equal(t, queue.Len(), 3)
	assert.Equal(t, queue.Peek(), 1)

	assert.Equal(t, queue.Pop(), 1)
	assert.Equal(t, queue.Pop(), 2)
	assert.Equal(t, queue.Pop(), 3)

	assert.Equal(t, queue.IsEmpty(), true)
}
