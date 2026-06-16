// Package retrievium provides a collection of searching algorithm
// implementations behind a single interface.
//
// Every algorithm satisfies the [Searcher] interface. Its Search method
// operates on a slice of integers and returns the index of the target value,
// or -1 if it is not found. [BinarySearcher], [TernarySearcher],
// [FibonacciSearcher], and [JumpSearcher] require the input slice to be sorted
// in ascending order; [LinearSearcher] works on any input.
//
// All searchers are stateless empty structs. The zero value of each is ready
// to use, and because they hold no state a single value may be shared and used
// concurrently by multiple goroutines.
package retrievium

// Searcher is the interface implemented by every searching algorithm in this
// package.
//
// Implementations are stateless: the zero value is ready to use and is safe
// for concurrent use by multiple goroutines.
type Searcher interface {
	// Search returns the index of target in haystack, or -1 if target is not
	// present. [BinarySearcher], [TernarySearcher], [FibonacciSearcher], and
	// [JumpSearcher] require haystack to be sorted in ascending order.
	Search(haystack []int, target int) int

	// Name returns the human-readable name of the algorithm.
	Name() string
}
