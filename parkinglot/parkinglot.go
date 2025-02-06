package parkinglot

import (
	"github.com/google/uuid"
	"parkinglot/errors"
	"parkinglot/slot"
	"parkinglot/ticket"
	"parkinglot/vehicle"
)

type ParkingLotObserver interface {
	OnParkingLotFull(fullLot *ParkingLot)
	OnParkingLotAvailable(availableLot *ParkingLot)
}

type ParkingLot struct {
	Id           string
	slots        []*slot.Slot
	ticketToSlot map[*ticket.Ticket]*slot.Slot
	observers    []ParkingLotObserver
}

func (l *ParkingLot) AddObserver(o ParkingLotObserver) {
	l.observers = append(l.observers, o)
}

func (l *ParkingLot) notifyObserversLotFull() {
	for _, o := range l.observers {
		o.OnParkingLotFull(l)
	}
}

func (l *ParkingLot) notifyObserversLotAvailable() {
	for _, o := range l.observers {
		o.OnParkingLotAvailable(l)
	}
}

func generateUniqueID() string {
	return uuid.New().String()
}

func (l *ParkingLot) Park(vehicle *vehicle.Vehicle) (*ticket.Ticket, error) {

	availableSlot := l.fetchNearestAvailableSlot()
	if availableSlot != nil {
		availableSlot.Park(vehicle)
		t := ticket.NewTicket()
		l.ticketToSlot[t] = availableSlot
		// Notify observers
		if l.IsFull() {
			l.notifyObserversLotFull()
		}
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
		wasFull := l.IsFull()
		err := s.UnPark()
		if err == nil {
			delete(l.ticketToSlot, t)
			if wasFull {
				l.notifyObserversLotAvailable()
			}
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
		Id:           generateUniqueID(),
		slots:        make([]*slot.Slot, size),
		ticketToSlot: make(map[*ticket.Ticket]*slot.Slot),
		observers:    []ParkingLotObserver{},
	}
	for i := 0; i < size; i++ {
		parkingLot.slots[i] = slot.NewSlot()
	}
	return parkingLot, nil
}

func (l *ParkingLot) CountParkedVehicles() int {
	count := 0
	for _, s := range l.slots {
		if s.IsOccupied() {
			count++
		}
	}
	return count
}

func (l *ParkingLot) IsFull() bool {
	for _, s := range l.slots {
		if !s.IsOccupied() {
			return false
		}
	}
	return true
}
