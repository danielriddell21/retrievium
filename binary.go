package retrievium

// BinarySearcher implements Binary Search.
// Repeatedly halves the search interval, discarding the half that cannot
// contain the target. Requires the input slice to be sorted in ascending order.
// Time: O(log n)  Space: O(1)
type BinarySearcher struct{}

func (BinarySearcher) Name() string { return "Binary Search" }

func (BinarySearcher) Search(haystack []int, target int) int {
	lo, hi := 0, len(haystack)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		switch {
		case haystack[mid] == target:
			return mid
		case haystack[mid] < target:
			lo = mid + 1
		default:
			hi = mid - 1
		}
	}
	return -1
}
