package internal

import (
	"errors"
)

var (
	ErrFieldRequired        = errors.New("field required")
	ErrFieldQuality         = errors.New("field quality")
	ErrProductAlreadyExists = errors.New("product already exists")
)

// MovieService is an interface that represents a movie service
// - business logic
// - validation
// - external services (e.g. apis, databases, etc.)
type ProductService interface {
	// Save saves a movie
	Save(product *Product) (err error)
}
