package main

import "fmt"

// declare a structure of type person
type Person struct {
	name  string
	age   int
	email string
}

func main() {
	// Create a new Person instance
	p := Person{
		name:  "Alice",
		age:   30,
		email: "alice@example.com",
	}

	// Accessing struct fields
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("Email:", p.email)
}
