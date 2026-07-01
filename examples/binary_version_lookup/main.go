// Package main finds a software version in a sorted changelog using Binary Search.
//
// Binary Search's O(log n) complexity makes it the natural choice for any
// large sorted dataset. A version registry with thousands of releases can be
// probed in under 20 comparisons — far better than scanning from the top.
package main

import (
	"fmt"

	"github.com/danielriddell21/retrievium"
)

func main() {
	// Sorted list of released version numbers (encoded as integers: major*10000 + minor*100 + patch).
	versions := []int{
		10000, // 1.0.0
		10001, // 1.0.1
		10100, // 1.1.0
		10200, // 1.2.0
		10201, // 1.2.1
		20000, // 2.0.0
		20100, // 2.1.0
		20200, // 2.2.0
		20201, // 2.2.1
		30000, // 3.0.0
		30100, // 3.1.0
		30101, // 3.1.1
	}

	queries := []struct {
		label   string
		encoded int
	}{
		{"2.1.0", 20100},
		{"3.1.1", 30101},
		{"1.3.0", 10300},
		{"2.0.0", 20000},
	}

	fmt.Println("=== Version Registry Lookup ===")
	fmt.Println()
	fmt.Printf("%-10s  %-7s  %s\n", "Version", "Index", "Status")
	fmt.Println("----------  -------  --------")

	s := retrievium.BinarySearcher[int]{}
	for _, q := range queries {
		idx, ok := s.Search(versions, q.encoded)
		status := "released"
		idxStr := fmt.Sprintf("%d", idx)
		if !ok {
			status = "not found"
			idxStr = "-"
		}
		fmt.Printf("%-10s  %-7s  %s\n", q.label, idxStr, status)
	}
}
