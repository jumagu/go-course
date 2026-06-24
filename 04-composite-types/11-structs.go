package main

import "fmt"

func structsGo() {
	type person struct {
		name    string
		age     int
		address string
	}

	var person1 person
	var person2 = person{}

	fmt.Println(person1) // { 0 }
	fmt.Println(person2) // { 0 }

	// fmt.Println(person1 == nil) // invalid operation: person1 == nil (mismatched types person and untyped nil)
	// fmt.Println(person2 == nil) // invalid operation: person1 == nil (mismatched types person and untyped nil)

	// ? Struct literal
	// var person3 person = person{
	// 	name:    "Chesley",
	// 	age:     22,
	// 	address: "62349 Kozey Spurs",
	// }

	person3 := person{
		name:    "Chesley",
		age:     22,
		address: "62349 Kozey Spurs",
	}

	fmt.Println(person3)         // {Chesley 22 62349 Kozey Spurs}
	fmt.Println(person3.name)    // Chesley
	fmt.Println(person3.age)     // 22
	fmt.Println(person3.address) // 62349 Kozey Spurs

	person3.name = "Brannon"

	fmt.Println(person3.name) // Brannon
}
