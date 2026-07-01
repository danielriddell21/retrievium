package retrievium

import "cmp"

// LinearSearcher searches a slice using linear search, scanning every element
// in sequence. It implements the [Searcher] interface and is the only
// algorithm in this package that does not require haystack to be sorted.
//
// The standard library's slices.Index does the same job; LinearSearcher exists
// so linear search sits alongside the other algorithms behind one interface.
//
// Linear search runs in O(n) time and O(1) space.
type LinearSearcher[E cmp.Ordered] struct{}

// Name returns the human-readable name of the algorithm, "Linear Search".
func (LinearSearcher[E]) Name() string { return "Linear Search" }

// Search returns the index of the first occurrence of target in haystack and
// true, or 0 and false if target is not present. haystack need not be sorted.
func (LinearSearcher[E]) Search(haystack []E, target E) (int, bool) {
	for i, v := range haystack {
		if v == target {
			return i, true
		}
	}
	return 0, false
}
