package retrievium_test

import (
	"fmt"
	"math/rand/v2"
	"sort"

	"github.com/danielriddell21/retrievium"
)

func sortedSlice(n int) []int {
	r := rand.New(rand.NewPCG(42, 0))
	s := make([]int, n)
	for i := range s {
		s[i] = r.IntN(10000) - 5000
	}
	sort.Ints(s)
	return s
}

// Every searcher satisfies the Searcher interface, so algorithms are interchangeable.
func Example() {
	searchers := []retrievium.Searcher{
		retrievium.BinarySearcher{},
		retrievium.TernarySearcher{},
		retrievium.FibonacciSearcher{},
	}
	haystack := []int{1, 3, 5, 7, 9, 11}
	target := 7
	for _, s := range searchers {
		fmt.Printf("%s: index %d\n", s.Name(), s.Search(haystack, target))
	}
	// Output:
	// Binary Search: index 3
	// Ternary Search: index 3
	// Fibonacci Search: index 3
}
