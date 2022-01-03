package queue

type Item interface{}

// Queue is FIFO (First In First Out)
type Queue interface {
	// Push enqueue
	Push(Item)
	// Pop dequeue
	Pop() Item
	Peek() Item
	Len() int
	IsEmpty() bool
}

type queue struct {
	data []Item
	size int
}

func New() Queue {
	return &queue{}
}

// Push enqueue
func (q *queue) Push(item Item) {
	q.data = append(q.data, item)
	q.size++
}

// Pop dequeue
func (q *queue) Pop() Item {
	var item Item
	item, q.data = q.data[0], q.data[1:]
	q.size--
	return item
}

func (q *queue) Peek() Item {
	return q.data[0]
}

func (q *queue) Len() int {
	return q.size
}

func (q *queue) IsEmpty() bool {
	return q.Len() == 0
}
