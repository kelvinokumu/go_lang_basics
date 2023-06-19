// write a go program to find duplicate values in an array

package main

import (
	"fmt"
)

func findDuplicateValues(arr []int) []int {
	duplicates := make([]int, 0)
	visited := make(map[int]bool)

	for _, num := range arr {
		// If the number is already visited, it's a duplicate
		if visited[num] {
			duplicates = append(duplicates, num)
		} else {
			visited[num] = true
		}
	}

	return duplicates
}

func main() {
	// Example usage
	array := []int{1, 2, 3, 4, 2, 5, 4, 6, 6}

	duplicateValues := findDuplicateValues(array)
	if len(duplicateValues) == 0 {
		fmt.Println("No duplicate values found.")
	} else {
		fmt.Println("Duplicate values:", duplicateValues)
	}
}
