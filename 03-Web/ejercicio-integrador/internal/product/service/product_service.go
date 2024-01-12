package service

import (
	"app/internal/product"
	"errors"
)

var (
	ErrFieldRequired        = errors.New("field required")
	ErrFieldQuality         = errors.New("field quality")
	ErrProductAlreadyExists = errors.New("product already exists")
	ErrInvalidID            = errors.New("invalid id")
)

type ProductService interface {
	Create(product *product.Product) (err error)
	Update(product *product.Product) (err error)
	GetByID(id int) (product product.Product, err error)
	Delete(id int) (err error)
	GetAll() (products []product.Product, err error)
	GetConsumerPrice(ids []string) (sumPrice float64, products []product.Product, err error)
}
