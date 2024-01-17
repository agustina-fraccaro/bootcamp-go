package internal

import "errors"

var (
	ErrFieldRequired   = errors.New("field is required")
	ErrVehicleNotFound = errors.New("vehicle not found")
)

// VehicleService is an interface that represents a vehicle service
type VehicleService interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	Create(vehicle *Vehicle) (err error)
	FindByColorAndYear(color string, year int) (v map[int]Vehicle, err error)
	FindByBrandAndYears(brand string, startYear int, endYear int) (v map[int]Vehicle, err error)
	FindAverageSpeedByBrand(brand string) (averageSpeed float64, err error)
	Delete(id int) (err error)
	FindByFuelType(fuelType string) (v map[int]Vehicle, err error)
	FindByTransmissionType(transmission string) (v map[int]Vehicle, err error)
	Update(vehicle *Vehicle) (err error)
	FindAverageCapacityByBrand(brand string) (averageCapacity float64, err error)
	FindByWeight(query map[string]float64) (v map[int]Vehicle, err error)
	FindByDimensions(query map[string]any) (v map[int]Vehicle, err error)
	GetById(id int) (v Vehicle, err error)
}
