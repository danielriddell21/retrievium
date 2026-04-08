package retrievium

// FibonacciSearcher implements Fibonacci Search.
// Uses Fibonacci numbers to divide the search range, avoiding division and
// working well when sequential memory access is costly. Requires the input
// slice to be sorted in ascending order.
// Time: O(log n)  Space: O(1)
type FibonacciSearcher struct{}

func (FibonacciSearcher) Name() string { return "Fibonacci Search" }

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
