package retrievium

import "cmp"

// BinarySearcher searches a sorted slice using binary search, repeatedly
// halving the search interval. It implements the [Searcher] interface and
// requires haystack to be sorted in ascending order.
//
// Binary search runs in O(log n) time and O(1) space.
type BinarySearcher[E cmp.Ordered] struct{}

// Name returns the human-readable name of the algorithm, "Binary Search".
func (BinarySearcher[E]) Name() string { return "Binary Search" }

// Search returns the index of target in haystack and true, or 0 and false if
// target is not present. haystack must be sorted in ascending order; the result
// is unspecified otherwise.
func (BinarySearcher[E]) Search(haystack []E, target E) (int, bool) {
	lo, hi := 0, len(haystack)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		switch {
		case haystack[mid] == target:
			return mid, true
		case haystack[mid] < target:
			lo = mid + 1
		default:
			hi = mid - 1
		}
	}
	return 0, false
}
