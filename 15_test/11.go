// Write a go program to find the largest element in an array of integers.

package main

import (
	"fmt"
)

func findLargestElement(arr []int) int {
	largest := arr[0]

	for _, num := range arr {
		if num > largest {
			largest = num
		}
	}

	return largest
}

func main() {
	// Example usage
	array := []int{10, 5, 8, 3, 15, 12}

	largestElement := findLargestElement(array)
	fmt.Println("Largest element:", largestElement)
}
