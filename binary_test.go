package retrievium_test

import (
	"fmt"
	"testing"

	"github.com/danielriddell21/retrievium/v2"
)

func TestBinarySearch(t *testing.T) {
	s := retrievium.BinarySearcher[int]{}
	runSearchCases(t, s, sortedCases)
	checkFindsInLargeSorted(t, s)
}

func BenchmarkBinarySearch(b *testing.B) {
	benchmarkSearcher(b, retrievium.BinarySearcher[int]{})
}

func ExampleBinarySearcher() {
	s := retrievium.BinarySearcher[int]{}
	fmt.Println(s.Search([]int{1, 3, 5, 7, 9}, 7))
	// Output:
	// 3 true
}

// ExampleBinarySearcher_Search shows that Search reports the index of a value
// that is present and false for one that is absent.
func ExampleBinarySearcher_Search() {
	s := retrievium.BinarySearcher[int]{}
	fmt.Println(s.Search([]int{1, 3, 5, 7, 9}, 5)) // present
	fmt.Println(s.Search([]int{1, 3, 5, 7, 9}, 4)) // absent
	// Output:
	// 2 true
	// 0 false
}

// ExampleBinarySearcher_Name prints the algorithm's human-readable name.
func ExampleBinarySearcher_Name() {
	fmt.Println(retrievium.BinarySearcher[int]{}.Name())
	// Output: Binary Search
}
