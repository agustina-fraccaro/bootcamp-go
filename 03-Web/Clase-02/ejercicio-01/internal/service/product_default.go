package service

import (
	"app/internal"
	"errors"
	"fmt"
	"strconv"
)

func NewProductDefault(rp internal.ProductRepository) *ProductDefault {
	return &ProductDefault{
		rp: rp,
	}
}

type ProductDefault struct {
	rp internal.ProductRepository
}

func (p *ProductDefault) Save(product *internal.Product) (err error) {
	if err = ValidateProduct(product); err != nil {
		return
	}

	err = p.rp.Save(product)
	if err != nil {
		switch err {
		case internal.ErrCodeValueAlreadyExists:
			err = fmt.Errorf("%w: code value", internal.ErrCodeValueAlreadyExists)
		}
		return
	}

	return
}

func (p *ProductDefault) GetAll() (products []internal.Product, err error) {
	products, err = p.rp.GetAll()
	if err != nil {
		switch err {
		case internal.ErrCodeValueAlreadyExists:
			err = fmt.Errorf("%w: code value", internal.ErrCodeValueAlreadyExists)
		}
		return
	}

	return
}

func ValidateProduct(product *internal.Product) (err error) {
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

	return
}

func (p *ProductDefault) Update(product *internal.Product) (err error) {
	if err = ValidateProduct(product); err != nil {
		return
	}

	err = p.rp.Update(product)
	if err != nil {
		switch err {
		case internal.ErrProductNotFound:
			err = fmt.Errorf("%w: id", internal.ErrProductNotFound)
		}
		return
	}
	return
}

func (m *ProductDefault) GetByID(id int) (product internal.Product, err error) {
	product, err = m.rp.GetByID(id)
	if err != nil {
		switch err {
		case internal.ErrProductNotFound:
			err = fmt.Errorf("%w: id", internal.ErrProductNotFound)
		}
		return
	}

	return
}

func (m *ProductDefault) Delete(id int) (err error) {
	err = m.rp.Delete(id)
	if err != nil {
		switch err {
		case internal.ErrProductNotFound:
			err = fmt.Errorf("%w: id", internal.ErrProductNotFound)
		}
		return
	}

	return
}
