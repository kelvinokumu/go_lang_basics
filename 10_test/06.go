// Write a Go program that finds the sum of all
// the numbers in a given array.

package main

import "fmt"

func main() {
	var sum int = 0
	numbers := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	max := len(numbers)
	// fmt.Println(max)
	for i := 0; i < max; i++ {
		sum = sum + numbers[i]
		fmt.Println(sum)
	}

}
