package main

import "fmt"

func main() {
	// how to declare variables
	fmt.Println("Hello, World!")
	var price int = 2 // 1
	fmt.Println(price)

	var number = 150.77 // 2
	fmt.Println(number)
	fmt.Printf("The type is %T \n", number)

	number_1 := 65 // 3
	fmt.Println(number_1)
	fmt.Printf("The type is %T \n", number_1) // Printf gives you the type of the variable
}
