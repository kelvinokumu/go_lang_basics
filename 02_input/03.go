package main

import "fmt"

func main() {
	// variable declaration
	var number int32
	var temperature float32
	var sunny bool

	// take integer input
	fmt.Println("Enter a number:")
	fmt.Scanf("%d", &number)
	fmt.Printf("The number is %d \n", number)

	// take float input
	fmt.Println("Enter the current temperature:")
	fmt.Scanf("%g", &temperature)

	// take boolean input
	fmt.Println("Is the day sunny? (True/False)")
	fmt.Scanf("%t", &sunny)

	fmt.Printf("Current temperature: %g\nIs the day Sunny? %t\n", temperature, sunny)
}
