package main

import "fmt"

func main() {
	// declare an array
	var numbers [5]int

	// assign values
	numbers[0] = 12
	numbers[1] = 20
	numbers[4] = 15

	fmt.Println(numbers)

	// change an element
	numbers[4] = 55

	fmt.Println(numbers)

	languages := []string{"GO", "C", "JAVA"}

	size := len(languages)

	fmt.Println(size)
}
