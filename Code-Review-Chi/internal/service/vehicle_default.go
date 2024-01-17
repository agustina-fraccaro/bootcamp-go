package service

import "app/internal"

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(rp internal.VehicleRepository) *VehicleDefault {
	return &VehicleDefault{rp: rp}
}

// VehicleDefault is a struct that represents the default service for vehicles
type VehicleDefault struct {
	// rp is the repository that will be used by the service
	rp internal.VehicleRepository
}

// FindAll is a method that returns a map of all vehicles
func (s *VehicleDefault) FindAll() (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindAll()
	return
}

// Create is a method that creates a vehicle
func (s *VehicleDefault) Create(vehicle *internal.Vehicle) (err error) {
	if err = s.ValidateVehicle(vehicle); err != nil {
		return
	}

	err = s.rp.Create(vehicle)
	return
}

func (s *VehicleDefault) FindByColorAndYear(color string, year int) (v map[int]internal.Vehicle, err error) {
	if v, err = s.rp.FindByColorAndYear(color, year); err != nil {
		return
	}

	if len(v) == 0 {
		err = internal.ErrVehicleNotFound
	}
	return
}

func (s *VehicleDefault) FindByBrandAndYears(brand string, startYear int, endYear int) (v map[int]internal.Vehicle, err error) {
	if v, err = s.rp.FindByBrandAndYears(brand, startYear, endYear); err != nil {
		return
	}

	if len(v) == 0 {
		err = internal.ErrVehicleNotFound
	}
	return
}

func (s *VehicleDefault) FindAverageSpeedByBrand(brand string) (averageSpeed float64, err error) {
	v := make(map[int]internal.Vehicle)
	if v, err = s.rp.FindByBrand(brand); err != nil {
		return
	}

	if len(v) == 0 {
		err = internal.ErrVehicleNotFound
		return
	}

	for _, vehicle := range v {
		averageSpeed += vehicle.VehicleAttributes.MaxSpeed
	}
	averageSpeed /= float64(len(v))
	return
}

func (s *VehicleDefault) Delete(id int) (err error) {
	if err = s.rp.Delete(id); err != nil {
		return
	}

	return
}

func (s *VehicleDefault) FindByFuelType(fuelType string) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindByFuelType(fuelType)
	if err != nil {
		return
	}

	if len(v) == 0 {
		err = internal.ErrVehicleNotFound
	}

	return
}

func (s *VehicleDefault) FindByTransmissionType(transmission string) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindByTransmissionType(transmission)
	if err != nil {
		return
	}

	if len(v) == 0 {
		err = internal.ErrVehicleNotFound
	}

	return
}

func (s *VehicleDefault) Update(vehicle *internal.Vehicle) (err error) {
	if err = s.ValidateVehicle(vehicle); err != nil {
		return
	}

	err = s.rp.Update(vehicle)
	return
}

func (s *VehicleDefault) GetById(id int) (v internal.Vehicle, err error) {
	v, err = s.rp.GetById(id)
	if err != nil {
		return
	}

	return
}

func (s *VehicleDefault) FindByDimensions(query map[string]any) (v map[int]internal.Vehicle, err error) {
	if len(query) != 4 {
		err = internal.ErrInvalidQuery
		return
	}

	v, err = s.rp.FindByDimensions(query)
	if err != nil {
		return
	}

	if len(v) == 0 {
		err = internal.ErrVehicleNotFound
	}

	return
}

func (s *VehicleDefault) FindByWeight(query map[string]float64) (v map[int]internal.Vehicle, err error) {
	if len(query) != 2 {
		err = internal.ErrInvalidQuery
		return
	}

	v, err = s.rp.FindByWeight(query)
	if err != nil {
		return
	}

	if len(v) == 0 {
		err = internal.ErrVehicleNotFound
	}

	return
}

func (s *VehicleDefault) FindAverageCapacityByBrand(brand string) (averageCapacity float64, err error) {
	v := make(map[int]internal.Vehicle)
	v, err = s.rp.FindByBrand(brand)
	if err != nil {
		return
	}

	if len(v) == 0 {
		err = internal.ErrVehicleNotFound
		return
	}

	for _, value := range v {
		averageCapacity += float64(value.VehicleAttributes.Capacity)
	}

	averageCapacity /= float64(len(v))
	return
}

// ValidateVehicle is a method that validates a vehicle
func (s *VehicleDefault) ValidateVehicle(vehicle *internal.Vehicle) (err error) {
	if (*vehicle).VehicleAttributes.Brand == "" {
		err = internal.ErrFieldRequired
		return
	}

	if (*vehicle).VehicleAttributes.Model == "" {
		err = internal.ErrFieldRequired
		return
	}

	if (*vehicle).VehicleAttributes.Registration == "" {
		err = internal.ErrFieldRequired
		return
	}

	if (*vehicle).VehicleAttributes.Color == "" {
		err = internal.ErrFieldRequired
		return
	}

	if (*vehicle).VehicleAttributes.FabricationYear == 0 {
		err = internal.ErrFieldRequired
		return
	}

	if (*vehicle).VehicleAttributes.Capacity == 0 {
		err = internal.ErrFieldRequired
		return
	}

	if (*vehicle).VehicleAttributes.MaxSpeed == 0 {
		err = internal.ErrFieldRequired
		return
	}

	if (*vehicle).VehicleAttributes.FuelType == "" {
		err = internal.ErrFieldRequired
		return
	}

	if (*vehicle).VehicleAttributes.Transmission == "" {
		err = internal.ErrFieldRequired
		return
	}

	return
}
