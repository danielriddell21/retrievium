package retrievium_test

import (
	"fmt"
	"math/rand/v2"
	"slices"

	"github.com/danielriddell21/retrievium/v2"
)

func sortedSlice(n int) []int {
	r := rand.New(rand.NewPCG(42, 0))
	s := make([]int, n)
	for i := range s {
		s[i] = r.IntN(10000) - 5000
	}
	slices.Sort(s)
	return s
}

// Every searcher satisfies Searcher[int], so algorithms are interchangeable.
func Example() {
	searchers := []retrievium.Searcher[int]{
		retrievium.BinarySearcher[int]{},
		retrievium.TernarySearcher[int]{},
		retrievium.FibonacciSearcher[int]{},
	}
	haystack := []int{1, 3, 5, 7, 9, 11}
	target := 7
	for _, s := range searchers {
		idx, ok := s.Search(haystack, target)
		fmt.Printf("%s: index %d found %v\n", s.Name(), idx, ok)
	}
	// Output:
	// Binary Search: index 3 found true
	// Ternary Search: index 3 found true
	// Fibonacci Search: index 3 found true
}
