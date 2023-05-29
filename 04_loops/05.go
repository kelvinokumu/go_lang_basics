// Program to print numbers for natural numbers 1 + 2 + 3 + ... +n

package main

import "fmt"

func main() {
	var n, sum = 10, 0

	for i := 1; i <= n; i++ {
		sum += i // sum = sum + i
		// fmt.Printf("%d + %d \n", sum, i)
		fmt.Println("sum =", sum)
	}

	// fmt.Println("sum =", sum)
}
