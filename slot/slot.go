package slot

import (
	"parkinglot/errors"
	"parkinglot/vehicle"
)

type Slot struct {
	vehicle *vehicle.Vehicle
}

func (s *Slot) Park(v vehicle.Vehicle) error {
	if s.isOccupied() {
		return errors.ErrSlotAlreadyOccupied
	}
	s.vehicle = &v
	return nil
}

func (s *Slot) isOccupied() bool {
	return s.vehicle != nil
}

func (s *Slot) UnPark() error {
	if s.vehicle == nil {
		return errors.ErrSlotEmpty
	}
	s.vehicle = nil
	return nil
}

func (s *Slot) HasVehicleColor(color vehicle.VehicleColor) interface{} {
	return s.vehicle != nil && s.vehicle.HasColor(color)
}

func NewSlot() *Slot {
	return &Slot{
		vehicle: nil,
	}
}
