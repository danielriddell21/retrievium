// Package main searches a library catalogue by call number using Fibonacci Search.
//
// Fibonacci Search divides the search range using Fibonacci numbers rather
// than halving. This avoids division operations and accesses elements that are
// closer to the front of the array — an advantage when memory access costs
// increase with distance, as in tape storage or certain cache layouts.
package main

import (
	"fmt"

	"github.com/danielriddell21/retrievium/v2"
)

func main() {
	type Book struct {
		CallNumber int
		Title      string
	}

	catalogue := []Book{
		{100, "Introduction to Philosophy"},
		{153, "Cognitive Psychology"},
		{220, "The Bible: A History"},
		{305, "Sociology of Everyday Life"},
		{370, "Principles of Education"},
		{428, "English Grammar in Use"},
		{510, "Discrete Mathematics"},
		{530, "Classical Mechanics"},
		{610, "Human Anatomy"},
		{641, "The Art of French Cooking"},
		{720, "Architectural Theory"},
		{780, "Music Theory Fundamentals"},
		{823, "Victorian Fiction"},
		{900, "World History"},
		{973, "American History"},
	}

	callNumbers := make([]int, len(catalogue))
	for i, b := range catalogue {
		callNumbers[i] = b.CallNumber
	}

	queries := []int{510, 220, 900, 999, 100}

	fmt.Println("=== Library Catalogue Search ===")
	fmt.Println()
	fmt.Printf("%-12s  %-7s  %s\n", "Call Number", "Index", "Title")
	fmt.Println("------------  -------  -----")

	s := retrievium.FibonacciSearcher[int]{}
	for _, cn := range queries {
		idx, ok := s.Search(callNumbers, cn)
		title := "not found"
		idxStr := "-"
		if ok {
			title = catalogue[idx].Title
			idxStr = fmt.Sprintf("%d", idx)
		}
		fmt.Printf("%-12d  %-7s  %s\n", cn, idxStr, title)
	}
}
