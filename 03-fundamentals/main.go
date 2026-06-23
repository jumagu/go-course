package main

import "fmt"

const IVA = 0.19

func main() {
	clientName := ""
	clientEmail := ""
	amount := 0.0

	fmt.Println("Ticket taxes calculator")

	fmt.Print("Client name: ")
	fmt.Scan(&clientName)

	fmt.Print("Client email: ")
	fmt.Scan(&clientEmail)

	fmt.Print("Amount: ")
	fmt.Scan(&amount)

	taxIVA := fmt.Sprintf("%d%%", int(IVA*100))
	taxAmount := amount * IVA
	totalAmount := amount + taxAmount

	result := `
=============Ticket=============

Client: %s - %s

--------------------------------

Taxes
IVA = %s

--------------------------------

Subtotal: %.2f
Tax amount: %.2f
Total: %.2f

================================
	`

	fmt.Printf(result, clientName, clientEmail, taxIVA, amount, taxAmount, totalAmount)
}
