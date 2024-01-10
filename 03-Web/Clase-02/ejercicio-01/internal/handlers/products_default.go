package handlers

import (
	"app/internal"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func NewDefaultProducts(products []internal.Product, lastId int) *DefaultProducts {
	return &DefaultProducts{
		products: products,
		lastID:   lastId,
	}
}

type DefaultProducts struct {
	products []internal.Product
	lastID   int
}

type BodyRequestProductJSON struct {
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func (d *DefaultProducts) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var body BodyRequestProductJSON
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid body"))
			return
		}

		product := internal.Product{
			Name:        body.Name,
			Quantity:    body.Quantity,
			CodeValue:   body.CodeValue,
			IsPublished: body.IsPublished,
			Expiration:  body.Expiration,
			Price:       body.Price,
		}

		(*d).lastID++

		product.Id = (*d).lastID

		if err := ValidateBussinessRule(&product); err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid product"))
			return
		}

		for _, mv := range (*d).products {
			if mv.CodeValue == product.CodeValue {
				w.Header().Set("Content-Type", "text/plain")
				w.WriteHeader(http.StatusConflict)
				w.Write([]byte("product already exists"))
				return
			}
		}

		(*d).products = append((*d).products, product)

		data := internal.Product{
			Id:          product.Id,
			Name:        product.Name,
			Quantity:    product.Quantity,
			CodeValue:   product.CodeValue,
			IsPublished: product.IsPublished,
			Expiration:  product.Expiration,
			Price:       product.Price,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]any{
			"message": "movie created",
			"data":    data,
		})
	}
}

func ValidateBussinessRule(prod *internal.Product) error {
	if prod.Name == "" {
		return errors.New("name is required")
	}
	if prod.Quantity == 0 {
		return errors.New("quantity is required")
	}

	if prod.Price == 0 {
		return errors.New("price is required")
	}

	if prod.Expiration == "" {
		return errors.New("expiration is required")
	}

	if prod.CodeValue == "" {
		return errors.New("code value is required")
	}

	day, err := strconv.Atoi(prod.Expiration[0:2])
	if err != nil {
		return errors.New("day is invalid")
	}
	month, err := strconv.Atoi(prod.Expiration[3:5])
	if err != nil {
		return errors.New("month is invalid")
	}
	year, err := strconv.Atoi(prod.Expiration[6:10])
	if err != nil {
		return errors.New("year is invalid")
	}

	if day < 1 || day > 31 {
		return errors.New("day is invalid")
	}

	if month < 1 || month > 12 {
		return errors.New("month is invalid")
	}

	if year < 2021 || year > 2025 {
		return errors.New("year is invalid")
	}

	return nil
}

func (d *DefaultProducts) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(d.products)
	}
}

func (d *DefaultProducts) GetById(id int) internal.Product {
	var productFound internal.Product
	for _, product := range (*d).products {
		if product.Id == id {
			productFound = product
		}
	}

	return productFound
}

func (d *DefaultProducts) GetByPrice() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		price := r.URL.Query().Get("priceGt")
		if price == "" {
			http.Error(w, "could not found param", http.StatusBadRequest)
			return
		}

		priceGt, err := strconv.ParseFloat(price, 64)
		if err != nil {
			http.Error(w, "invalid price", http.StatusBadRequest)
			return
		}
		resultados := []internal.Product{}
		for _, producto := range d.products {
			if producto.Price > priceGt {
				resultados = append(resultados, producto)
			}
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resultados)
	}
}
