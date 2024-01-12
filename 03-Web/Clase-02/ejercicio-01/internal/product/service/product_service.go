package service

import (
	"app/internal/product"
	"errors"
)

var (
	ErrFieldRequired        = errors.New("field required")
	ErrFieldQuality         = errors.New("field quality")
	ErrProductAlreadyExists = errors.New("product already exists")
)

type ProductService interface {
	Save(product *product.Product) (err error)
	Update(product *product.Product) (err error)
	GetByID(id int) (product product.Product, err error)
	Delete(id int) (err error)
	GetAll() (products []product.Product, err error)
}
