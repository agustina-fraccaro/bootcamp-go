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

func (r *RepositoryProductDB) FindById(id int) (p internal.Product, err error) {
	row := r.db.QueryRow("SELECT id, name, price, quantity, code_value, is_published, expiration FROM product WHERE id = ?", id)
	if err = row.Err(); err != nil {
		return
	}

	if err = row.Scan(&p.Id, &p.Name, &p.Quantity, &p.Price, &p.CodeValue, &p.Expiration, &p.IsPublished); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return
		}
	}

	return
}

func (r *RepositoryProductDB) Save(p *internal.Product) (err error) {
	result, err := r.db.Exec("INSERT INTO product (name, price, quantity, code_value, is_published, expiration) VALUES (?, ?, ?, ?, ?, ?)", (*p).Name, (*p).Price, (*p).Quantity, (*p).CodeValue, (*p).IsPublished, (*p).Expiration)

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
	result, err := r.db.Exec("UPDATE product SET name = ?, price = ?, quantity = ?, code_value = ?, is_published = ?, expiration = ? WHERE id = ?", (*p).Name, (*p).Price, (*p).Quantity, (*p).CodeValue, (*p).IsPublished, (*p).Expiration, (*p).Id)
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
	result, err := r.db.Exec("UPDATE product SET name = ?, price = ?, quantity = ?, code_value = ?, is_published = ?, expiration = ? WHERE id = ?", (*p).Name, (*p).Price, (*p).Quantity, (*p).CodeValue, (*p).IsPublished, (*p).Expiration, (*p).Id)
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
	_, err = r.db.Exec("DELETE FROM product WHERE id = ?", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = internal.ErrRepositoryProductNotFound
		}
		return
	}
	return
}
