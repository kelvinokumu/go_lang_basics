// Write a Go program that converts a temperature from Celsius to Fahrenheit.

package main

import (
	"fmt"
)

func celsiusToFahrenheit(celsius float64) float64 {
	fahrenheit := (celsius * 9 / 5) + 32
	return fahrenheit
}

func main() {
	// Example usage
	celsius := 25.0

	fahrenheit := celsiusToFahrenheit(celsius)
	fmt.Printf("%.1f degrees Celsius is %.1f degrees Fahrenheit\n", celsius, fahrenheit)
}

// In this program, the `celsiusToFahrenheit` function takes a temperature in Celsius as a `float64` input and returns the equivalent temperature in Fahrenheit as a `float64`.

// The function uses the formula `(Celsius * 9/5) + 32` to convert Celsius to Fahrenheit. It multiplies the Celsius temperature by 9/5 and then adds 32 to get the Fahrenheit temperature.

// In the `main` function, an example usage is shown where the variable `celsius` is set to `25.0`. The `celsiusToFahrenheit` function is called with this Celsius temperature, and the resulting Fahrenheit temperature is printed to the console using `Printf` to format the output. In this case, the output would be `25.0 degrees Celsius is 77.0 degrees Fahrenheit`.
