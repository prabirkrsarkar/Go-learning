package main

import (
	"fmt"
	"strings"
)

var productMaster = map[string]float64{ // short hand declaration is ONLY allowed at function level, not at package level. Hence we have to use var keyword here.
	"Cap":       150.00,
	"T-Shirt":   500.99,
	"Jeans":     999.50,
	"Socks":     200.00,
	"Jacket":    800.00,
	"Mug":       100.00,
	"DinnerSet": 1200.00,
	"LunchBox":  450.00,
}

func calculateItemPrice(itemCode string) (float64, bool) {
	basePrice, found := productMaster[itemCode]
	//Direct match, Item in productMaster, hence not on Sale
	if found {
		fmt.Println(" - Item : Base Price", itemCode, basePrice)
		return basePrice, true
	}

	//Itemcode not in productMaster
	if strings.HasSuffix(itemCode, "_Sale") { // Item is on sale
		itemCode = strings.TrimSuffix(itemCode, "_Sale")
		basePrice, found = productMaster[itemCode] // Lookup the productMaster map for baseprice
		if found {                                 // If item which is on Sale is found in the inventory
			salePrice := basePrice * 0.8 // Apply a 20% discount for sale items
			fmt.Println(itemCode, "on Sale!, Original price:", basePrice, "Sale price is", salePrice)
			return salePrice, true
		}
		// Item on Sale but not in productMaster!
		fmt.Println("Item on Sale but not found in productMaster!:", itemCode)
		return 0.0, false
	}

	// Item is neither in productMaster, nor on Sale!
	fmt.Println("Item not found in productMaster and Not on Sale!:", itemCode)
	return 0.0, false
}

func main() {
	custOrder1 := []string{"Cap", "T-Shirt_Sale", "Jeans_Sale", "Socks", "Mug", "Umbrella"}
	var subtotal float64 = 0.0

	for _, item := range custOrder1 {
		price, found := calculateItemPrice(item)
		if found {

			subtotal += price
		}
	}
	fmt.Printf("Subtotal for the order is: %.2f\n", subtotal)
}
