// Write  go program that takes a slice of integers as input and calculates the average of the numbers.

package main

import (
	"fmt"
)

func calculateAverage(numbers []int) float64 {
	sum := 0
	count := len(numbers)

	for _, num := range numbers {
		sum += num
	}

	average := float64(sum) / float64(count)
	return average
}

func main() {
	// Example usage
	numbers := []int{5, 10, 15, 20, 25}

	average := calculateAverage(numbers)
	fmt.Printf("Average: %.2f\n", average)
}
