package slot

import (
	"parkinglot/errors"
	"parkinglot/vehicle"
)

type Slot struct {
	vehicle *vehicle.Vehicle
}

func (s *Slot) Park(v *vehicle.Vehicle) error {
	if s.IsOccupied() {
		return errors.ErrSlotAlreadyOccupied
	}
	s.vehicle = v
	return nil
}

func (s *Slot) IsOccupied() bool {
	return s.vehicle != nil
}

// particular excepiton

func (s *Slot) UnPark() error {
	if s.vehicle == nil {
		return errors.ErrSlotEmpty
	}
	s.vehicle = nil
	return nil
}

func (s *Slot) HasVehicleColor(color vehicle.VehicleColor) bool {
	return s.vehicle != nil && s.vehicle.HasColor(color)
}

func (s *Slot) IsVehicleParked(registrationNumber string) bool {
	return s.vehicle != nil && s.vehicle.HasRegistrationNumber(registrationNumber)
}

func New() *Slot {
	return &Slot{
		vehicle: nil,
	}
}
