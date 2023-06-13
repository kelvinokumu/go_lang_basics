package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a number as a command-line argument.")
		return
	}

	number, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid number provided.")
		return
	}

	// Check if the number is odd
	if number%2 == 1 {
		fmt.Println("The number is odd.")
	} else {
		fmt.Println("The number is even.")
	}

	// Check if the number is divisible by 8
	if number%8 == 0 {
		fmt.Println("The number is divisible by 8.")
	} else {
		fmt.Println("The number is not divisible by 8.")
	}
}
