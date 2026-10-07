package retrievium_test

import (
	"fmt"
	"testing"

	"github.com/danielriddell21/retrievium/v2"
)

func TestFibonacciSearch(t *testing.T) {
	s := retrievium.FibonacciSearcher[int]{}
	runSearchCases(t, s, withCases(
		searchCase{"two elements first", []int{3, 7}, 3, 0},
		searchCase{"two elements second", []int{3, 7}, 7, 1},
		searchCase{"fibonacci-sized input n=8", []int{1, 2, 3, 5, 8, 13, 21, 34}, 13, 5},
	))
	checkFindsInLargeSorted(t, s)
}

func BenchmarkFibonacciSearch(b *testing.B) {
	benchmarkSearcher(b, retrievium.FibonacciSearcher[int]{})
}

func ExampleFibonacciSearcher() {
	s := retrievium.FibonacciSearcher[int]{}
	fmt.Println(s.Search([]int{1, 3, 5, 7, 9}, 7))
	// Output:
	// 3 true
}
