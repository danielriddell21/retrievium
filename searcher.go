// Package retrievium provides a collection of searching algorithm implementations.
//
// All implementations satisfy the Searcher interface. Search operates on a
// slice of integers and returns the index of the target value, or -1 if not
// found. BinarySearcher, TernarySearcher, FibonacciSearcher, and JumpSearcher
// require the input slice to be sorted in ascending order. LinearSearcher
// works on any input.
package retrievium

// Searcher is implemented by every searching algorithm in this package.
type Searcher interface {
	// Search returns the index of target in haystack, or -1 if not present.
	// BinarySearcher, TernarySearcher, FibonacciSearcher, and JumpSearcher
	// require haystack to be sorted in ascending order.
	Search(haystack []int, target int) int

	// Name returns the human-readable name of the algorithm.
	Name() string
}
