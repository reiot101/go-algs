package set

import (
	"math/rand"
	"testing"
)

func TestSet(t *testing.T) {
	// fixtures data
	n := 65536
	data := make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = rand.Int()
	}

	// Fill the set with the data and verify that they're all set
	set := New()
	for _, elem := range data {
		set.Add(elem)
	}
	for _, elem := range data {
		if !set.Has(elem) {
			t.Errorf("failed to locate element in set: %v in %v", elem, set)
		}
	}

	// Remove a few elements and ensure they's out.
	rems := data[:1024]
	for _, elem := range rems {
		set.Remove(elem)
	}
	for _, elem := range rems {
		if set.Has(elem) {
			t.Errorf("element exists after remove: %v in %v", elem, set)
		}
	}

	// Calc the sum of the remainder and verify
	var sum int
	set.Do(func(elem interface{}) {
		sum += elem.(int)
	})

	var want int
	for _, elem := range data {
		want += elem
	}
	for _, elem := range rems {
		want -= elem
	}
	if sum != want {
		t.Errorf("iteration sum mismatch: have %v, want %v", sum, want)
	}

	// Clear the set and ensure nothing's left
	set.Reset()
	for _, elem := range data {
		if set.Has(elem) {
			t.Errorf("element exists after reset: %v in %v", elem, set)
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	data := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		data[i] = rand.Int()
	}

	// Execute the benchmark
	b.ResetTimer()
	set := New()
	for _, elem := range data {
		set.Add(elem)
	}
}

func BenchmarkRemove(b *testing.B) {
	data := rand.Perm(b.N)
	set := New()
	for _, elem := range data {
		set.Add(elem)
	}

	// Execute the benchmark
	rems := rand.Perm(b.N)
	b.ResetTimer()
	for _, elem := range rems {
		set.Remove(elem)
	}
}
