package internal

import "errors"

var (
	ErrCodeValueAlreadyExists = errors.New("code value already exists")
	ErrProductNotFound        = errors.New("product not found")
	ErrNoProductsFound        = errors.New("no products found")
)

type ProductRepository interface {
	Save(product *Product) (err error)
	Update(product *Product) (err error)
	GetByID(id int) (product Product, err error)
	Delete(id int) (err error)
	GetAll() (products []Product, err error)
}
