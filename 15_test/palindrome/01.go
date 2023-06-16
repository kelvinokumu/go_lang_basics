package main

import (
	"fmt"
	"strings"
)

func isPalindrome(input string) bool {
	// Remove spaces and convert to lowercase
	input = strings.ReplaceAll(input, " ", "")
	input = strings.ToLower(input)

	// Check if the input is a palindrome
	length := len(input)
	for i := 0; i < length/2; i++ {
		if input[i] != input[length-1-i] {
			return false
		}
	}

	return true
}

func main() {
	// Prompt the user to enter a string
	fmt.Print("Enter a string: ")
	var input string
	fmt.Scanln(&input)

	// Check if the input string is a palindrome
	if isPalindrome(input) {
		fmt.Println("The string is a palindrome.")
	} else {
		fmt.Println("The string is not a palindrome.")
	}
}
