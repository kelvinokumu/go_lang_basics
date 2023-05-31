// Program to check if it's February or not using switch without expression

package main

import "fmt"

func main() {
	numberOfDays := 31

	// switch without any expression
	switch {

	case 31 == numberOfDays:
		fmt.Println("It's January")

	case 28 == numberOfDays:
		fmt.Println("It's February")

	default:
		fmt.Println("Not February")
	}
}
