// Write a go program that find the second smallest element in an array of integers.

package main

import (
	"fmt"
	"math"
)

func findSecondSmallestElement(arr []int) int {
	smallest := math.MaxInt32
	secondSmallest := math.MaxInt32

	for _, num := range arr {
		if num < smallest {
			secondSmallest = smallest
			smallest = num
		} else if num < secondSmallest && num != smallest {
			secondSmallest = num
		}
	}

	return secondSmallest
}

func main() {
	// Example usage
	array := []int{10, 5, 8, 3, 15, 12}

	secondSmallestElement := findSecondSmallestElement(array)
	if secondSmallestElement == math.MaxInt32 {
		fmt.Println("There is no second smallest element.")
	} else {
		fmt.Println("Second smallest element:", secondSmallestElement)
	}
}
