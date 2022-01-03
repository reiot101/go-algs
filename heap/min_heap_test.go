package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinHeap(t *testing.T) {
	data := []int{5, 1, 3, 5, 3, 4, 3, 7}
	heap := NewMinHeap()

	for _, elem := range data {
		heap.Push(elem)
	}

	var output []int
	for heap.Len() != 0 {
		output = append(output, heap.Pop())
	}

	want := []int{1, 3, 3, 3, 4, 5, 5, 7}
	assert.Equal(t, output, want)
}
