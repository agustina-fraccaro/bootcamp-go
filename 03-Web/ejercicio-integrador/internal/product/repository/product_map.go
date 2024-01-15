package repository

import (
	"app/internal/product"
)

func NewProductMap(db map[int]product.Product, lastId int) *ProductMap {
	return &ProductMap{
		db:     db,
		lastId: lastId,
	}
}

type ProductMap struct {
	db     map[int]product.Product
	lastId int
}

func (p *ProductMap) Create(product *product.Product) (err error) {
	if err = p.ValidateProductCodeValue((*product).CodeValue); err != nil {
		return
	}

	(*p).lastId++
	(*product).Id = (*p).lastId

	(*p).db[(*product).Id] = *product

	return
}

func (p *ProductMap) ValidateProductCodeValue(codeValue string) (err error) {
	for _, v := range (*p).db {
		if v.CodeValue == codeValue {
			return ErrCodeValueAlreadyExists
		}
	}

	return
}

func (p *ProductMap) Update(product *product.Product) (err error) {
	_, ok := p.db[(*product).Id]
	if !ok {
		err = ErrProductNotFound
		return
	}

	p.db[(*product).Id] = *product
	return
}

func (m *ProductMap) GetByID(id int) (product product.Product, err error) {
	product, ok := m.db[id]
	if !ok {
		err = ErrProductNotFound
		return
	}

	return
}

func (m *ProductMap) Delete(id int) (err error) {
	_, ok := m.db[id]
	if !ok {
		err = ErrProductNotFound
		return
	}

	delete(m.db, id)
	return
}

func (m *ProductMap) GetAll() (products []product.Product, err error) {
	for _, v := range m.db {
		products = append(products, v)
	}

	return
}
