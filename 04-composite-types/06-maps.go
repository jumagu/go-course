package main

import "fmt"

func mapsGo() {
	// [keyType]valueType
	var map1 map[string]int

	// map1["a"] = 0 // Panic: https://pkg.go.dev/golang.org/x/tools/go/analysis/passes/nilness#nilderef

	fmt.Println(map1)        // map[]
	fmt.Println(map1 == nil) // true

	// var map2 map[string]int = map[string]int{}
	map2 := map[string]int{}

	fmt.Println(map2)        // map[]
	fmt.Println(map2 == nil) // false

	map2["a"] = 1
	map2["b"] = 2
	map2["c"] = 3
	map2["d"] = 4

	fmt.Println(map2)      // map[a:1 b:2 c:3 d:4]
	fmt.Println(map2["c"]) // 3
}
