// Write a go program that checks whether a given number is even or odd.

package main

import "fmt"

func main() {

	var number int
	fmt.Println("Enter a number")
	fmt.Scan(&number)
	isEven(number)

}

func isEven(number int) {
	status := number%2 == 0
	if status {
		fmt.Println(status)
	} else {
		fmt.Println(status)
	}
}

func odd() {

}
