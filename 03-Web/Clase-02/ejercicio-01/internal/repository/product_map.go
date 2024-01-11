package repository

import "app/internal"

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
	for _, v := range (*p).db {
		if v.CodeValue == product.CodeValue {
			return internal.ErrCodeValueAlreadyExists
		}
	}

	(*p).lastId++
	(*product).Id = (*p).lastId

	(*p).db[(*product).Id] = *product

	return
}
