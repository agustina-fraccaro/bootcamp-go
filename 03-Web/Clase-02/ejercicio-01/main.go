package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi"
)

type producto struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func main() {
	jsonFile, err := os.Open("products.json")
	if err != nil {
		fmt.Println("Error al abrir el archivo JSON:", err)
		return
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)

	var productos []producto

	err = decoder.Decode(&productos)
	if err != nil {
		fmt.Println("Error al decodificar el archivo JSON:", err)
		return
	}

	router := chi.NewRouter()

	http.ListenAndServe(":8080", router)

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	router.Route("/products", func(r chi.Router) {
		router.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(productos)
		})

		router.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
			id, err := strconv.Atoi(chi.URLParam(r, "id"))
			if err != nil {
				http.Error(w, "Error al parsear 'id'", http.StatusBadRequest)
				return
			}
			for _, producto := range productos {
				if producto.Id == id {
					w.WriteHeader(http.StatusOK)
					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(producto)
					return
				}
			}
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Producto no encontrado"))
		})

		router.Get("/search", func(w http.ResponseWriter, r *http.Request) {
			price := r.URL.Query().Get("priceGt")
			if price == "" {
				http.Error(w, "Parámetro 'priceGt' no proporcionado", http.StatusBadRequest)
				return
			}

			priceGt, err := strconv.ParseFloat(price, 64)
			if err != nil {
				http.Error(w, "Error al parsear 'priceGt'", http.StatusBadRequest)
				return
			}
			resultados := []producto{}
			for _, producto := range productos {
				if producto.Price > priceGt {
					resultados = append(resultados, producto)
				}
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resultados)
		})
	})

}
