package main

import (
	"fmt"
	"maps"
)

func mapCompare() {
	map1 := map[int]string{
		1: "Hello",
		2: "World",
	}

	map2 := map[int]string{
		1: "Hello",
		2: "World",
	}

	map3 := map[int]string{
		1: "Hello",
		2: "World",
		3: "Go",
	}

	/*
		Just like with slices, maps can only be compared using the Equal fn
		from the standard package of the type, in this case it is "maps"
	*/

	fmt.Println(maps.Equal(map1, map2)) // true
	fmt.Println(maps.Equal(map1, map3)) // false
}
