// Package main locates a salary in a pay scale using Ternary Search.
//
// Ternary Search divides the search space into three parts per iteration.
// While Binary Search is typically faster in practice (fewer comparisons per
// element), Ternary Search demonstrates that the divide-and-conquer principle
// generalises beyond halving — useful when partition costs differ across ranges.
package main

import (
	"fmt"

	"github.com/danielriddell21/retrievium/v2"
)

func main() {
	// Sorted annual salaries (£) for each pay grade in a company.
	payscale := []int{
		22000,  // Grade 1
		26000,  // Grade 2
		30000,  // Grade 3
		35000,  // Grade 4
		42000,  // Grade 5
		50000,  // Grade 6
		60000,  // Grade 7
		72000,  // Grade 8
		88000,  // Grade 9
		110000, // Grade 10
	}

	queries := []struct {
		name   string
		salary int
	}{
		{"Alice", 42000},
		{"Bob", 72000},
		{"Carol", 55000},
		{"Dave", 22000},
		{"Eve", 110000},
	}

	fmt.Println("=== Salary Band Lookup ===")
	fmt.Println()
	fmt.Printf("%-8s  %-10s  %-7s  %s\n", "Employee", "Salary (£)", "Index", "Grade")
	fmt.Println("--------  ----------  -------  -------")

	s := retrievium.TernarySearcher[int]{}
	for _, q := range queries {
		idx, ok := s.Search(payscale, q.salary)
		grade := "—"
		idxStr := "-"
		if ok {
			grade = fmt.Sprintf("Grade %d", idx+1)
			idxStr = fmt.Sprintf("%d", idx)
		}
		fmt.Printf("%-8s  %-10d  %-7s  %s\n", q.name, q.salary, idxStr, grade)
	}
}
