package main

import "fmt"

func sliceFunctions() {
	slice1 := []int{1, 2, 3}
	fmt.Println(slice1)

	// * len()
	fmt.Println(len(slice1))

	// * append()
	slice1 = append(slice1, 10, 11, 12)
	fmt.Println(len(slice1))
	fmt.Println(slice1)

	// * capacity()
	fmt.Println(cap(slice1)) // Capacity is dynamic in slices

	// * make()
	// slice2 := make([]int, 5) // This creates a slice of integers with a size of 5
	// fmt.Println(slice2) // [0 0 0 0 0] Initialized with Zero Values

	slice3 := make([]int, 0, 5) // This creates an integer slice with a size of 0 and a capacity of 5
	fmt.Println(slice3)         // []

	/*
		Capacity is not a restriction; it indicates how many
		elements are available in that array from the starting position
		of the slice to the end of the array.
	*/
	slice3 = append(slice3, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(slice3)                   // [1 2 3 4 5 6 7]
	fmt.Println(len(slice3), cap(slice3)) // 7 10

	// * clear()
	clear(slice3)
	/*
		Contrary to what it might seem, `clear` does not actually leave
		the array empty; rather, it sets all elements to zero. It keeps
		the size and capacity intact.
	*/
	fmt.Println(slice3)                   // [0 0 0 0 0 0 0]
	fmt.Println(len(slice3), cap(slice3)) // 7 10
}
