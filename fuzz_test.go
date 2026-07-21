package retrievium_test

import (
	"slices"
	"testing"

	"github.com/danielriddell21/retrievium/v2"
)

func allSearchers() []retrievium.Searcher[int] {
	return []retrievium.Searcher[int]{
		retrievium.LinearSearcher[int]{},
		retrievium.BinarySearcher[int]{},
		retrievium.TernarySearcher[int]{},
		retrievium.JumpSearcher[int]{},
		retrievium.FibonacciSearcher[int]{},
	}
}

// FuzzSearchers checks every searcher against the standard library: on a sorted
// haystack, the found flag must match slices.Contains, and a reported index
// must actually hold the target.
func FuzzSearchers(f *testing.F) {
	f.Add([]byte{1, 3, 5, 7, 9}, byte(7))
	f.Add([]byte{}, byte(0))
	f.Add([]byte{2, 2, 2}, byte(2))
	f.Fuzz(func(t *testing.T, data []byte, target byte) {
		haystack := make([]int, len(data))
		for i, b := range data {
			haystack[i] = int(b)
		}
		slices.Sort(haystack)
		want := slices.Contains(haystack, int(target))
		for _, s := range allSearchers() {
			idx, ok := s.Search(haystack, int(target))
			if ok != want {
				t.Errorf("%s.Search(%v, %d) found = %v, want %v", s.Name(), haystack, target, ok, want)
			}
			if ok && haystack[idx] != int(target) {
				t.Errorf("%s.Search(%v, %d) = index %d holding %d, want %d",
					s.Name(), haystack, target, idx, haystack[idx], target)
			}
		}
	})
}
