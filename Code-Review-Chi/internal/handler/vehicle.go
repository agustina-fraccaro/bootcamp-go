package handler

import (
	"app/internal"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

// VehicleJSON is a struct that represents a vehicle in JSON format
type VehicleJSON struct {
	ID              int     `json:"id"`
	Brand           string  `json:"brand"`
	Model           string  `json:"model"`
	Registration    string  `json:"registration"`
	Color           string  `json:"color"`
	FabricationYear int     `json:"year"`
	Capacity        int     `json:"passengers"`
	MaxSpeed        float64 `json:"max_speed"`
	FuelType        string  `json:"fuel_type"`
	Transmission    string  `json:"transmission"`
	Weight          float64 `json:"weight"`
	Height          float64 `json:"height"`
	Length          float64 `json:"length"`
	Width           float64 `json:"width"`
}

type BodyRequestVehicleJson struct {
	ID              int     `json:"id"`
	Brand           string  `json:"brand"`
	Model           string  `json:"model"`
	Registration    string  `json:"registration"`
	Color           string  `json:"color"`
	FabricationYear int     `json:"year"`
	Capacity        int     `json:"passengers"`
	MaxSpeed        float64 `json:"max_speed"`
	FuelType        string  `json:"fuel_type"`
	Transmission    string  `json:"transmission"`
	Weight          float64 `json:"weight"`
	Height          float64 `json:"height"`
	Length          float64 `json:"length"`
	Width           float64 `json:"width"`
}

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(sv internal.VehicleService) *VehicleDefault {
	return &VehicleDefault{sv: sv}
}

// VehicleDefault is a struct with methods that represent handlers for vehicles
type VehicleDefault struct {
	// sv is the service that will be used by the handler
	sv internal.VehicleService
}

// GetAll is a method that returns a handler for the route GET /vehicles
func (h *VehicleDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		// - get all vehicles
		v, err := h.sv.FindAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

// Create is a method that returns a handler for the route POST /vehicles
func (h *VehicleDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - parse body
		var body BodyRequestVehicleJson
		if err := request.JSON(r, &body); err != nil {
			response.Text(w, http.StatusBadRequest, "invalid body")
			return
		}

		vehicle := internal.Vehicle{
			Id: body.ID,
			VehicleAttributes: internal.VehicleAttributes{
				Brand:           body.Brand,
				Model:           body.Model,
				Registration:    body.Registration,
				Color:           body.Color,
				FabricationYear: body.FabricationYear,
				Capacity:        body.Capacity,
				MaxSpeed:        body.MaxSpeed,
				FuelType:        body.FuelType,
				Transmission:    body.Transmission,
				Weight:          body.Weight,
				Dimensions: internal.Dimensions{
					Height: body.Height,
					Length: body.Length,
					Width:  body.Width,
				},
			},
		}

		// process
		// - create vehicle
		if err := h.sv.Create(&vehicle); err != nil {
			switch {
			case errors.Is(err, internal.ErrFieldRequired):
				response.Text(w, http.StatusBadRequest, "invalid body")
			case errors.Is(err, internal.ErrVehicleIdAlreadyExists):
				response.Text(w, http.StatusConflict, "vehicle id already exists")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		data := VehicleJSON{
			ID:              vehicle.Id,
			Brand:           vehicle.Brand,
			Model:           vehicle.Model,
			Registration:    vehicle.Registration,
			Color:           vehicle.Color,
			FabricationYear: vehicle.FabricationYear,
			Capacity:        vehicle.Capacity,
			MaxSpeed:        vehicle.MaxSpeed,
			FuelType:        vehicle.FuelType,
			Transmission:    vehicle.Transmission,
			Weight:          vehicle.Weight,
			Height:          vehicle.VehicleAttributes.Dimensions.Height,
			Length:          vehicle.VehicleAttributes.Dimensions.Length,
			Width:           vehicle.VehicleAttributes.Dimensions.Width,
		}

		// response
		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "vehicle created",
			"data":    data,
		})
	}
}

// GetByColorAndYear is a method that returns a handler for the route GET /vehicles/color/{color}/year/{year}
func (h *VehicleDefault) GetByColorAndYear() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - parse path params
		color := chi.URLParam(r, "color")
		year, err := strconv.Atoi(chi.URLParam(r, "year"))
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid year")
			return
		}

		if color == "" {
			response.Text(w, http.StatusBadRequest, "invalid color")
			return
		}

		// process
		// - get vehicles by color and year
		v, err := h.sv.FindByColorAndYear(color, year)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrVehicleNotFound):
				response.Text(w, http.StatusNotFound, "vehicle not found")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.VehicleAttributes.Brand,
				Model:           value.VehicleAttributes.Model,
				Registration:    value.VehicleAttributes.Registration,
				Color:           value.VehicleAttributes.Color,
				FabricationYear: value.VehicleAttributes.FabricationYear,
				Capacity:        value.VehicleAttributes.Capacity,
				MaxSpeed:        value.VehicleAttributes.MaxSpeed,
				FuelType:        value.VehicleAttributes.FuelType,
				Transmission:    value.VehicleAttributes.Transmission,
				Weight:          value.VehicleAttributes.Weight,
				Height:          value.VehicleAttributes.Dimensions.Height,
				Length:          value.VehicleAttributes.Dimensions.Length,
				Width:           value.VehicleAttributes.Dimensions.Width,
			}
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

