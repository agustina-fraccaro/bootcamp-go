package internal

// CustomerAttributes is the struct that represents the attributes of a customer.
type CustomerAttributes struct {
	// FirstName is the first name of the customer.
	FirstName string
	// LastName is the last name of the customer.
	LastName string
	// Condition is the condition of the customer.
	Condition int
}

// Customer is the struct that represents a customer.
type Customer struct {
	// Id is the unique identifier of the customer.
	Id int
	// CustomerAttributes is the attributes of the customer.
	CustomerAttributes
}

// CustomerInvoicesByCondition is the struct that represents the total invoices by customer condition.
type CustomerByCondition struct {
	// Condition is the condition of the customer.
	Condition int
	// Total is the total invoices by customer condition
	Total float64
}

type CustomerTop5ActiveCustomersMostSpend struct {
	FirstName string
	LastName  string
	Amount    float64
}
