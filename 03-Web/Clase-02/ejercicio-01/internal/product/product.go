package product

func NewProduct(id int, name string, quantity int, codeValue string, isPublished bool, expiration string, price float64) *Product {
	return &Product{id, name, quantity, codeValue, isPublished, expiration, price}
}

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}