// GetByBrandAndYears is a method that returns a handler for the route GET /vehicles/brand/{brand}/between/{start_year}/{end_year}
func (h *VehicleDefault) GetByBrandAndYears() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		brand := chi.URLParam(r, "brand")
		if brand == "" {
			response.Text(w, http.StatusBadRequest, "invalid brand")
			return
		}
		startYear, err := strconv.Atoi(chi.URLParam(r, "start_year"))
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid start year")
			return
		}
		endYear, err := strconv.Atoi(chi.URLParam(r, "end_year"))
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid end year")
			return
		}

		if startYear > endYear {
			response.Text(w, http.StatusBadRequest, "invalid years")
			return
		}

		v, err := h.sv.FindByBrandAndYears(brand, startYear, endYear)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrVehicleNotFound):
				response.Text(w, http.StatusNotFound, "vehicle not found")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.VehicleAttributes.Brand,
				Model:           value.VehicleAttributes.Model,
				Registration:    value.VehicleAttributes.Registration,
				Color:           value.VehicleAttributes.Color,
				FabricationYear: value.VehicleAttributes.FabricationYear,
				Capacity:        value.VehicleAttributes.Capacity,
				MaxSpeed:        value.VehicleAttributes.MaxSpeed,
				FuelType:        value.VehicleAttributes.FuelType,
				Transmission:    value.VehicleAttributes.Transmission,
				Weight:          value.VehicleAttributes.Weight,
				Height:          value.VehicleAttributes.Dimensions.Height,
				Length:          value.VehicleAttributes.Dimensions.Length,
				Width:           value.VehicleAttributes.Dimensions.Width,
			}
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

// GetAverageSpeedByBrand is a method that returns a handler for the route GET /vehicles/average_speed/brand/{brand}
func (h *VehicleDefault) GetAverageSpeedByBrand() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		brand := chi.URLParam(r, "brand")
		if brand == "" {
			response.Text(w, http.StatusBadRequest, "invalid brand")
			return
		}

		speed, err := h.sv.FindAverageSpeedByBrand(brand)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrVehicleNotFound):
				response.Text(w, http.StatusNotFound, "vehicle not found")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message":       "success",
			"average_speed": speed,
		})
	}
}

func (h *VehicleDefault) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid id")
			return
		}

		err = h.sv.Delete(id)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrVehicleNotFound):
				response.Text(w, http.StatusNotFound, "vehicle not found")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		response.Text(w, http.StatusNoContent, "vehicle deleted")

	}
}

