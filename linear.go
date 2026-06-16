package retrievium

// LinearSearcher searches a slice using linear search, scanning every element
// in sequence. It implements the [Searcher] interface and is the only
// algorithm in this package that does not require haystack to be sorted.
//
// Linear search runs in O(n) time and O(1) space.
type LinearSearcher struct{}

// Name returns the human-readable name of the algorithm, "Linear Search".
func (LinearSearcher) Name() string { return "Linear Search" }

// Search returns the index of the first occurrence of target in haystack, or
// -1 if target is not present. haystack need not be sorted.
func (LinearSearcher) Search(haystack []int, target int) int {
	for i, v := range haystack {
		if v == target {
			return i
		}
	}
	return -1
}
