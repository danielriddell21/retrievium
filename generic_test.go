package retrievium_test

import (
	"testing"

	"github.com/danielriddell21/retrievium/v2"
)

// The searchers are generic, so they search any cmp.Ordered type, not just int.
func TestGenericStringSearch(t *testing.T) {
	haystack := []string{"apple", "banana", "cherry", "date", "elderberry"}
	searchers := []retrievium.Searcher[string]{
		retrievium.LinearSearcher[string]{},
		retrievium.BinarySearcher[string]{},
		retrievium.TernarySearcher[string]{},
		retrievium.JumpSearcher[string]{},
		retrievium.FibonacciSearcher[string]{},
	}
	for _, s := range searchers {
		if idx, ok := s.Search(haystack, "cherry"); !ok || idx != 2 {
			t.Errorf("%s.Search(cherry) = (%d, %v), want (2, true)", s.Name(), idx, ok)
		}
		if _, ok := s.Search(haystack, "fig"); ok {
			t.Errorf("%s.Search(fig) found a value that is absent", s.Name())
		}
	}
}

// Every searcher satisfies Searcher[int]; the assertions also confirm they
// satisfy Searcher for other cmp.Ordered element types.
var (
	_ retrievium.Searcher[int] = retrievium.LinearSearcher[int]{}
	_ retrievium.Searcher[int] = retrievium.BinarySearcher[int]{}
	_ retrievium.Searcher[int] = retrievium.TernarySearcher[int]{}
	_ retrievium.Searcher[int] = retrievium.JumpSearcher[int]{}
	_ retrievium.Searcher[int] = retrievium.FibonacciSearcher[int]{}

	_ retrievium.Searcher[string] = retrievium.BinarySearcher[string]{}
)
