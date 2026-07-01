// Package retrievium provides a collection of searching algorithm
// implementations behind a single interface.
//
// Every algorithm satisfies the [Searcher] interface, generic over any
// [cmp.Ordered] element type. Its Search method returns the index of the target
// value and a boolean reporting whether it was found. [BinarySearcher],
// [TernarySearcher], [FibonacciSearcher], and [JumpSearcher] require the input
// slice to be sorted in ascending order; [LinearSearcher] works on any input.
//
// All searchers are stateless empty structs. The zero value of each is ready
// to use, and because they hold no state a single value may be shared and used
// concurrently by multiple goroutines.
package retrievium

import "cmp"

// Searcher is the interface implemented by every searching algorithm in this
// package. E is the element type, which may be any [cmp.Ordered] type.
//
// Implementations are stateless: the zero value is ready to use and is safe
// for concurrent use by multiple goroutines.
type Searcher[E cmp.Ordered] interface {
	// Search returns the index of target in haystack and true, or 0 and false
	// if target is not present. [BinarySearcher], [TernarySearcher],
	// [FibonacciSearcher], and [JumpSearcher] require haystack to be sorted in
	// ascending order.
	Search(haystack []E, target E) (int, bool)

	// Name returns the human-readable name of the algorithm.
	Name() string
}
