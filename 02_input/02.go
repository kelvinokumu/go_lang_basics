package main

import "fmt"

func main() {
	// create variables
	var name string
	var age int
	var city string

	// takes input value for city
	fmt.Print("Enter name for your city: ")
	fmt.Scan(&city)

	fmt.Printf("Name: %s \n", city)

	// take name and age input
	fmt.Println("Enter your name and age:")
	fmt.Scan(&name, &age)
	// fmt.Scanf("%s %d", &name, &age)

	// print input values
	fmt.Printf("Name: %s\n Age: %d \n", name, age)

}