func (h *VehicleDefault) GetByFuelType() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fuelType := chi.URLParam(r, "type")
		if fuelType == "" {
			response.Text(w, http.StatusBadRequest, "invalid fuel type")
			return
		}

		v, err := h.sv.FindByFuelType(fuelType)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrVehicleNotFound):
				response.Text(w, http.StatusNotFound, "vehicle not found")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.VehicleAttributes.Brand,
				Model:           value.VehicleAttributes.Model,
				Registration:    value.VehicleAttributes.Registration,
				Color:           value.VehicleAttributes.Color,
				FabricationYear: value.VehicleAttributes.FabricationYear,
				Capacity:        value.VehicleAttributes.Capacity,
				MaxSpeed:        value.VehicleAttributes.MaxSpeed,
				FuelType:        value.VehicleAttributes.FuelType,
				Transmission:    value.VehicleAttributes.Transmission,
				Weight:          value.VehicleAttributes.Weight,
				Height:          value.VehicleAttributes.Dimensions.Height,
				Length:          value.VehicleAttributes.Dimensions.Length,
				Width:           value.VehicleAttributes.Dimensions.Width,
			}
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (h *VehicleDefault) GetByTransmission() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		transmission := chi.URLParam(r, "type")
		if transmission == "" {
			response.Text(w, http.StatusBadRequest, "invalid transmission type")
		}

		v, err := h.sv.FindByTransmissionType(transmission)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrVehicleNotFound):
				response.Text(w, http.StatusNotFound, "vehicle not found")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		data := map[int]VehicleJSON{}
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.VehicleAttributes.Brand,
				Model:           value.VehicleAttributes.Model,
				Registration:    value.VehicleAttributes.Registration,
				Color:           value.VehicleAttributes.Color,
				FabricationYear: value.VehicleAttributes.FabricationYear,
				Capacity:        value.VehicleAttributes.Capacity,
				MaxSpeed:        value.VehicleAttributes.MaxSpeed,
				FuelType:        value.VehicleAttributes.FuelType,
				Transmission:    value.VehicleAttributes.Transmission,
				Weight:          value.VehicleAttributes.Weight,
				Height:          value.VehicleAttributes.Dimensions.Height,
				Length:          value.VehicleAttributes.Dimensions.Length,
				Width:           value.VehicleAttributes.Dimensions.Width,
			}
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (h *VehicleDefault) GetAverageCapacityByBrand() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		brand := chi.URLParam(r, "brand")
		if brand == "" {
			response.Text(w, http.StatusBadRequest, "invalid brand")
			return
		}

		capacity, err := h.sv.FindAverageCapacityByBrand(brand)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrVehicleNotFound):
				response.Text(w, http.StatusNotFound, "vehicle not found")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message":          "success",
			"average_capacity": capacity,
		})
	}
}

