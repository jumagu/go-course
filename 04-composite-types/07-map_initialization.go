package main

import "fmt"

func mapInitialization() {
	// Map literal
	footballTeams := map[string][]string{
		"France":  {"Mbappe", "Kanté", "Olise"},
		"England": {"Pickford", "Kane", "Rice"},
		"Germany": {"Neuer", "Sané", "Havertz"},
	}

	fmt.Println(footballTeams)              // map[England:[Pickford Kane Rice] France:[Mbappe Kanté Olise] Germany:[Neuer Sané Havertz]]
	fmt.Println(footballTeams["England"])   // [Pickford Kane Rice]
	fmt.Println(footballTeams["France"][2]) // Olise

	// *********************************************

	ages := make(map[int][]string, 10) // Can create a map using make(), assigning it a predefined size for memory efficiency

	fmt.Println(ages) // map[]

	// *********************************************

	// ? Comma ok idom
	value, ok := footballTeams["Germany"]
	fmt.Println(value, ok) // [Neuer Sané Havertz] true

	spain, spainExist := footballTeams["Spain"]
	fmt.Println(spain, spainExist) // [] false
	fmt.Println(spain == nil)      // true
}
