package main

import "fmt"

func main() {
	// array of numbers
	numbers := [5]int{32, 33, 20, 12, 22}

	// using range to iterate over the elements of array
	for index, item := range numbers {
		fmt.Printf("numbers[%d] = %d \n", index, item)
	}
}
