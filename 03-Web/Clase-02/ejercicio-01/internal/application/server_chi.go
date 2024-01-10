package application

import (
	"app/internal"
	"app/internal/handlers"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func NewServerChi(address string) *ServerChi {
	// default config / values
	defaultAddress := ":8080"
	if address != "" {
		defaultAddress = address
	}

	return &ServerChi{
		address: defaultAddress,
	}
}

type ServerChi struct {
	address string
}

func (s *ServerChi) Run() error {
	products, err := getProductsFromJSON()
	if err != nil {
		return err
	}
	lastID := findMaxID(products)
	hd := handlers.NewDefaultProducts(products, lastID)
	rt := chi.NewRouter()

	rt.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	rt.Route("/products", func(r chi.Router) {
		r.Get("/", hd.Get())

		r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
			id := chi.URLParam(r, "id")
			idInt, err := strconv.Atoi(id)
			if err != nil {
				w.Header().Set("Content-Type", "text/plain")
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("invalid id"))
				return
			}

			product := hd.GetById(idInt)
			if err != nil {
				w.Header().Set("Content-Type", "text/plain")
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("product not found"))
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(product)
		})

		r.Get("/search", hd.GetByPrice())
	})
	rt.Post("/products", hd.Create())

	return http.ListenAndServe(s.address, rt)
}

func getProductsFromJSON() ([]internal.Product, error) {
	jsonFile, err := os.Open("products.json")
	if err != nil {
		fmt.Println("could not open JSON:", err)
		return nil, err
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)

	var products []internal.Product
	if err := decoder.Decode(&products); err != nil {
		fmt.Println("could not decode JSON:", err)
		return nil, err
	}

	return products, nil
}

func findMaxID(products []internal.Product) int {
	var maxID int = 0

	for id := range products {
		if id > maxID {
			maxID = id
		}
	}

	return maxID
}
