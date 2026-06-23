package main

import "fmt"

func zeroValues() {
	// * Zero Values
	// * By default, Go assigns a default value to non initilized vars

	var i int      // Defaults to 0 for ints
	var s string   // Defaults to empty string
	var ok bool    // Defaults to false for booleans
	var nums []int // Defaults to an empty array. Type have no affect

	fmt.Printf("i=%d\n", i)       // i=0
	fmt.Printf("s=%q\n", s)       // s=""
	fmt.Printf("ok=%v\n", ok)     // ok=false
	fmt.Printf("nums=%v\n", nums) // nums=[]
}
