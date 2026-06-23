package main

import "fmt"

func explicitConvertions() {
	// * Explicit conversions
	var number1 int = 10
	var number2 float64 = 3.5

	// total := number1 + int(number2) // This is explicit
	total := float64(number1) + number2 // Did use float64 instead in order to keept the .5

	fmt.Println(total)
}
