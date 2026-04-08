package retrievium

// LinearSearcher implements Linear Search.
// Scans every element in sequence until the target is found or the slice is
// exhausted. Works on unsorted input.
// Time: O(n)  Space: O(1)
type LinearSearcher struct{}

func (LinearSearcher) Name() string { return "Linear Search" }

func (LinearSearcher) Search(haystack []int, target int) int {
	for i, v := range haystack {
		if v == target {
			return i
		}
	}
	return -1
}
