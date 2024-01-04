package product

import "fmt"

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

var products []Product

func (p Product) Save() {
	products = append(products, p)
}

func (p Product) GetAll() {
	for _, product := range products {
		fmt.Println(product)
	}
}

func getById(id int) Product {
	for _, product := range products {
		if product.ID == id {
			return product
		}
	}
	return Product{}
}
