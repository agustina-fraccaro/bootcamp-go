package handler

import (
	"app/internal"
	"app/platform/web/request"
	"app/platform/web/response"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func NewDefaultProducts(sv internal.ProductService) *DefaultProducts {
	return &DefaultProducts{
		sv: sv,
	}
}

type DefaultProducts struct {
	sv internal.ProductService
}

type ProductJSON struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

type BodyRequestProductJSON struct {
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func (d *DefaultProducts) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid id")
			return
		}

		product, err := d.sv.GetByID(id)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrProductNotFound):
				response.Text(w, http.StatusNotFound, "product not found")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		data := ProductJSON{
			ID:          product.Id,
			Name:        product.Name,
			Quantity:    product.Quantity,
			CodeValue:   product.CodeValue,
			IsPublished: product.IsPublished,
			Expiration:  product.Expiration,
			Price:       product.Price,
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "product found",
			"data":    data,
		})
	}
}

func (d *DefaultProducts) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body BodyRequestProductJSON
		if err := request.JSON(r, &body); err != nil {
			response.Text(w, http.StatusBadRequest, "invalid body")
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

		if err := d.sv.Save(&product); err != nil {
			switch {
			case errors.Is(err, internal.ErrFieldRequired), errors.Is(err, internal.ErrFieldQuality):
				response.Text(w, http.StatusBadRequest, "invalid body")
			case errors.Is(err, internal.ErrCodeValueAlreadyExists):
				response.Text(w, http.StatusConflict, "product already exists")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		data := ProductJSON{
			ID:          product.Id,
			Name:        product.Name,
			Quantity:    product.Quantity,
			CodeValue:   product.CodeValue,
			IsPublished: product.IsPublished,
			Expiration:  product.Expiration,
			Price:       product.Price,
		}
		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "product created",
			"data":    data,
		})
	}
}

func (d *DefaultProducts) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid id")
			return
		}

		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid body")
			return
		}

		var bodyMap map[string]any
		if err := json.Unmarshal(bytes, &bodyMap); err != nil {
			response.Text(w, http.StatusBadRequest, "invalid body")
			return
		}

		if err := ValidateKeyExistante(bodyMap, "name", "quantity", "code_value", "is_published", "expiration", "price"); err != nil {
			response.Text(w, http.StatusBadRequest, "invalid body")
			return
		}

		var body BodyRequestProductJSON
		if err := json.Unmarshal(bytes, &body); err != nil {
			response.Text(w, http.StatusBadRequest, "invalid body")
			return
		}

		product := internal.Product{
			Id:          id,
			Name:        body.Name,
			Quantity:    body.Quantity,
			CodeValue:   body.CodeValue,
			IsPublished: body.IsPublished,
			Expiration:  body.Expiration,
			Price:       body.Price,
		}

		if err := d.sv.Update(&product); err != nil {
			switch {
			case errors.Is(err, internal.ErrProductNotFound):
				response.Text(w, http.StatusNotFound, "product not found")
			case errors.Is(err, internal.ErrFieldRequired), errors.Is(err, internal.ErrFieldQuality):
				response.Text(w, http.StatusBadRequest, "invalid body")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		data := ProductJSON{
			ID:          product.Id,
			Name:        product.Name,
			Quantity:    product.Quantity,
			CodeValue:   product.CodeValue,
			IsPublished: product.IsPublished,
			Expiration:  product.Expiration,
			Price:       product.Price,
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "product updated",
			"data":    data,
		})
	}
}

func (d *DefaultProducts) UpdatePartial() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid id")
			return
		}

		product, err := d.sv.GetByID(id)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrProductNotFound):
				response.Text(w, http.StatusNotFound, "product not found")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		reqBody := BodyRequestProductJSON{
			Name:        product.Name,
			Quantity:    product.Quantity,
			CodeValue:   product.CodeValue,
			IsPublished: product.IsPublished,
			Expiration:  product.Expiration,
			Price:       product.Price,
		}

		if err := request.JSON(r, &reqBody); err != nil {
			response.Text(w, http.StatusBadRequest, "invalid body")
			return
		}

		product = internal.Product{
			Id:          id,
			Name:        reqBody.Name,
			Quantity:    reqBody.Quantity,
			CodeValue:   reqBody.CodeValue,
			IsPublished: reqBody.IsPublished,
			Expiration:  reqBody.Expiration,
			Price:       reqBody.Price,
		}

		if err := d.sv.Update(&product); err != nil {
			switch {
			case errors.Is(err, internal.ErrProductNotFound):
				response.Text(w, http.StatusNotFound, "product not found")
			case errors.Is(err, internal.ErrFieldRequired), errors.Is(err, internal.ErrFieldQuality):
				response.Text(w, http.StatusBadRequest, "invalid body")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		data := ProductJSON{
			ID:          id,
			Name:        reqBody.Name,
			Quantity:    reqBody.Quantity,
			CodeValue:   reqBody.CodeValue,
			IsPublished: reqBody.IsPublished,
			Expiration:  reqBody.Expiration,
			Price:       reqBody.Price,
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "product updated",
			"data":    data,
		})

	}
}

func (d *DefaultProducts) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid id")
			return
		}

		if err := d.sv.Delete(id); err != nil {
			switch {
			case errors.Is(err, internal.ErrProductNotFound):
				response.Text(w, http.StatusNotFound, "product not found")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		response.Text(w, http.StatusNoContent, "movie deleted")
	}
}

func ValidateKeyExistante(mp map[string]any, keys ...string) (err error) {
	for _, k := range keys {
		if _, ok := mp[k]; !ok {
			return fmt.Errorf("key %s not found", k)
		}
	}
	return
}