func (h *VehicleDefault) CreateMultipleVehicles() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body []BodyRequestVehicleJson
		err := request.JSON(r, &body)
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid body")
			return
		}
		data := map[int]VehicleJSON{}
		for _, value := range body {
			vehicle := internal.Vehicle{
				Id: value.ID,
				VehicleAttributes: internal.VehicleAttributes{
					Brand:           value.Brand,
					Model:           value.Model,
					Registration:    value.Registration,
					Color:           value.Color,
					FabricationYear: value.FabricationYear,
					Capacity:        value.Capacity,
					MaxSpeed:        value.MaxSpeed,
					FuelType:        value.FuelType,
					Transmission:    value.Transmission,
					Weight:          value.Weight,
					Dimensions: internal.Dimensions{
						Height: value.Height,
						Length: value.Length,
						Width:  value.Width,
					},
				},
			}

			if err := h.sv.Create(&vehicle); err != nil {
				switch {
				case errors.Is(err, internal.ErrFieldRequired):
					response.Text(w, http.StatusBadRequest, "invalid body")
				case errors.Is(err, internal.ErrVehicleIdAlreadyExists):
					response.Text(w, http.StatusConflict, "vehicle id already exists")
				default:
					response.Text(w, http.StatusInternalServerError, "internal server error")
				}
				break
			}

			data[value.ID] = VehicleJSON{
				ID:              value.ID,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}

		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (h *VehicleDefault) GetByWeight() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := make(map[string]float64)
		weightMin, ok := r.URL.Query()["min"]
		if ok {
			weightMinFloat, err := strconv.ParseFloat(weightMin[0], 64)
			if err != nil {
				response.Text(w, http.StatusBadRequest, "min_weight must be a number")
				return
			}
			query["min_weight"] = weightMinFloat
		}
		weightMax, ok := r.URL.Query()["max"]
		if ok {
			weightMaxFloat, err := strconv.ParseFloat(weightMax[0], 64)
			if err != nil {
				response.Text(w, http.StatusBadRequest, "max_weight must be a number")
				return
			}
			if weightMaxFloat < query["min_weight"] {
				response.Text(w, http.StatusBadRequest, "invalid range of weights")
				return
			}

			query["max_weight"] = weightMaxFloat
		}

		v, err := h.sv.FindByWeight(query)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrVehicleNotFound):
				response.Text(w, http.StatusNotFound, "vehicle not found")
			case errors.Is(err, internal.ErrInvalidQuery):
				response.Text(w, http.StatusBadRequest, "invalid query")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.VehicleAttributes.Brand,
				Model:           value.VehicleAttributes.Model,
				Registration:    value.VehicleAttributes.Registration,
				Color:           value.VehicleAttributes.Color,
				FabricationYear: value.VehicleAttributes.FabricationYear,
				Capacity:        value.VehicleAttributes.Capacity,
				MaxSpeed:        value.VehicleAttributes.MaxSpeed,
				FuelType:        value.VehicleAttributes.FuelType,
				Transmission:    value.VehicleAttributes.Transmission,
				Weight:          value.VehicleAttributes.Weight,
				Height:          value.VehicleAttributes.Dimensions.Height,
				Length:          value.VehicleAttributes.Dimensions.Length,
				Width:           value.VehicleAttributes.Dimensions.Width,
			}
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (h *VehicleDefault) GetByDimensions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := make(map[string]float64)
		length, ok := r.URL.Query()["length"]
		if ok {
			lengthSplit := strings.Split(length[0], "-")
			if len(lengthSplit) != 2 {
				response.Text(w, http.StatusBadRequest, "invalid length")
				return
			}
			lengthMin, err := strconv.ParseFloat(lengthSplit[0], 64)
			if err != nil {
				response.Text(w, http.StatusBadRequest, "invalid length")
				return
			}
			lengthMax, err := strconv.ParseFloat(lengthSplit[1], 64)
			if err != nil {
				response.Text(w, http.StatusBadRequest, "invalid length")
				return
			}

			if lengthMax < lengthMin {
				response.Text(w, http.StatusBadRequest, "invalid length range")
				return
			}
			query["min_length"] = lengthMin
			query["max_length"] = lengthMax
		}

		width, ok := r.URL.Query()["width"]
		if ok {
			widthSplit := strings.Split(width[0], "-")
			if len(widthSplit) != 2 {
				response.Text(w, http.StatusBadRequest, "invalid width")
				return
			}
			widthMin, err := strconv.ParseFloat(widthSplit[0], 64)
			if err != nil {
				response.Text(w, http.StatusBadRequest, "invalid width")
				return
			}
			widthMax, err := strconv.ParseFloat(widthSplit[1], 64)
			if err != nil {
				response.Text(w, http.StatusBadRequest, "invalid width")
				return
			}
			if widthMax < widthMin {
				response.Text(w, http.StatusBadRequest, "invalid width range")
				return
			}
			query["min_width"] = widthMin
			query["max_width"] = widthMax
		}

		v, err := h.sv.FindByDimensions(query)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrVehicleNotFound):
				response.Text(w, http.StatusNotFound, "vehicle not found")
			case errors.Is(err, internal.ErrInvalidQuery):
				response.Text(w, http.StatusBadRequest, "invalid query")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.VehicleAttributes.Brand,
				Model:           value.VehicleAttributes.Model,
				Registration:    value.VehicleAttributes.Registration,
				Color:           value.VehicleAttributes.Color,
				FabricationYear: value.VehicleAttributes.FabricationYear,
				Capacity:        value.VehicleAttributes.Capacity,
				MaxSpeed:        value.VehicleAttributes.MaxSpeed,
				FuelType:        value.VehicleAttributes.FuelType,
				Transmission:    value.VehicleAttributes.Transmission,
				Weight:          value.VehicleAttributes.Weight,
				Height:          value.VehicleAttributes.Dimensions.Height,
				Length:          value.VehicleAttributes.Dimensions.Length,
				Width:           value.VehicleAttributes.Dimensions.Width,
			}
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})

	}
}

