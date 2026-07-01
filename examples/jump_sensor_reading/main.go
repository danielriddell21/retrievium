// Package main finds a temperature reading in a sorted sensor log using Jump Search.
//
// Jump Search's O(√n) complexity sits between Linear and Binary Search. It
// works well when jumping forward in the dataset is cheap but backtracking is
// costly — for instance, reading sequentially from a sorted log file where
// seeking backward is expensive.
package main

import (
	"fmt"

	"github.com/danielriddell21/retrievium"
)

func main() {
	// Sorted temperature readings (°C × 10 to avoid floats) from a weather station.
	readings := []int{
		-82, -75, -61, -50, -43, -38, -29, -20,
		-15, -10, -5, 0, 6, 12, 18, 23,
		27, 31, 35, 38, 42, 47, 51, 56,
		60, 64, 68, 72, 75, 79, 83, 88,
		91, 95, 98, 102, 107, 111, 116, 120,
	}

	queries := []struct {
		label string
		value int
	}{
		{"0.0°C (freezing point)", 0},
		{"3.8°C (spring morning)", 38},
		{"9.8°C (summer peak)", 98},
		{"12.0°C (above range)", 120},
		{"5.5°C (not recorded)", 55},
	}

	fmt.Println("=== Sensor Log Lookup ===")
	fmt.Println()
	fmt.Printf("%-28s  %-7s  %s\n", "Query", "Index", "Result")
	fmt.Println("----------------------------  -------  ------")

	s := retrievium.JumpSearcher[int]{}
	for _, q := range queries {
		idx, ok := s.Search(readings, q.value)
		result := "not found"
		idxStr := "-"
		if ok {
			result = fmt.Sprintf("found at index %d", idx)
			idxStr = fmt.Sprintf("%d", idx)
		}
		fmt.Printf("%-28s  %-7s  %s\n", q.label, idxStr, result)
	}
}
