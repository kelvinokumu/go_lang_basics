package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)             // Convert string to lowercase
	s = strings.ReplaceAll(s, " ", "") // Remove spaces from the string

	// Check if the string is a palindrome
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	word := "Madam"
	if isPalindrome(word) {
		fmt.Printf("%s is a palindrome.\n", word)
	} else {
		fmt.Printf("%s is not a palindrome.\n", word)
	}
}
