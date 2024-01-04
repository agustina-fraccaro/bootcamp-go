package product

const (
	largeType  = "large"
	mediumType = "medium"
	smallType  = "small"
)

func NewProduct(productType string) Product {
	switch productType {
	case largeType:
		return ProductLarge{
			Price: 100,
		}
	case mediumType:
		return ProductMedium{
			Price: 50,
		}
	case smallType:
		return ProductSmall{
			Price: 25,
		}
	default:
		return nil
	}
}
