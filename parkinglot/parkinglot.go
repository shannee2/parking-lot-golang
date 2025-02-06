package parkinglot

import (
	"parkinglot/errors"
	"parkinglot/slot"
	"parkinglot/ticket"
	"parkinglot/vehicle"
)

type ParkingLot struct {
	slots        []*slot.Slot
	ticketToSlot map[*ticket.Ticket]*slot.Slot
}

func (l *ParkingLot) Park(vehicle *vehicle.Vehicle) (*ticket.Ticket, error) {

	availableSlot := l.fetchNearestAvailableSlot()
	if availableSlot != nil {
		availableSlot.Park(vehicle)
		t := ticket.NewTicket()
		l.ticketToSlot[t] = availableSlot
		return t, nil
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

func (l *ParkingLot) UnPark(t *ticket.Ticket) error {
	s, exists := l.ticketToSlot[t]
	if exists {
		err := s.UnPark()
		if err == nil {
			delete(l.ticketToSlot, t)
			return nil
		}
	}
	return errors.ErrTicketNotFound
}

func (l *ParkingLot) countVehiclesWithColor(color vehicle.VehicleColor) int {
	count := 0
	for _, s := range l.slots {
		if s.HasVehicleColor(color) {
			count++
		}
	}
	return count
}

func NewParkingLot(size int) (*ParkingLot, error) {
	if size <= 0 {
		return nil, errors.ErrInvalidSlotSize
	}

	parkingLot := &ParkingLot{
		slots:        make([]*slot.Slot, size),
		ticketToSlot: make(map[*ticket.Ticket]*slot.Slot),
	}
	for i := 0; i < size; i++ {
		parkingLot.slots[i] = slot.NewSlot()
	}
	return parkingLot, nil
}
