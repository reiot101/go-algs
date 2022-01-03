package bag

// Item is element of bag.
type Item interface{}

// Bag data structure (multiset).
type Bag interface {
	// Add an element into the bag.
	Add(Item)
	// Remove an element from the bag.
	Remove(Item)
	// Len returns the total number of elements in the bag.
	Len() int
	// Count the number of occurances of an element in the bag.
	Count(Item) int
	// Do execute a func for every element in bag.
	Do(func(Item))
	// Reset clear the contents of a bag.
	Reset()
}

// bag implemented Bag interface.
type bag struct {
	data map[Item]int
	size int
}

func New() Bag {
	return &bag{data: make(map[Item]int), size: 0}
}

// Add an element into the bag.
func (b *bag) Add(item Item) {
	b.data[item]++
	b.size++
}

// Remove an element from the bag.
func (b *bag) Remove(item Item) {
	old, ok := b.data[item]
	if !ok {
		return
	}

	if old > 1 {
		b.data[item] = old - 1
	} else {
		delete(b.data, item)
	}
	b.size--
}

// Len returns the total number of elements in the bag.
func (b *bag) Len() int {
	return b.size
}

// Count the number of occurances of an element in the bag.
func (b *bag) Count(item Item) int {
	return b.data[item]
}

// Do execute a func for every element in bag.
func (b *bag) Do(cb func(Item)) {
	for k, v := range b.data {
		for v > 0 {
			cb(k)
			v--
		}
	}
}

// Reset clear the contents of a bag.
func (b *bag) Reset() {
	newBag := &bag{data: make(map[Item]int), size: 0}
	*b = *newBag
}
