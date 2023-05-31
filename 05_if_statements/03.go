// Program to relate two integers using =, > or < symbol

package main

import "fmt"

func main() {

	number1 := 12
	number2 := 20

	if number1 == number2 {
		fmt.Printf("Result: %d == %d", number1, number2)
	} else if number1 > number2 {
		fmt.Printf("Result: %d > %d", number1, number2)
	} else {
		fmt.Printf("Result: %d < %d", number1, number2)
	}
}

// if number1 == number2 {
// 	fmt.Printf("Result: %d == %d", number1, number2)
// } else if number1 == 1 || number == 2 || number == 3 {
// 	fmt.Printf("Result: %d > %d", number1, number2)
// } else {
// 	fmt.Printf("Result: %d < %d", number1, number2)
// }
