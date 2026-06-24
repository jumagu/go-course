package main

import "fmt"

func commaOkIdiom() {
	map1 := map[int]string{
		1: "Hello",
		2: "World",
	}

	/*
		Accessing a value that doesn't exist does not
		return a nil, as we might think. It returns a zero value.

		That is why this exists (as if we were extracting from a tuple in Python)
		to get the value and a boolean flag that tells whether the value exists or not.
	*/

	value, ok := map1[3]

	fmt.Println(value, ok) // "" false
}
