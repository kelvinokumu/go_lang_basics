package main

import "fmt"

func main() {
	fmt.Println("Start point")
	for i := 0; i < 5; i++ {
		greetings()
	}

}

func greetings() {
	morning_greetings := "Good Morning"
	fmt.Println(morning_greetings)
}
