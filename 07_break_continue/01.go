package main

import "fmt"

func main() {
	for counter := 1; counter <= 10; counter++ {
		if counter%2 == 0 {
			continue
		}
		fmt.Println(counter)
	}

	for counter := 1; counter <= 10; counter++ {
		if counter%2 == 0 {
			break
		}
		fmt.Println(counter)
	}
}
