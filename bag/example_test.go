package bag_test

import (
	"fmt"

	"github.com/reiot777/go-algs/bag"
)

func Example_usage() {
	// Create a new bag with some integers in it
	b := bag.New()
	for i := 0; i < 10; i++ {
		b.Add(i)
	}
	b.Add(8)

	// Remove every odd integer
	for i := 0; i < 10; i += 2 {
		b.Remove(i)
	}

	// Print the element count of all numbers
	for i := 0; i < 10; i++ {
		fmt.Printf("#%d: %d\n", i, b.Count(i))
	}

	var sum int
	b.Do(func(item bag.Item) {
		sum += item.(int)
	})

	fmt.Println("Sum:", sum)
}
