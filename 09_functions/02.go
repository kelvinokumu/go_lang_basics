package main

import "fmt"

func main() {
	// function calling
	addNumbers(45, 12, "kelvin")
	getSquare(15)
}

// function with parameters
func addNumbers(n1 int, n2 int, name string) {
	fmt.Printf("n1 is %d n2 is %d \n", n1, n2)
	sum := n1 + n2
	fmt.Println("Sum : ", sum)
	fmt.Println("Good Morning ", name)
}

func getSquare(num int) {
	square := num * num
	fmt.Printf("The square of %d is %d \n", num, square)
}
