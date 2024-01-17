package repository

import (
	"app/internal"
)

// NewVehicleMap is a function that returns a new instance of VehicleMap
func NewVehicleMap(db map[int]internal.Vehicle) *VehicleMap {
	// default db
	defaultDb := make(map[int]internal.Vehicle)
	if db != nil {
		defaultDb = db
	}
	return &VehicleMap{db: defaultDb}
}

// VehicleMap is a struct that represents a vehicle repository
type VehicleMap struct {
	// db is a map of vehicles
	db map[int]internal.Vehicle
}

// FindAll is a method that returns a map of all vehicles
func (r *VehicleMap) FindAll() (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db
	for key, value := range r.db {
		v[key] = value
	}

	return
}

// Create is a method that creates a vehicle
func (r *VehicleMap) Create(vehicle *internal.Vehicle) (err error) {
	if err = r.ValidateVehicleId((*vehicle).Id); err != nil {
		return
	}

	(*r).db[(*vehicle).Id] = *vehicle
	return
}

// ValidateVehicleId is a method that validates if a vehicle id already exists
func (r *VehicleMap) ValidateVehicleId(id int) (err error) {
	if _, ok := (*r).db[id]; ok {
		err = internal.ErrVehicleIdAlreadyExists
		return
	}
	return
}

func (r *VehicleMap) FindByColorAndYear(color string, year int) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)
	for key, value := range (*r).db {
		if value.VehicleAttributes.Color == color && value.VehicleAttributes.FabricationYear == year {
			v[key] = value
		}
	}

	return
}

func (r *VehicleMap) FindByBrandAndYears(brand string, startYear int, endYear int) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)
	for key, value := range (*r).db {
		if value.VehicleAttributes.Brand == brand && value.VehicleAttributes.FabricationYear >= startYear && value.VehicleAttributes.FabricationYear <= endYear {
			v[key] = value
		}
	}

	return
}

func (r *VehicleMap) FindByBrand(brand string) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)
	for key, value := range (*r).db {
		if value.VehicleAttributes.Brand == brand {
			v[key] = value
		}
	}

	return
}

func (r *VehicleMap) Delete(id int) (err error) {
	_, ok := (*r).db[id]
	if !ok {
		err = internal.ErrVehicleNotFound
		return
	}

	delete((*r).db, id)
	return
}

func (r *VehicleMap) FindByFuelType(fuelType string) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)
	for key, value := range (*r).db {
		if value.VehicleAttributes.FuelType == fuelType {
			v[key] = value
		}
	}
	return
}

func (r *VehicleMap) FindByTransmissionType(transmission string) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)
	for key, value := range (*r).db {
		if value.VehicleAttributes.Transmission == transmission {
			v[key] = value
		}
	}
	return
}

func (r *VehicleMap) Update(vehicle *internal.Vehicle) (err error) {
	_, ok := (*r).db[(*vehicle).Id]
	if !ok {
		err = internal.ErrVehicleNotFound
		return
	}

	(*r).db[(*vehicle).Id] = *vehicle
	return
}

func (r *VehicleMap) GetById(id int) (v internal.Vehicle, err error) {
	v, ok := (*r).db[id]
	if !ok {
		err = internal.ErrVehicleNotFound
		return
	}
	return
}

func (r *VehicleMap) FindByWeight(query map[string]float64) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)
	for key, value := range (*r).db {
		weightMin, ok := query["min_weight"]
		if !ok {
			err = internal.ErrInvalidQuery
			return
		}

		weightMax, ok := query["max_weight"]
		if !ok {
			err = internal.ErrInvalidQuery
			return
		}
		if value.VehicleAttributes.Weight >= weightMin && value.VehicleAttributes.Weight <= weightMax {
			v[key] = value
		}
	}
	return
}

func (r *VehicleMap) FindByDimensions(query map[string]any) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)
	for key, value := range (*r).db {
		lengthMin, ok := query["min_length"]
		if !ok {
			err = internal.ErrInvalidQuery
			return
		}

		lengthMinFloat, ok := lengthMin.(float64)
		if !ok {
			err = internal.ErrInvalidQuery
			return
		}

		lengthMax, ok := query["max_length"]
		if !ok {
			err = internal.ErrInvalidQuery
			return
		}
		lengthMaxFloat, ok := lengthMax.(float64)
		if !ok {
			err = internal.ErrInvalidQuery
			return
		}

		widthMin, ok := query["min_width"]
		if !ok {
			err = internal.ErrInvalidQuery
			return
		}

		widthMinFloat, ok := widthMin.(float64)
		if !ok {
			err = internal.ErrInvalidQuery
			return
		}

		widthMax, ok := query["max_width"]
		if !ok {
			err = internal.ErrInvalidQuery
			return
		}

		widthMaxFloat, ok := widthMax.(float64)
		if !ok {
			err = internal.ErrInvalidQuery
			return
		}

		if value.VehicleAttributes.Dimensions.Length >= lengthMinFloat && value.VehicleAttributes.Dimensions.Length <= lengthMaxFloat && value.VehicleAttributes.Dimensions.Width >= widthMinFloat && value.VehicleAttributes.Dimensions.Width <= widthMaxFloat {
			v[key] = value
		}

	}

	return

}
