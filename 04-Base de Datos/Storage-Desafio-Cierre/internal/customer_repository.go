package internal

// RepositoryCustomer is the interface that wraps the basic methods that a customer repository should implement.
type RepositoryCustomer interface {
	// FindAll returns all customers saved in the database.
	FindAll() (c []Customer, err error)
	// Save saves a customer into the database.
	Save(c *Customer) (err error)
	// FindByCondition returns all customers invoices by condition saved in the database.
	FindByCondition() (c []CustomerByCondition, err error)
	// FindTop5ActiveCustomersMostSpend returns top 5 active customers who spend the most saved in the database.
	FindTop5ActiveCustomersMostSpend() (c []CustomerTop5ActiveCustomersMostSpend, err error)
}
