package retrievium

// TernarySearcher implements Ternary Search.
// Divides the search interval into three equal parts each iteration, discarding
// the two-thirds that cannot contain the target. Requires the input slice to be
// sorted in ascending order.
// Time: O(log₃ n)  Space: O(1)
type TernarySearcher struct{}

func (TernarySearcher) Name() string { return "Ternary Search" }

func (TernarySearcher) Search(haystack []int, target int) int {
	lo, hi := 0, len(haystack)-1
	for lo <= hi {
		third := (hi - lo) / 3
		mid1 := lo + third
		mid2 := hi - third
		switch {
		case haystack[mid1] == target:
			return mid1
		case haystack[mid2] == target:
			return mid2
		case target < haystack[mid1]:
			hi = mid1 - 1
		case target > haystack[mid2]:
			lo = mid2 + 1
		default:
			lo = mid1 + 1
			hi = mid2 - 1
		}
	}
	return -1
}
