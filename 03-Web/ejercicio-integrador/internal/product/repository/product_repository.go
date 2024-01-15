package repository

import (
	"app/internal/product"
	"errors"
)

var (
	ErrCodeValueAlreadyExists = errors.New("code value already exists")
	ErrProductNotFound        = errors.New("product not found")
)

type ProductRepository interface {
	Create(product *product.Product) (err error)
	Update(product *product.Product) (err error)
	GetByID(id int) (product product.Product, err error)
	Delete(id int) (err error)
	GetAll() (products []product.Product, err error)
}
