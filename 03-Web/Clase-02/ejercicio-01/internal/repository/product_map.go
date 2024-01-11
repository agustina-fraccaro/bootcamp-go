package repository

import (
	"app/internal"
)

func NewProductMap(db map[int]internal.Product, lastId int) *ProductMap {
	return &ProductMap{
		db:     db,
		lastId: lastId,
	}
}

type ProductMap struct {
	db     map[int]internal.Product
	lastId int
}

func (p *ProductMap) Save(product *internal.Product) (err error) {
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
			return internal.ErrCodeValueAlreadyExists
		}
	}

	return
}

func (p *ProductMap) Update(product *internal.Product) (err error) {
	_, ok := p.db[(*product).Id]
	if !ok {
		err = internal.ErrProductNotFound
		return
	}

	p.db[(*product).Id] = *product
	return
}

func (m *ProductMap) GetByID(id int) (product internal.Product, err error) {
	product, ok := m.db[id]
	if !ok {
		err = internal.ErrProductNotFound
		return
	}

	return
}

func (m *ProductMap) Delete(id int) (err error) {
	_, ok := m.db[id]
	if !ok {
		err = internal.ErrProductNotFound
		return
	}

	delete(m.db, id)
	return
}
