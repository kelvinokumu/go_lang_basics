package main

import (
	"fmt"
	"strings"
)

func main() {
	// Declaring a string variable
	var myString string = "Hello, World!"

	// Printing the string
	fmt.Println(myString)

	// Concatenating strings
	var firstName string = "John"
	var lastName string = "Doe"
	var fullName string = firstName + " " + lastName
	var fullNames string = firstName + lastName
	fmt.Println(fullName)
	fmt.Println(fullNames)

	// Getting the length of a string
	fmt.Println(len(myString))
	// length, size

	// Accessing individual characters in a string
	fmt.Println(myString[3])

	// Check if a string contains a substring
	fmt.Println(strings.Contains(myString, "kenya"))

	// Replacing a substring
	myString = strings.Replace(myString, "World", "Go", 1)
	fmt.Println(myString)
}
