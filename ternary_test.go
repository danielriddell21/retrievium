package retrievium_test

import (
	"fmt"
	"testing"

	"github.com/danielriddell21/retrievium/v2"
)

func TestTernarySearch(t *testing.T) {
	s := retrievium.TernarySearcher[int]{}
	runSearchCases(t, s, withCases(
		searchCase{"two elements first", []int{3, 7}, 3, 0},
		searchCase{"two elements second", []int{3, 7}, 7, 1},
	))
	checkFindsInLargeSorted(t, s)
}

func BenchmarkTernarySearch(b *testing.B) {
	benchmarkSearcher(b, retrievium.TernarySearcher[int]{})
}

func ExampleTernarySearcher() {
	s := retrievium.TernarySearcher[int]{}
	fmt.Println(s.Search([]int{1, 3, 5, 7, 9}, 7))
	// Output:
	// 3 true
}
