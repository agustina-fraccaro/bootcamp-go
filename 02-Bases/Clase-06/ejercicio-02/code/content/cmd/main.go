package main

import (
	"content/product"
	"fmt"
)

func main() {
	productMedium := product.NewProduct("medium")

	priceMedium := productMedium.GetPrice()

	fmt.Println("Precio medium:", priceMedium)

	productLarge := product.NewProduct("large")

	priceLarge := productLarge.GetPrice()

	fmt.Println("Precio large:", priceLarge)

	productSmall := product.NewProduct("small")

	priceSmall := productSmall.GetPrice()

	fmt.Println("Precio small:", priceSmall)
}
