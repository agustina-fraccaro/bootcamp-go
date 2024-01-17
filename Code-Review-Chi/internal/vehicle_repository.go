package internal

import "errors"

var (
	// ErrVehicleIdAlreadyExists is the error returned when a vehicle id already exists
	ErrVehicleIdAlreadyExists = errors.New("vehicle id already exists")
	// ErrInvalidQuery is the error returned when a query is invalid
	ErrInvalidQuery = errors.New("repository: invalid query")
)

// VehicleRepository is an interface that represents a vehicle repository
type VehicleRepository interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	Create(vehicle *Vehicle) (err error)
	FindByColorAndYear(color string, year int) (v map[int]Vehicle, err error)
	FindByBrandAndYears(brand string, startYear int, endYear int) (v map[int]Vehicle, err error)
	FindByBrand(brand string) (v map[int]Vehicle, err error)
	Delete(id int) (err error)
	FindByFuelType(fuelType string) (v map[int]Vehicle, err error)
	FindByTransmissionType(transmission string) (v map[int]Vehicle, err error)
	Update(vehicle *Vehicle) (err error)
	FindByDimensions(query map[string]any) (v map[int]Vehicle, err error)
	FindByWeight(query map[string]float64) (v map[int]Vehicle, err error)
	GetById(id int) (v Vehicle, err error)
}
