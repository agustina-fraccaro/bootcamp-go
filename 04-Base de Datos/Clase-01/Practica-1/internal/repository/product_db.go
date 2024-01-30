package repository

import (
	"app/internal"
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

func NewRepositoryProductDB(db *sql.DB) (r *RepositoryProductDB) {
	r = &RepositoryProductDB{
		db: db,
	}
	return
}

type RepositoryProductDB struct {
	// db is the underlying database.
	db *sql.DB
}

func (r *RepositoryProductDB) FindAll() (ps []internal.Product, err error) {
	rows, err := r.db.Query("SELECT id, name, price, quantity, code_value, is_published, expiration, id_warehouse FROM products")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var p internal.Product
		if err = rows.Scan(&p.Id, &p.Name, &p.Price, &p.Quantity, &p.CodeValue, &p.IsPublished, &p.Expiration, &p.WarehouseId); err != nil {
			return
		}

		ps = append(ps, p)
	}

	return

}

func (r *RepositoryProductDB) FindById(id int) (p internal.Product, err error) {
	row := r.db.QueryRow("SELECT id, name, price, quantity, code_value, is_published, expiration, id_warehouse FROM products WHERE id = ?", id)
	if err = row.Err(); err != nil {
		return
	}

	if err = row.Scan(&p.Id, &p.Name, &p.Price, &p.Quantity, &p.CodeValue, &p.IsPublished, &p.Expiration, &p.WarehouseId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = internal.ErrRepositoryProductNotFound
			return
		}
	}

	return
}

func (r *RepositoryProductDB) Save(p *internal.Product) (err error) {
	result, err := r.db.Exec("INSERT INTO products (name, price, quantity, code_value, is_published, expiration, id_warehouse) VALUES (?, ?, ?, ?, ?, ?, ?)", (*p).Name, (*p).Price, (*p).Quantity, (*p).CodeValue, (*p).IsPublished, (*p).Expiration, (*p).WarehouseId)

	if err != nil {
		return
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return
	}

	(*p).Id = int(lastInsertId)
	return
}

func (r *RepositoryProductDB) UpdateOrSave(p *internal.Product) (err error) {
	result, err := r.db.Exec("UPDATE products SET name = ?, price = ?, quantity = ?, code_value = ?, is_published = ?, expiration = ?, id_warehouse = ? WHERE id = ?", (*p).Name, (*p).Price, (*p).Quantity, (*p).CodeValue, (*p).IsPublished, (*p).Expiration, (*p).WarehouseId, (*p).Id)
	if err != nil {
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return
	}

	if rowsAffected == 0 {
		err = r.Save(p)
		if err != nil {
			return
		}
	}

	return
}

func (r *RepositoryProductDB) Update(p *internal.Product) (err error) {
	result, err := r.db.Exec("UPDATE product SET name = ?, price = ?, quantity = ?, code_value = ?, is_published = ?, expiration = ?, id_warehouse = ? WHERE id = ?", (*p).Name, (*p).Price, (*p).Quantity, (*p).CodeValue, (*p).IsPublished, (*p).Expiration, (*p).WarehouseId, (*p).Id)
	if err != nil {
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return
	}

	if rowsAffected == 0 {
		err = internal.ErrRepositoryProductNotFound
		return
	}

	return
}

func (r *RepositoryProductDB) Delete(id int) (err error) {
	_, err = r.db.Exec("DELETE FROM products WHERE id = ?", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = internal.ErrRepositoryProductNotFound
		}
		return
	}
	return
}

func (r *RepositoryProductDB) GetReport(query map[string]int) (p map[string]int, err error) {
	var rows *sql.Rows
	p = make(map[string]int)
	if query == nil {
		rows, err = r.db.Query("SELECT w.name, COUNT(p.id_warehouse) FROM products p INNER JOIN warehouses w ON p.id_warehouse = w.id GROUP BY p.id_warehouse")
		if err != nil {
			return
		}
	} else {
		rows, err = r.db.Query("SELECT w.name, COUNT(p.id_warehouse) FROM products p INNER JOIN warehouses w ON p.id_warehouse = w.id WHERE p.id_warehouse = ? GROUP BY p.id_warehouse", query["id"])
		if err != nil {
			return
		}
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		var count int
		if err = rows.Scan(&name, &count); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				err = internal.ErrRepositoryProductNotFound
				return
			}
		}

		p[name] = count
	}

	return
}
