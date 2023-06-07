package main

import (
	"fmt"
)

func main() {

	// function call
	// fmt.Println("Start")
	// result := addNumbers(21, 13)

	// fmt.Println("Back to the main function")
	// fmt.Println("Sum:", result)
	ans1, ans2 := performCalculations(23, 5)
	fmt.Println("Sum ", ans1, "Mult ", ans2)
}

func addNumbers(n1 int, n2 int) int {
	fmt.Println("Inside the second function")
	sum := n1 + n2
	return sum
}

func performCalculations(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	mult := n1 * n2
	return sum, mult
}
