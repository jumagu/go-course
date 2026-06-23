package main

import "fmt"

// const Pi = 3.1415 // Untyped const - they adapt to the context
const Pi float32 = 3.1415 // Typed const - they have declared a static type

func typedConstants() {
	var number float64 = float64(Pi)

	fmt.Println(number)
}
