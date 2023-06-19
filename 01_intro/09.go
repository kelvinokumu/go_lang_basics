package main

import "fmt"

func main() {
	// Integer
	num := 42
	fmt.Printf("Integer: %d\n", num)

	// Floating-point
	pi := 3.14159
	fmt.Printf("Floating-point: %.2f\n", pi)

	// String
	name := "Alice"
	fmt.Printf("String: %s\n", name)

	// Boolean
	flag := true
	fmt.Printf("Boolean: %t\n", flag)

	// Character
	ch := 'A'
	fmt.Printf("Character: %c\n", ch)

	// Pointer
	ptr := &num
	fmt.Printf("Pointer: %p\n", ptr)
}
