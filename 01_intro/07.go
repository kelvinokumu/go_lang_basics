package main

import (
	"fmt"
)

func main() {
	// Declare a float64 variable
	x := 5.5
	fmt.Printf(" %T \n", x)

	// Convert x to an int and assign it to a new variable
	y := int(x)
	fmt.Printf(" %T \n", y)

	// Print the value of y
	fmt.Printf("Value of y: %d \n", y)
}
