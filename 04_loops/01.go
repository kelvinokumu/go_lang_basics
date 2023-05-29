// Program to print numbers between 1 and 5
package main

import (
	"fmt"
)

func main() {
	number := 1
	// for condition{
	// 	do something
	// 	increment/decrement
	// }

	fmt.Println("Before for")
	for number < 15 {
		fmt.Println(number)
		number++
	}
	fmt.Println("After for")
}
