package retrievium

import "cmp"

// TernarySearcher searches a sorted slice using ternary search, dividing the
// search interval into three equal parts each iteration. It implements the
// [Searcher] interface and requires haystack to be sorted in ascending order.
//
// Ternary search runs in O(log₃ n) time and O(1) space.
type TernarySearcher[E cmp.Ordered] struct{}

// Name returns the human-readable name of the algorithm, "Ternary Search".
func (TernarySearcher[E]) Name() string { return "Ternary Search" }

// Search returns the index of target in haystack and true, or 0 and false if
// target is not present. haystack must be sorted in ascending order; the result
// is unspecified otherwise.
func (TernarySearcher[E]) Search(haystack []E, target E) (int, bool) {
	lo, hi := 0, len(haystack)-1
	for lo <= hi {
		third := (hi - lo) / 3
		mid1 := lo + third
		mid2 := hi - third
		switch {
		case haystack[mid1] == target:
			return mid1, true
		case haystack[mid2] == target:
			return mid2, true
		case target < haystack[mid1]:
			hi = mid1 - 1
		case target > haystack[mid2]:
			lo = mid2 + 1
		default:
			lo = mid1 + 1
			hi = mid2 - 1
		}
	}
	return 0, false
}
