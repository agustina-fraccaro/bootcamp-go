package service

import (
	"app/internal/product"
	"app/internal/product/repository"
	"errors"
	"fmt"
	"strconv"
)

func NewProductDefault(rp repository.ProductRepository) *ProductDefault {
	return &ProductDefault{
		rp: rp,
	}
}

type ProductDefault struct {
	rp repository.ProductRepository
}

func (p *ProductDefault) Save(prod *product.Product) (err error) {
	if err = ValidateProduct(prod); err != nil {
		return
	}

	err = p.rp.Save(prod)
	if err != nil {
		switch err {
		case repository.ErrCodeValueAlreadyExists:
			err = fmt.Errorf("%w: code value", repository.ErrCodeValueAlreadyExists)
		}
		return
	}

	return
}

func (p *ProductDefault) GetAll() (products []product.Product, err error) {
	products, err = p.rp.GetAll()
	if err != nil {
		switch err {
		case repository.ErrCodeValueAlreadyExists:
			err = fmt.Errorf("%w: code value", repository.ErrCodeValueAlreadyExists)
		}
		return
	}

	return
}

func ValidateProduct(product *product.Product) (err error) {
	if (*product).Name == "" {
		return fmt.Errorf("%w: title", ErrFieldRequired)
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

func (p *ProductDefault) Update(product *product.Product) (err error) {
	if err = ValidateProduct(product); err != nil {
		return
	}

	err = p.rp.Update(product)
	if err != nil {
		switch err {
		case repository.ErrProductNotFound:
			err = fmt.Errorf("%w: id", repository.ErrProductNotFound)
		}
		return
	}
	return
}

func (m *ProductDefault) GetByID(id int) (product product.Product, err error) {
	product, err = m.rp.GetByID(id)
	if err != nil {
		switch err {
		case repository.ErrProductNotFound:
			err = fmt.Errorf("%w: id", repository.ErrProductNotFound)
		}
		return
	}

	return
}

func (p *ProductDefault) Delete(id int) (err error) {
	err = p.rp.Delete(id)
	if err != nil {
		switch err {
		case repository.ErrProductNotFound:
			err = fmt.Errorf("%w: id", repository.ErrProductNotFound)
		}
		return
	}

	return
}

func (p *ProductDefault) GetConsumerPrice(ids []string) (sumPrice float64, products []product.Product, err error) {
	var prod product.Product
	var id int

	for _, v := range ids {
		id, err = strconv.Atoi(v)
		if err != nil {
			err = fmt.Errorf("%w: id", ErrInvalidID)
			return
		}

		prod, err = p.rp.GetByID(id)
		if err != nil {
			switch err {
			case repository.ErrProductNotFound:
				err = fmt.Errorf("%w: id", repository.ErrProductNotFound)
			}
			return
		}

		if prod.IsPublished {
			sumPrice += prod.Price
			products = append(products, prod)
		}
	}

	cantProducts := len(ids)
	switch {
	case cantProducts < 10:
		sumPrice += (sumPrice * 1.21)
	case cantProducts >= 10 && cantProducts <= 20:
		sumPrice += (sumPrice * 1.17)
	case cantProducts > 20:
		sumPrice += (sumPrice * 1.15)
	}
	return
}
