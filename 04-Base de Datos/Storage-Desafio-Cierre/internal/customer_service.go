package internal

// ServiceCustomer is the interface that wraps the basic methods that a customer service should implement.
type ServiceCustomer interface {
	// FindAll returns all customers
	FindAll() (c []Customer, err error)
	// Save saves a customer
	Save(c *Customer) (err error)
	// FindByCondition returns all customers invoices by condition
	FindByCondition() (c []CustomerByCondition, err error)
	// FindTop5ActiveCustomersMostSpend returns top 5 active customers who spend the most
	FindTop5ActiveCustomersMostSpend() (c []CustomerTop5ActiveCustomersMostSpend, err error)
}