func (h *VehicleDefault) UpdateFuelType() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid id")
			return
		}

		vehicle, err := h.sv.GetById(id)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrVehicleNotFound):
				response.Text(w, http.StatusNotFound, "vehicle not found")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		reqBody := BodyRequestVehicleJson{
			ID:              vehicle.Id,
			Brand:           vehicle.VehicleAttributes.Brand,
			Model:           vehicle.VehicleAttributes.Model,
			Registration:    vehicle.VehicleAttributes.Registration,
			Color:           vehicle.VehicleAttributes.Color,
			FabricationYear: vehicle.VehicleAttributes.FabricationYear,
			Capacity:        vehicle.VehicleAttributes.Capacity,
			MaxSpeed:        vehicle.VehicleAttributes.MaxSpeed,
			FuelType:        vehicle.VehicleAttributes.FuelType,
			Transmission:    vehicle.VehicleAttributes.Transmission,
			Weight:          vehicle.VehicleAttributes.Weight,
			Height:          vehicle.VehicleAttributes.Dimensions.Height,
			Length:          vehicle.VehicleAttributes.Dimensions.Length,
			Width:           vehicle.VehicleAttributes.Dimensions.Width,
		}

		if err := request.JSON(r, &reqBody); err != nil {
			response.Text(w, http.StatusBadRequest, "invalid body")
			return
		}

		vehicle.VehicleAttributes.FuelType = reqBody.FuelType

		err = h.sv.Update(&vehicle)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrVehicleNotFound):
				response.Text(w, http.StatusNotFound, "vehicle not found")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		data := VehicleJSON{
			ID:              reqBody.ID,
			Brand:           reqBody.Brand,
			Model:           reqBody.Model,
			Registration:    reqBody.Registration,
			Color:           reqBody.Color,
			FabricationYear: reqBody.FabricationYear,
			Capacity:        reqBody.Capacity,
			MaxSpeed:        reqBody.MaxSpeed,
			FuelType:        reqBody.FuelType,
			Transmission:    reqBody.Transmission,
			Weight:          reqBody.Weight,
			Height:          reqBody.Height,
			Length:          reqBody.Length,
			Width:           reqBody.Width,
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (h *VehicleDefault) UpdateSpeed() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid id")
			return
		}

		vehicle, err := h.sv.GetById(id)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrVehicleNotFound):
				response.Text(w, http.StatusNotFound, "vehicle not found")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}
		reqBody := BodyRequestVehicleJson{
			ID:              vehicle.Id,
			Brand:           vehicle.VehicleAttributes.Brand,
			Model:           vehicle.VehicleAttributes.Model,
			Registration:    vehicle.VehicleAttributes.Registration,
			Color:           vehicle.VehicleAttributes.Color,
			FabricationYear: vehicle.VehicleAttributes.FabricationYear,
			Capacity:        vehicle.VehicleAttributes.Capacity,
			MaxSpeed:        vehicle.VehicleAttributes.MaxSpeed,
			FuelType:        vehicle.VehicleAttributes.FuelType,
			Transmission:    vehicle.VehicleAttributes.Transmission,
			Weight:          vehicle.VehicleAttributes.Weight,
			Height:          vehicle.VehicleAttributes.Dimensions.Height,
			Length:          vehicle.VehicleAttributes.Dimensions.Length,
			Width:           vehicle.VehicleAttributes.Dimensions.Width,
		}

		if err := request.JSON(r, &reqBody); err != nil {
			response.Text(w, http.StatusBadRequest, "invalid body")
			return
		}
		fmt.Println(reqBody)
		vehicle.VehicleAttributes.MaxSpeed = reqBody.MaxSpeed
		fmt.Println(vehicle)
		err = h.sv.Update(&vehicle)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrVehicleNotFound):
				response.Text(w, http.StatusNotFound, "vehicle not found")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		data := VehicleJSON{
			ID:              reqBody.ID,
			Brand:           reqBody.Brand,
			Model:           reqBody.Model,
			Registration:    reqBody.Registration,
			Color:           reqBody.Color,
			FabricationYear: reqBody.FabricationYear,
			Capacity:        reqBody.Capacity,
			MaxSpeed:        reqBody.MaxSpeed,
			FuelType:        reqBody.FuelType,
			Transmission:    reqBody.Transmission,
			Weight:          reqBody.Weight,
			Height:          reqBody.Height,
			Length:          reqBody.Length,
			Width:           reqBody.Width,
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (h *VehicleDefault) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid id")
			return
		}

		vehicle, err := h.sv.GetById(id)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrVehicleNotFound):
				response.Text(w, http.StatusNotFound, "vehicle not found")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		data := VehicleJSON{
			ID:              vehicle.Id,
			Brand:           vehicle.VehicleAttributes.Brand,
			Model:           vehicle.VehicleAttributes.Model,
			Registration:    vehicle.VehicleAttributes.Registration,
			Color:           vehicle.VehicleAttributes.Color,
			FabricationYear: vehicle.VehicleAttributes.FabricationYear,
			Capacity:        vehicle.VehicleAttributes.Capacity,
			MaxSpeed:        vehicle.VehicleAttributes.MaxSpeed,
			FuelType:        vehicle.VehicleAttributes.FuelType,
			Transmission:    vehicle.VehicleAttributes.Transmission,
			Weight:          vehicle.VehicleAttributes.Weight,
			Height:          vehicle.VehicleAttributes.Dimensions.Height,
			Length:          vehicle.VehicleAttributes.Dimensions.Length,
			Width:           vehicle.VehicleAttributes.Dimensions.Width,
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}
