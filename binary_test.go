package retrievium_test

import (
	"fmt"
	"testing"

	"github.com/danielriddell21/retrievium"
)

func TestBinarySearch(t *testing.T) {
	s := retrievium.BinarySearcher[int]{}
	cases := []struct {
		name     string
		haystack []int
		target   int
		want     int
	}{
		{"empty", []int{}, 1, -1},
		{"single found", []int{5}, 5, 0},
		{"single not found", []int{5}, 3, -1},
		{"first element", []int{1, 3, 5, 7, 9}, 1, 0},
		{"last element", []int{1, 3, 5, 7, 9}, 9, 4},
		{"middle element", []int{1, 3, 5, 7, 9}, 5, 2},
		{"not found below range", []int{1, 3, 5, 7, 9}, 0, -1},
		{"not found above range", []int{1, 3, 5, 7, 9}, 10, -1},
		{"not found in range", []int{1, 3, 5, 7, 9}, 4, -1},
		{"with negatives found", []int{-5, -3, 0, 2, 4}, -3, 1},
		{"with negatives not found", []int{-5, -3, 0, 2, 4}, 1, -1},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := s.Search(tc.haystack, tc.target)
			switch {
			case tc.want == -1 && ok:
				t.Errorf("Search(%v, %d) = (%d, true), want not found", tc.haystack, tc.target, got)
			case tc.want != -1 && (!ok || got != tc.want):
				t.Errorf("Search(%v, %d) = (%d, %v), want (%d, true)", tc.haystack, tc.target, got, ok, tc.want)
			}
		})
	}
	t.Run("random n=1000", func(t *testing.T) {
		data := sortedSlice(1000)
		target := data[500]
		got, ok := s.Search(data, target)
		if !ok {
			t.Errorf("Search did not find target %d in slice", target)
		}
		if data[got] != target {
			t.Errorf("Search(%d) = index %d, haystack[%d] = %d", target, got, got, data[got])
		}
	})
}

func BenchmarkBinarySearch(b *testing.B) {
	s := retrievium.BinarySearcher[int]{}
	for _, size := range []int{100, 1000, 10000} {
		data := sortedSlice(size)
		target := data[size/2]
		b.Run(fmt.Sprintf("n=%d", size), func(b *testing.B) {
			for b.Loop() {
				s.Search(data, target)
			}
		})
	}
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
