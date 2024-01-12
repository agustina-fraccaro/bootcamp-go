package storage

import "app/internal/product"

type ProductStorage interface {
	ReadProducts() ([]product.Product, error)
	WriteProducts([]product.Product) error
}
