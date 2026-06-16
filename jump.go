package retrievium

import "math"

// JumpSearcher searches a sorted slice using jump search, advancing in fixed
// blocks of √n elements to find the block containing the target and then
// scanning within it. It implements the [Searcher] interface and requires
// haystack to be sorted in ascending order.
//
// Jump search runs in O(√n) time and O(1) space.
type JumpSearcher struct{}

// Name returns the human-readable name of the algorithm, "Jump Search".
func (JumpSearcher) Name() string { return "Jump Search" }

// Search returns the index of target in haystack, or -1 if target is not
// present. haystack must be sorted in ascending order; the result is
// unspecified otherwise.
func (JumpSearcher) Search(haystack []int, target int) int {
	n := len(haystack)
	if n == 0 {
		return -1
	}

	step := int(math.Sqrt(float64(n)))
	prev := 0

	for haystack[min(step, n)-1] < target {
		prev = step
		step += int(math.Sqrt(float64(n)))
		if prev >= n {
			return -1
		}
	}

	for haystack[prev] < target {
		prev++
		if prev == min(step, n) {
			return -1
		}
	}

	if haystack[prev] == target {
		return prev
	}
	return -1
}
