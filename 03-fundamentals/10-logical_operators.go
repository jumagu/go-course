package main

import "fmt"

func logicalOperators() {
	// Common logial operators
	// AND: &&
	// OR: ||
	// NOT: !

	a := true
	b := false

	fmt.Println(a && b) // false
	fmt.Println(a || b) // true
	fmt.Println(!a)     // false
	fmt.Println(!b)     // true
}
