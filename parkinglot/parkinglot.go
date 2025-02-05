package parkinglot

import (
	"parkinglot/errors"
	"parkinglot/slot"
	"parkinglot/ticket"
	"parkinglot/vehicle"
)

type ParkingLot struct {
	slots []*slot.Slot
}

func (l *ParkingLot) Park(vehicle *vehicle.Vehicle) (*ticket.Ticket, error) {

	availableSlot := l.fetchNearestAvailableSlot()
	if availableSlot != nil {
		availableSlot.Park(vehicle)
		return ticket.NewTicket(), nil
	}
	return nil, errors.ErrAllSlotsOccupied
}

func (l *ParkingLot) fetchNearestAvailableSlot() *slot.Slot {
	for _, s := range l.slots {
		if !s.IsOccupied() {
			return s
		}
	}
	return nil
}

func (l *ParkingLot) IsVehicleParked(registrationNumber string) bool {
	for _, s := range l.slots {
		if s.IsVehicleParked(registrationNumber) {
			return true
		}

	}
	return false
}

func NewParkingLot(size int) (*ParkingLot, error) {
	if size <= 0 {
		return nil, errors.ErrInvalidSlotSize
	}

	parkingLot := &ParkingLot{
		slots: make([]*slot.Slot, size),
	}
	for i := 0; i < size; i++ {
		parkingLot.slots[i] = slot.NewSlot()
	}
	return parkingLot, nil
}
