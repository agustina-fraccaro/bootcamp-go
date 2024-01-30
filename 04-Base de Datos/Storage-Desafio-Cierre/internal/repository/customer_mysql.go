package repository

import (
	"database/sql"

	"app/internal"
)

// NewCustomersMySQL creates new mysql repository for customer entity.
func NewCustomersMySQL(db *sql.DB) *CustomersMySQL {
	return &CustomersMySQL{db}
}

// CustomersMySQL is the MySQL repository implementation for customer entity.
type CustomersMySQL struct {
	// db is the database connection.
	db *sql.DB
}

// FindAll returns all customers from the database.
func (r *CustomersMySQL) FindAll() (c []internal.Customer, err error) {
	// execute the query
	rows, err := r.db.Query("SELECT `id`, `first_name`, `last_name`, `condition` FROM customers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var cs internal.Customer
		// scan the row into the customer
		err := rows.Scan(&cs.Id, &cs.FirstName, &cs.LastName, &cs.Condition)
		if err != nil {
			return nil, err
		}
		// append the customer to the slice
		c = append(c, cs)
	}
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

// Save saves the customer into the database.
func (r *CustomersMySQL) Save(c *internal.Customer) (err error) {
	// execute the query
	res, err := r.db.Exec(
		"INSERT INTO customers (`first_name`, `last_name`, `condition`) VALUES (?, ?, ?)",
		(*c).FirstName, (*c).LastName, (*c).Condition,
	)
	if err != nil {
		return err
	}

	// get the last inserted id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set the id
	(*c).Id = int(id)

	return
}

func (r *CustomersMySQL) FindByCondition() (c []internal.CustomerByCondition, err error) {
	result, err := r.db.Query("SELECT c.condition, ROUND(SUM(i.total), 2) as total FROM CUSTOMERS c INNER JOIN INVOICES i ON c.id = i.customer_id GROUP BY c.condition")
	if err != nil {
		return
	}
	defer result.Close()
	for result.Next() {
		var cs internal.CustomerByCondition
		err := result.Scan(&cs.Condition, &cs.Total)
		if err != nil {
			return nil, err
		}
		c = append(c, cs)
	}
	err = result.Err()
	if err != nil {
		return
	}
	return
}

func (r *CustomersMySQL) FindTop5ActiveCustomersMostSpend() (c []internal.CustomerTop5ActiveCustomersMostSpend, err error) {
	result, err := r.db.Query("SELECT c.first_name, c.last_name, ROUND(SUM(i.total), 2) as total FROM CUSTOMERS c INNER JOIN INVOICES i ON c.id = i.customer_id WHERE c.condition = 1 GROUP BY c.id ORDER BY total DESC LIMIT 5")
	if err != nil {
		return
	}

	defer result.Close()

	for result.Next() {
		var cs internal.CustomerTop5ActiveCustomersMostSpend
		err := result.Scan(&cs.FirstName, &cs.LastName, &cs.Amount)
		if err != nil {
			return nil, err
		}
		c = append(c, cs)
	}

	return
}
