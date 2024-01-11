package internal

import (
	"errors"
)

var (
	ErrFieldRequired        = errors.New("field required")
	ErrFieldQuality         = errors.New("field quality")
	ErrProductAlreadyExists = errors.New("product already exists")
)

type ProductService interface {
	Save(product *Product) (err error)
	Update(product *Product) (err error)
	GetByID(id int) (product Product, err error)
	Delete(id int) (err error)
}
