package main

import "fmt"

type ShoppingItem struct {
	title        string
	description  string
	quanity      int
	pricePerUnit int
}

func main() {
	var (
		total          = 0
		shoppingBasket = []ShoppingItem{{title: "Lego Set", description: "4000 pieces", quanity: 1, pricePerUnit: 600}, {title: "Plushy", description: "plush toy", quanity: 3, pricePerUnit: 5}}
	)

	fmt.Println("Title, Description, Quantity, Price per unit, Total")

	for _, item := range shoppingBasket {
		itemTotal := item.quanity * item.pricePerUnit
		total += itemTotal
		fmt.Printf("%s, %s, %d, %dGBP, %dGBP\n", item.title, item.description, item.quanity, item.pricePerUnit, itemTotal)
	}

	fmt.Printf("\nTotal: %d GBP\n", total)
}
