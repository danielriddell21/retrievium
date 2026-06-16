package retrievium

// FibonacciSearcher searches a sorted slice using Fibonacci search, dividing
// the search range with Fibonacci numbers instead of halving it. It implements
// the [Searcher] interface and requires haystack to be sorted in ascending
// order.
//
// Avoiding division makes Fibonacci search useful when sequential memory
// access is costly. It runs in O(log n) time and O(1) space.
type FibonacciSearcher struct{}

// Name returns the human-readable name of the algorithm, "Fibonacci Search".
func (FibonacciSearcher) Name() string { return "Fibonacci Search" }

// Search returns the index of target in haystack, or -1 if target is not
// present. haystack must be sorted in ascending order; the result is
// unspecified otherwise.
func (FibonacciSearcher) Search(haystack []int, target int) int {
	n := len(haystack)
	if n == 0 {
		return -1
	}

	// Find the smallest Fibonacci number >= n.
	fibMm2 := 0 // fib(m-2)
	fibMm1 := 1 // fib(m-1)
	fibM := 1   // fib(m)
	for fibM < n {
		fibMm2 = fibMm1
		fibMm1 = fibM
		fibM = fibMm1 + fibMm2
	}

	offset := -1
	for fibM > 1 {
		i := min(offset+fibMm2, n-1)
		switch {
		case haystack[i] < target:
			fibM = fibMm1
			fibMm1 = fibMm2
			fibMm2 = fibM - fibMm1
			offset = i
		case haystack[i] > target:
			fibM = fibMm2
			fibMm1 = fibMm1 - fibMm2
			fibMm2 = fibM - fibMm1
		default:
			return i
		}
	}

	if fibMm1 == 1 && offset+1 < n && haystack[offset+1] == target {
		return offset + 1
	}
	return -1
}
