package main

import "fmt"

// Write a program which will find all such numbers which are divisible by 7
// but are not a multiple of 5, between 1 and 50 (both included).

func main() {
	start := 1
	stop := 50
	// and  &&  -  both conditions have to be true
	// or   ||  -  one condition has to be true
	// not   !  -  none is true
	for i := start; i <= stop; i++ {
		if i%7 == 0 && i%5 != 0 {
			fmt.Println(i)
		}
	}

}
