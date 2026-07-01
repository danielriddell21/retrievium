package retrievium

import (
	"cmp"
	"math"
)

// JumpSearcher searches a sorted slice using jump search, advancing in fixed
// blocks of √n elements to find the block containing the target and then
// scanning within it. It implements the [Searcher] interface and requires
// haystack to be sorted in ascending order.
//
// Jump search runs in O(√n) time and O(1) space.
type JumpSearcher[E cmp.Ordered] struct{}

// Name returns the human-readable name of the algorithm, "Jump Search".
func (JumpSearcher[E]) Name() string { return "Jump Search" }

// Search returns the index of target in haystack and true, or 0 and false if
// target is not present. haystack must be sorted in ascending order; the result
// is unspecified otherwise.
func (JumpSearcher[E]) Search(haystack []E, target E) (int, bool) {
	n := len(haystack)
	if n == 0 {
		return 0, false
	}

	step := int(math.Sqrt(float64(n)))
	prev := 0

	for haystack[min(step, n)-1] < target {
		prev = step
		step += int(math.Sqrt(float64(n)))
		if prev >= n {
			return 0, false
		}
	}

	for haystack[prev] < target {
		prev++
		if prev == min(step, n) {
			return 0, false
		}
	}

	if haystack[prev] == target {
		return prev, true
	}
	return 0, false
}
