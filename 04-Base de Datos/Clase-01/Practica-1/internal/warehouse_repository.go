package internal

import "errors"

var (
	ErrRepositoryWarehouseNotFound = errors.New("repository: warehouse not found")
)

type WarehouseRepository interface {
	FindAll() (ws []Warehouse, err error)
	FindById(id int) (w Warehouse, err error)
	Save(w *Warehouse) (err error)
}
