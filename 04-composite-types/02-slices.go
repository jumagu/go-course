package main

import (
	"fmt"
	"slices"
)

func slicesExample() {
	// ? Slice literal
	var slice1 = []int{1, 2, 3} // A subtle difference from arrays: the square brackets are left empty

	fmt.Println(slice1)

	slice1[0] = 6

	fmt.Println(slice1)
	fmt.Println(slice1[0])

	var emptySlice []int
	var otherEmptySlice = []int{}

	fmt.Println(emptySlice)        // []
	fmt.Println(emptySlice == nil) // true - Slice was not initialized and is empty. So this it is a nil (Also known as null in other progamming langs)

	fmt.Println(otherEmptySlice)        // []
	fmt.Println(otherEmptySlice == nil) // false - Slice was initialized and is empty. So this it is not a nil

	// ****************************************

	/*
		? Slices only can be compared using Equal method from 'slices' package
	*/

	x := []int{1, 2, 3, 4, 5}
	y := []int{1, 2, 3, 4, 5}
	z := []int{1, 2, 3, 4, 5, 6}
	// s := []string{"a", "b", "c"}

	fmt.Println(slices.Equal(x, y)) // true
	fmt.Println(slices.Equal(x, z)) // false
	// fmt.Println(slices.Equal(x, s)) // Can't compare slices with different types
}
