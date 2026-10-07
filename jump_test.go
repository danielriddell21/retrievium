package retrievium_test

import (
	"fmt"
	"testing"

	"github.com/danielriddell21/retrievium/v2"
)

func TestJumpSearch(t *testing.T) {
	s := retrievium.JumpSearcher[int]{}
	runSearchCases(t, s, withCases(
		searchCase{"perfect square length n=9", []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 7, 6},
	))
	checkFindsInLargeSorted(t, s)
}

func BenchmarkJumpSearch(b *testing.B) {
	benchmarkSearcher(b, retrievium.JumpSearcher[int]{})
}

func ExampleJumpSearcher() {
	s := retrievium.JumpSearcher[int]{}
	fmt.Println(s.Search([]int{1, 3, 5, 7, 9}, 7))
	// Output:
	// 3 true
}
