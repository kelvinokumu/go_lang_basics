package main

import "fmt"

func main() {
	fmt.Println("main")
	var number int

	fmt.Println("Enter a number")
	// get number from the user
	fmt.Scan(&number)

	// calling function
	getSquare(number)
}

func getSquare(num int) {
	// perform calculation
	square := num * num

	// display the results
	fmt.Println("The square is ", square)
}
