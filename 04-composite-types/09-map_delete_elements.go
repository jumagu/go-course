package main

import "fmt"

func mapDeleteElements() {
	map1 := map[int]string{
		1: "Hello",
		2: "World",
	}

	fmt.Println(map1, len(map1)) // map[1:Hello 2:World] 2

	// Delete one element
	delete(map1, 1)
	fmt.Println(map1, len(map1)) // map[2:World] 1

	// Clear the entire map
	clear(map1)
	fmt.Println(map1, len(map1)) // map[] 0
}
