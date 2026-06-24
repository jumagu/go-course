package main

import "fmt"

func stringBytes() {
	greeting := "Hello World"

	character := greeting[0] // saves the word 'H' in its byte representation

	fmt.Println(character) // 72

	// ? Slicing strings
	slicedGreeting := greeting[3:7]
	fmt.Println(slicedGreeting) // lo W

	// ? Convert string to slice
	byteSlice := []byte(greeting)
	runeSlice := []rune(greeting)
	fmt.Println(byteSlice) // [72 101 108 108 111 32 87 111 114 108 100]
	fmt.Println(runeSlice) // [72 101 108 108 111 32 87 111 114 108 100]
}
