// Package main scans an unsorted application log for error codes using Linear Search.
//
// Linear Search is the only algorithm in this package that works on unsorted
// input — and sometimes it is the right tool. When the dataset is small, the
// input cannot be pre-sorted, or the target is statistically near the front,
// a sequential scan requires no setup and is straightforward to reason about.
package main

import (
	"fmt"

	"github.com/danielriddell21/retrievium"
)

func main() {
	// HTTP status codes recorded in arrival order from an application log.
	// Unsorted — events arrive as they happen.
	statusCodes := []int{
		200, 200, 301, 200, 404, 200, 500,
		200, 403, 200, 200, 302, 200, 503,
		200, 200, 404, 200, 401, 200,
	}

	queries := []struct {
		code    int
		meaning string
	}{
		{404, "Not Found"},
		{500, "Internal Server Error"},
		{401, "Unauthorised"},
		{418, "I'm a Teapot"},
		{200, "OK"},
	}

	fmt.Println("=== Application Log Scanner ===")
	fmt.Println()
	fmt.Printf("%-5s  %-26s  %-7s  %s\n", "Code", "Meaning", "Index", "Result")
	fmt.Println("-----  --------------------------  -------  ------")

	s := retrievium.LinearSearcher[int]{}
	for _, q := range queries {
		idx, ok := s.Search(statusCodes, q.code)
		result := "not present"
		idxStr := "-"
		if ok {
			result = fmt.Sprintf("first occurrence at index %d", idx)
			idxStr = fmt.Sprintf("%d", idx)
		}
		fmt.Printf("%-5d  %-26s  %-7s  %s\n", q.code, q.meaning, idxStr, result)
	}
}
