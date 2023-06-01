package main

import "fmt"

func main() {

	county := "Nairobi"                      // value assigned
	fmt.Printf("My county is %s \n", county) // display

	var yourCounty string //variable declaration
	fmt.Println("Enter your County")
	fmt.Scan(&yourCounty) //recieve a value
	fmt.Printf("Your County is %s \n", yourCounty)
}
