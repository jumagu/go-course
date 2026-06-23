package main

import "fmt"

// ? Declaring multiple variables

// Declared type
var a, b, c int = 1, 2, 3

// Infered type
var name, city = "María", "Bogotá"
var x, y = 10, 3.1415

// Recommended for related variables
var (
	age      int  = 22
	isActive bool = true
)

func multipleVariables() {
	a, b, c := 1, 2, 3

	fmt.Println(a, b, c)

	a, b, c = 10, 2, 3 // Re-assigning new values

	fmt.Println(a, b, c)

	a, b, c, d := 3, 4, 5, 9 // Re-assigning and creating a new one

	fmt.Println(a, b, c, d)
}
