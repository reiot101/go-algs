package set

// Set data structure.
type Set interface {
	// Add an element into the set.
	Add(interface{})
	// Remove an element from the set.
	Remove(interface{})
	// Len returns the number of elements in the set.
	Len() int
	// Checks whether an element is inside the set.
	Has(interface{}) bool
	// Do executes a func for every element in the set.
	Do(func(interface{}))
	// Reset clears the contents of a set.
	Reset()
}

// set implemented Set interface.
type set struct {
	data map[interface{}]struct{}
}

func New() Set {
	return &set{
		data: make(map[interface{}]struct{}),
	}
}

func (s *set) new() *set {
	return &set{
		data: make(map[interface{}]struct{}),
	}
}

// Add an element into the set.
func (s *set) Add(elem interface{}) {
	s.data[elem] = struct{}{}
}

// Remove an element from the set.
func (s *set) Remove(elem interface{}) {
	delete(s.data, elem)
}

// Len returns the number of elements in the set.
func (s *set) Len() int {
	return len(s.data)
}

// Checks whether an element is inside the set.
func (s *set) Has(elem interface{}) bool {
	_, ok := s.data[elem]
	return ok
}

// Do executes a func for every element in the set.
func (s *set) Do(cb func(interface{})) {
	for k := range s.data {
		cb(k)
	}
}

// Reset clears the contents of a set.
func (s *set) Reset() {
	newSet := &set{
		data: make(map[interface{}]struct{}),
	}
	*s = *newSet
}
