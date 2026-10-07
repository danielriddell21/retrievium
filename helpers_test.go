package retrievium_test

import (
	"fmt"
	"testing"

	"github.com/danielriddell21/retrievium/v2"
)

type searchCase struct {
	name     string
	haystack []int
	target   int
	want     int // -1 means not found
}

// sortedCases is what every searcher of a sorted slice has to get right.
// Each algorithm's test adds the shapes its own arithmetic is fragile on.
var sortedCases = []searchCase{
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

func withCases(extra ...searchCase) []searchCase {
	return append(append([]searchCase{}, sortedCases...), extra...)
}

func runSearchCases(t *testing.T, s retrievium.Searcher[int], cases []searchCase) {
	t.Helper()
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
}

// checkFinds asserts that s finds target in data, at an index holding it.
func checkFinds(t *testing.T, s retrievium.Searcher[int], data []int, target int) {
	t.Helper()
	got, ok := s.Search(data, target)
	if !ok {
		t.Fatalf("Search did not find target %d in slice", target)
	}
	if data[got] != target {
		t.Errorf("Search(%d) = index %d, haystack[%d] = %d", target, got, got, data[got])
	}
}

func checkFindsInLargeSorted(t *testing.T, s retrievium.Searcher[int]) {
	t.Helper()
	t.Run("random n=1000", func(t *testing.T) {
		data := sortedSlice(1000)
		checkFinds(t, s, data, data[500])
	})
}

func benchmarkSearcher(b *testing.B, s retrievium.Searcher[int]) {
	b.Helper()
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
