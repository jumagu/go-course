package main

import "fmt"

type Item struct {
	sku   string
	name  string
	price float64
	qty   int
}

type Order struct {
	id       string
	customer string
	items    []Item
	meta     map[string]string
}

func main() {
	order1 := Order{
		id:       "7afd29f9-0394-4aad-bd0a-cce277dcf0a1",
		customer: "Davin",
		items:    []Item{{sku: "4AS85SAD6", name: "Handmade Metal Mouse", price: 10.99, qty: 10}, {sku: "554AS96S2", name: "Gorgeous Granite Chair", price: 562.00, qty: 2}},
		meta:     map[string]string{"city": "Pollichbury", "source": "online"},
	}

	fmt.Println(order1.id)           // 7afd29f9-0394-4aad-bd0a-cce277dcf0a1
	fmt.Println(order1.customer)     // Davin
	fmt.Println(order1.items[0])     // {4AS85SAD6 Handmade Metal Mouse 10.99 10}
	fmt.Println(order1.meta["city"]) // Pollichbury

	order1.meta["coupon"] = "JUNE2026"

	fmt.Println(order1.meta["coupon"]) // JUNE2026

	fmt.Println(order1) // {7afd29f9-0394-4aad-bd0a-cce277dcf0a1 Davin [{4AS85SAD6 Handmade Metal Mouse 10.99 10} {554AS96S2 Gorgeous Granite Chair 562 2}] map[city:Pollichbury coupon:JUNE2026 source:online]}
}
