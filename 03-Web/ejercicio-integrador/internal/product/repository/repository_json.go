package repository

import (
	"app/internal/product"
	"app/internal/product/storage"
)

type ProductStore struct {
	st     storage.StorageProductJSON
	lastID int
}

func NewRepositoryProductStore(st storage.StorageProductJSON, lastId int) *ProductStore {
	return &ProductStore{
		st:     st,
		lastID: lastId,
	}
}

func (r *ProductStore) GetAll() (p []product.Product, err error) {
	p, err = r.st.ReadProducts()
	return
}

func (r *ProductStore) GetByID(id int) (p product.Product, err error) {
	ps, err := r.st.ReadProducts()
	if err != nil {
		return
	}

	var exists bool
	for _, v := range ps {
		if v.Id == id {
			p = v
			exists = true
			break
		}
	}

	if !exists {
		err = ErrProductNotFound
		return
	}

	return
}

func (r *ProductStore) Create(p *product.Product) (err error) {
	ps, err := r.st.ReadProducts()
	if err != nil {
		return
	}

	r.lastID++
	p.Id = r.lastID

	ps = append(ps, *p)

	err = r.st.WriteProducts(ps)
	if err != nil {
		return
	}

	return
}

func (r *ProductStore) Update(product *product.Product) (err error) {
	prods, err := r.st.ReadProducts()
	if err != nil {
		return
	}

	var exists bool
	for _, v := range prods {
		if v.Id == product.Id {
			exists = true
			break
		}
	}

	if !exists {
		err = ErrProductNotFound
		return
	}

	prods = append(prods, *product)

	err = r.st.WriteProducts(prods)
	if err != nil {
		return
	}

	return
}

func (r *ProductStore) Delete(id int) (err error) {
	ps, err := r.st.ReadProducts()
	if err != nil {
		return
	}

	var exists bool
	var ix int
	for k, v := range ps {
		if v.Id == id {
			ix = k
			exists = true
			break
		}
	}

	if !exists {
		err = ErrProductNotFound
		return
	}

	ps = append(ps[:ix], ps[ix+1:]...)

	err = r.st.WriteProducts(ps)
	if err != nil {
		return
	}

	return
}
