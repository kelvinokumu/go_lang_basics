// Program to create an array and prints its elements

package main

import "fmt"

func main() {
	// An array is a collection of similar types of data.

	// syntax
	// var array_variable = [size]datatype{elements of array}

	// declare array variable of type integer
	// defined size [size=5]

	// 3 ways of declaring an array

	// var arrayOfInteger = [5]int{1, 5, 8, 0, 3}

	// var arrayOfInteger = [...]int{1, 5, 8, 0, 3}

	arrayOfInteger := []int{1, 5, 8, 0, 3}
	fmt.Printf("%T \n", arrayOfInteger) // checking the type

	// loop through elements in an array
	for i := 0; i < 5; i++ {
		fmt.Println(arrayOfInteger[i] + 1000)
	}

	languages := []string{"GO", "C", "JAVA"}
	fmt.Printf("%T \n", languages) // checking the type
	fmt.Println(languages[0])
	fmt.Println(languages[1])
	fmt.Println(languages[2])

	// fmt.Println(arrayOfInteger)
}
