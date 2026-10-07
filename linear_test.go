package retrievium_test

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/danielriddell21/retrievium/v2"
)

func TestLinearSearch(t *testing.T) {
	s := retrievium.LinearSearcher[int]{}
	cases := []searchCase{
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
	runSearchCases(t, s, cases)
	t.Run("random n=1000 unsorted", func(t *testing.T) {
		r := rand.New(rand.NewPCG(99, 0))
		data := make([]int, 1000)
		for i := range data {
			data[i] = r.IntN(2000) - 1000
		}
		checkFinds(t, s, data, data[r.IntN(len(data))])
	})
}

func BenchmarkLinearSearch(b *testing.B) {
	benchmarkSearcher(b, retrievium.LinearSearcher[int]{})
}

func ExampleLinearSearcher() {
	s := retrievium.LinearSearcher[int]{}
	fmt.Println(s.Search([]int{9, 3, 7, 1, 5}, 7))
	// Output:
	// 2 true
}
