package heap

type Heap interface {
	// Peek return the min/max element.
	Peek() int
	// Pop extract the min/max element.
	Pop() int
	// Push add a new element.
	Push(int)
	// Len return the length of heap.
	Len() int

	siftUp(index int)
	siftDown(index int)
}
