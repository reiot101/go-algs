package stack

type Item interface{}

// Stack is LIFO (Last In First Out) data structure
type Stack interface {
	Push(Item)
	Pop() Item
	IsEmpty() bool
	Peek() Item
	Len() int
}

// stack implemented Stack interface (with array)
type stack struct {
	data []Item
	size int
}

func New() Stack {
	return &stack{}
}

func (s *stack) Push(item Item) {
	s.data = append([]Item{item}, s.data...)
	s.size++
}

func (s *stack) Pop() Item {
	var item Item

	item, s.data = s.data[0], s.data[1:]
	s.size--

	return item
}

func (s *stack) IsEmpty() bool {
	return s.size == 0
}

func (s *stack) Peek() Item {
	return s.data[0]
}

func (s *stack) Len() int {
	return s.size
}
