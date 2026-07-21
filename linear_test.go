package retrievium_test

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/danielriddell21/retrievium/v2"
)

func TestLinearSearch(t *testing.T) {
	s := retrievium.LinearSearcher[int]{}
	cases := []struct {
		name     string
		haystack []int
		target   int
		want     int
	}{
		{"empty", []int{}, 1, -1},
		{"single found", []int{5}, 5, 0},
		{"single not found", []int{5}, 3, -1},
		{"first element", []int{9, 3, 7, 1, 5}, 9, 0},
		{"last element", []int{9, 3, 7, 1, 5}, 5, 4},
		{"middle element", []int{9, 3, 7, 1, 5}, 7, 2},
		{"not found", []int{9, 3, 7, 1, 5}, 4, -1},
		{"with negatives found", []int{0, -3, 4, -1, 2}, -3, 1},
		{"with negatives not found", []int{0, -3, 4, -1, 2}, 99, -1},
		{"unsorted input", []int{5, 2, 8, 1, 9, 3}, 8, 2},
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
	t.Run("random n=1000 unsorted", func(t *testing.T) {
		r := rand.New(rand.NewPCG(99, 0))
		data := make([]int, 1000)
		for i := range data {
			data[i] = r.IntN(2000) - 1000
		}
		target := data[r.IntN(len(data))]
		got, ok := s.Search(data, target)
		if !ok {
			t.Errorf("Search did not find target %d in slice", target)
		}
		if data[got] != target {
			t.Errorf("Search(%d) = index %d, haystack[%d] = %d", target, got, got, data[got])
		}
	})
}

func BenchmarkLinearSearch(b *testing.B) {
	s := retrievium.LinearSearcher[int]{}
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

func ExampleLinearSearcher() {
	s := retrievium.LinearSearcher[int]{}
	fmt.Println(s.Search([]int{9, 3, 7, 1, 5}, 7))
	// Output:
	// 2 true
}
