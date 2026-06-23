package main

import "fmt"

func consoleInputs() {
	var name string
	var age int

	fmt.Printf("%q %d\n", name, age)
	fmt.Println("name memory address ->", &name)
	fmt.Println("age memory address ->", &age)

	fmt.Println("\nEnter your name:")
	/*
		'&' provides the memory address of the variable name, where the scanned value will be stored.
		This means the scanned value will point to the memory address of name var.
		In short, the scanned text will be stored in name.
	*/
	fmt.Scan(&name)

	fmt.Println("Enter your age:")
	fmt.Scan(&age)

	fmt.Printf("\nHi %s, you are %d years old\n", name, age)
	fmt.Println("name memory address ->", &name)
	fmt.Println("age memory address ->", &age)
}
