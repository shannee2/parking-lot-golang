package attendant

import (
	"parkinglot/errors"
	"parkinglot/parkinglot"
	"parkinglot/ticket"
	"parkinglot/vehicle"
)

type Attendant struct {
	parkingLots []*parkinglot.ParkingLot
}

func (a *Attendant) AssignParkingLot(lot *parkinglot.ParkingLot) {
	a.parkingLots = append(a.parkingLots, lot)
}

func (a *Attendant) Park(vehicle *vehicle.Vehicle) (*ticket.Ticket, error) {
	if len(a.parkingLots) == 0 {
		return nil, errors.ErrNoParkingLotAssignedToAttendant
	}

	var selectedLot *parkinglot.ParkingLot

	for _, l := range a.parkingLots {
		if !l.IsFull() {
			if selectedLot == nil || l.CountParkedVehicles() < selectedLot.CountParkedVehicles() {
				selectedLot = l
			}
		}
	}

	if selectedLot != nil {
		t, err := selectedLot.Park(vehicle)
		if err == nil {
			return t, nil
		}
	}

	return nil, errors.ErrAllLotsAreFull
}

func (a *Attendant) IsParked(registrationNumber string) bool {
	for _, l := range a.parkingLots {
		if l.IsVehicleParked(registrationNumber) {
			return true
		}
	}
	return false
}

func (a *Attendant) Unpark(t *ticket.Ticket) error {
	for _, l := range a.parkingLots {
		err := l.UnPark(t)
		return err
	}
	return errors.ErrTicketNotFound
}

func NewAttendant() *Attendant {
	return &Attendant{
		parkingLots: make([]*parkinglot.ParkingLot, 0),
	}
}
