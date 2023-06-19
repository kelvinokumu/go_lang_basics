// Write a go program to check whether a given year is a leap year or not.

package main

import (
	"fmt"
)

func isLeapYear(year int) bool {
	// Leap years are divisible by 4
	// But century years are not leap years unless they are divisible by 400
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func main() {
	// Example usage
	year := 2024

	if isLeapYear(year) {
		fmt.Println(year, "is a leap year.")
	} else {
		fmt.Println(year, "is not a leap year.")
	}
}
