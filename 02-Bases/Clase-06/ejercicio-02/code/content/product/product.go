package product

type Product interface {
	GetPrice() float64
}

type ProductSmall struct {
	Price float64
}

func (p ProductSmall) GetPrice() float64 {
	return p.Price
}

type ProductMedium struct {
	Price float64
}

func (p ProductMedium) GetPrice() float64 {
	return (p.Price + p.Price*0.03)
}

type ProductLarge struct {
	Price float64
}

func (p ProductLarge) GetPrice() float64 {
	return (p.Price + p.Price*0.06 + 2500)
}
