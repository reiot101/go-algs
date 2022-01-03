package heap

var _ Heap = &minHeap{}

type minHeap struct {
	data []int
}

func NewMinHeap() Heap {
	return &minHeap{}
}

// Peek return the min/max element.
func (h *minHeap) Peek() int {
	return h.data[0]
}

// Pop extract the min/max element.
func (h *minHeap) Pop() int {
	// Swap the min element with the last one
	h.swap(0, h.Len()-1)

	// Get the min element. O(1)
	last := h.data[h.Len()-1]

	// Evict the min element. O(1)
	h.data = h.data[:h.Len()-1]

	// Maintain heap property. O(logn)
	h.siftDown(0)
	return last
}

// Push add a new element.
func (h *minHeap) Push(elem int) {
	// Add the element to the end of array.
	h.data = append(h.data, elem)
	// Maintain heap property. O(logn)
	h.siftUp(h.Len() - 1)
}

// Len return the length of heap.
func (h *minHeap) Len() int {
	return len(h.data)
}

func (h *minHeap) siftUp(index int) {
	// Stop at root.
	if index == 0 {
		return
	}

	parent := (index - 1) / 2
	// Stop if greater or equal to parent.
	if h.data[index] >= h.data[parent] {
		return
	}

	// Swap the parent.
	h.swap(index, parent)

	// Continue siftUp on parent.
	h.siftUp(parent)
}

func (h *minHeap) siftDown(index int) {
	left, right := index*2+1, index*2+2

	if left >= h.Len() {
		return
	}

	// Get the min child.
	minChild := left
	if right < h.Len() && h.data[right] < h.data[left] {
		minChild = right
	}
	if h.data[index] <= h.data[minChild] {
		return
	}

	// Swap with min child.
	h.swap(index, minChild)

	// Continue siftDown on min child.
	h.siftDown(minChild)
}

func (h *minHeap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}
