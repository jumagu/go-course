package main

import "fmt"

func sliceOfSlice() {
	x := []string{"a", "b", "c", "d"}

	y := x[:2]
	z := x[2:3]
	a := x[:] // Make a copy of the original slice (keeps memory reference)

	fmt.Println(x) // [a b c d]
	fmt.Println(y) // [a b]
	fmt.Println(z) // [c]
	fmt.Println(a) // [a b c d]

	a[0] = "j"
	fmt.Println(a) // [j b c d]
	fmt.Println(x) // [j b c d]

	// ? Create a copy breaking the memory reference
	b := make([]string, 4)
	nCopiedElements := copy(b, x)

	fmt.Println(b, x)            // [j b c d] [j b c d]
	fmt.Println(nCopiedElements) // 4

	b[0] = "k"
	fmt.Println(b, x) // [k b c d] [j b c d]
}
