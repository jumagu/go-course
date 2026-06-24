package main

import "fmt"

func structsAnonymous() {
	// ?  Anonymous struct
	var person1 struct {
		name    string
		age     int
		address string
	}

	fmt.Println(person1) // { 0 }

	person1.name = "Jacinthe"
	person1.age = 35
	person1.address = "78133 Lynch Mission"

	fmt.Println(person1) // {Jacinthe 35 78133 Lynch Mission}

	pet := struct {
		name string
		kind string
	}{
		name: "Dante",
		kind: "Dog",
	}

	fmt.Println(pet) // {Dante Dog}
}
