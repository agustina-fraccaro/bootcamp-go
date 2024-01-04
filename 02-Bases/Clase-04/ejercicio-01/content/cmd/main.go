package main

import (
	"content/product"
)

func main() {
	product := product.Product{ID: 1, Name: "Product 1", Price: 100, Description: "Description 1", Category: "Category 1"}
	product.Save()
	product.GetAll()
}
