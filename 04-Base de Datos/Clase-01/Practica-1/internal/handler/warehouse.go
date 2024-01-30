package handler

import (
	"app/internal"
	"app/platform/web/request"
	"app/platform/web/response"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func NewHandlerWarehouse(rp internal.WarehouseRepository) (h *HandlerWarehouse) {
	h = &HandlerWarehouse{
		rp: rp,
	}
	return
}

type HandlerWarehouse struct {
	rp internal.WarehouseRepository
}

type WarehouseJSON struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	Telephone string `json:"telephone"`
	Capacity  int    `json:"capacity"`
}

type RequestBodyWarehouse struct {
	Name      string `json:"name"`
	Address   string `json:"address"`
	Telephone string `json:"telephone"`
	Capacity  int    `json:"capacity"`
}

func (h *HandlerWarehouse) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "invalid id")
			return
		}

		warehouse, err := h.rp.FindById(id)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrRepositoryWarehouseNotFound):
				response.JSON(w, http.StatusNotFound, "warehouse not found")
			default:
				response.JSON(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		data := WarehouseJSON{
			Id:        warehouse.Id,
			Name:      warehouse.Name,
			Address:   warehouse.Address,
			Telephone: warehouse.Telephone,
			Capacity:  warehouse.Capacity,
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (h *HandlerWarehouse) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body RequestBodyWarehouse
		err := request.JSON(r, &body)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "invalid body")
			return
		}

		warehouse := internal.Warehouse{
			Name:      body.Name,
			Address:   body.Address,
			Telephone: body.Telephone,
			Capacity:  body.Capacity,
		}

		err = h.rp.Save(&warehouse)
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, "internal server error")
			return
		}

		data := WarehouseJSON{
			Id:        warehouse.Id,
			Name:      warehouse.Name,
			Address:   warehouse.Address,
			Telephone: warehouse.Telephone,
			Capacity:  warehouse.Capacity,
		}

		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (h *HandlerWarehouse) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		warehouses, err := h.rp.FindAll()
		if err != nil {
			fmt.Println(err)
			response.JSON(w, http.StatusInternalServerError, "internal server error")
			return
		}

		var data []WarehouseJSON
		for _, warehouse := range warehouses {
			data = append(data, WarehouseJSON{
				Id:        warehouse.Id,
				Name:      warehouse.Name,
				Address:   warehouse.Address,
				Telephone: warehouse.Telephone,
				Capacity:  warehouse.Capacity,
			})
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}
