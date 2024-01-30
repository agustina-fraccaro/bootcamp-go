package repository

import (
	"app/internal"
	"database/sql"
	"errors"
)

func NewRepositoryWarehouseDB(db *sql.DB) (r *RepositoryWarehouseDB) {
	return &RepositoryWarehouseDB{
		db: db,
	}
}

type RepositoryWarehouseDB struct {
	db *sql.DB
}

func (r *RepositoryWarehouseDB) FindAll() (ws []internal.Warehouse, err error) {
	rows, err := r.db.Query("SELECT id, name, adress, telephone, capacity FROM warehouses")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var w internal.Warehouse
		if err = rows.Scan(&w.Id, &w.Name, &w.Address, &w.Telephone, &w.Capacity); err != nil {
			return
		}

		ws = append(ws, w)
	}

	return
}

func (r *RepositoryWarehouseDB) FindById(id int) (w internal.Warehouse, err error) {
	row := r.db.QueryRow("SELECT id, name, adress, telephone, capacity FROM warehouses WHERE id = ?", id)
	if err = row.Err(); err != nil {
		return
	}

	if err = row.Scan(&w.Id, &w.Name, &w.Address, &w.Telephone, &w.Capacity); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = internal.ErrRepositoryWarehouseNotFound
			return
		}
	}

	return
}

func (r *RepositoryWarehouseDB) Save(w *internal.Warehouse) (err error) {
	result, err := r.db.Exec("INSERT INTO warehouses (name, adress, telephone, capacity) VALUES (?, ?, ?, ?)", (*w).Name, (*w).Address, (*w).Telephone, (*w).Capacity)

	if err != nil {
		return
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return
	}

	(*w).Id = int(lastInsertId)
	return
}
