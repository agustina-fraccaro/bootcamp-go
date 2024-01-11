package service

import (
	"app/internal"
	"errors"
	"fmt"
	"strconv"
)

// NewMovieDefault creates a new instance of a movie service
func NewProductDefault(rp *internal.ProductRepository) *ProductDefault {
	return &ProductDefault{
		rp: rp,
	}
}

type ProductDefault struct {
	rp internal.ProductRepository
}

func (p *ProductDefault) Save(product *internal.Product) (err error) {
	if (*product).Name == "" {
		return fmt.Errorf("%w: title", internal.ErrFieldRequired)
	}
	if (*product).Quantity == 0 {
		return errors.New("quantity is required")
	}

	if (*product).Price == 0 {
		return errors.New("price is required")
	}

	if (*product).Expiration == "" {
		return errors.New("expiration is required")
	}

	if (*product).CodeValue == "" {
		return errors.New("code value is required")
	}

	day, err := strconv.Atoi((*product).Expiration[0:2])
	if err != nil {
		return errors.New("day is invalid")
	}
	month, err := strconv.Atoi((*product).Expiration[3:5])
	if err != nil {
		return errors.New("month is invalid")
	}
	year, err := strconv.Atoi((*product).Expiration[6:10])
	if err != nil {
		return errors.New("year is invalid")
	}

	if day < 1 || day > 31 {
		return errors.New("day is invalid")
	}

	if month < 1 || month > 12 {
		return errors.New("month is invalid")
	}

	if year < 2021 || year > 2025 {
		return errors.New("year is invalid")
	}
	err = p.rp.Save(product)
	if err != nil {
		switch err {
		case internal.ErrCodeValueAlreadyExists:
			err = fmt.Errorf("%w: title", internal.ErrCodeValueAlreadyExists)
		}
		return
	}

	return
}
