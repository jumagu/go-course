package main

import "fmt"

// ? Constants: usually declared outside the functions
// ? can't be changed during runtime

// const appName = "go course" // For per file scoped constants, use LowerCamelCase

const AppName = "Go Course" // use UpperCamelCase case for global scoped
const MaxUsers = 1000

func constants() {
	// ? Variables
	var name string = "Juan" // Long way, more descriptive
	last_name := "Gutierrez" // Short way, easy to read

	fmt.Printf("Name: %s %s \n", name, last_name)
	fmt.Println(AppName, MaxUsers)
}
