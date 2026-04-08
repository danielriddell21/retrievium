package retrievium

import "math"

// JumpSearcher implements Jump Search.
// Jumps ahead by √n steps to find the block containing the target, then
// performs a linear scan within that block. Requires the input slice to be
// sorted in ascending order.
// Time: O(√n)  Space: O(1)
type JumpSearcher struct{}

func (JumpSearcher) Name() string { return "Jump Search" }

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
