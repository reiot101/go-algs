package set_test

import (
	"fmt"

	"github.com/reiot777/go-algs/set"
)

func Example_usage() {
	s := set.New()
	s.Add(1)
	s.Add(2)
	s.Add(3)

	s.Add(4)
	s.Remove(4)

	if !s.Has(4) {
		fmt.Println("wow")
	} else {
		fmt.Println("...")
	}

	var sum int
	s.Do(func(elem interface{}) {
		sum += elem.(int)
	})

	fmt.Println("Sum:", sum)
}
