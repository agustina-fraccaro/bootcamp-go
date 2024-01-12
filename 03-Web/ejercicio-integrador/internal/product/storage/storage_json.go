package storage

import (
	"app/internal/product"
	"encoding/json"
	"os"
)

type StorageProductJSON struct {
	filePath string
}

func NewStorageProductJSON(filePath string) *StorageProductJSON {

	return &StorageProductJSON{filePath: filePath}
}

type ProductAttributesJSON struct {
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func (s *StorageProductJSON) ReadProducts() (p []product.Product, err error) {
	f, err := os.Open(s.filePath)
	if err != nil {
		return
	}
	defer f.Close()

	pr := make(map[int]ProductAttributesJSON)
	err = json.NewDecoder(f).Decode(&pr)
	if err != nil {
		return
	}

	for k, v := range pr {
		p = append(p, *product.NewProduct(k, v.Name, v.Quantity, v.CodeValue, v.IsPublished, v.Expiration, v.Price))
	}

	return
}

func (s *StorageProductJSON) WriteProducts(p []product.Product) (err error) {
	f, err := os.OpenFile(s.filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return
	}

	defer f.Close()

	pr := make(map[int]ProductAttributesJSON)
	for _, v := range p {
		pr[v.Id] = ProductAttributesJSON{v.Name, v.Quantity, v.CodeValue, v.IsPublished, v.Expiration, v.Price}
	}

	err = json.NewEncoder(f).Encode(pr)
	if err != nil {
		return
	}

	return
}
