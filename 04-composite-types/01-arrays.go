package main

import "fmt"

func arrays() {
	// var numberList [3]int // Initialized with Zero Values [0 0 0]

	// ? Array literal
	// var numberList = [3]int{10, 20, 30}
	var numberList = [...]int{10, 20, 30} // Simplified sintax

	fmt.Println(numberList)

	numberList[0] = 55 // Reassing the value of position [0]

	fmt.Println(numberList)
	fmt.Println(numberList[0])   // 55 - Access the value in the [0] position
	fmt.Println(len(numberList)) // 3 - Array length
}
