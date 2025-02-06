package attendant

import (
	"parkinglot/errors"
	"parkinglot/parkinglot"
	"parkinglot/ticket"
	"parkinglot/vehicle"
)

type Attendant struct {
	ParkingLots []*parkinglot.ParkingLot
}

func (a *Attendant) AssignParkingLot(lot *parkinglot.ParkingLot) {
	a.ParkingLots = append(a.ParkingLots, lot)
}

func (a *Attendant) Park(vehicle *vehicle.Vehicle) (*ticket.Ticket, error) {
	for _, l := range a.ParkingLots {
		t, _ := l.Park(vehicle)
		if t != nil {
			return t, nil
		}
	}
	return nil, errors.ErrAllLotsAreFull
}

func (a *Attendant) IsParked(registrationNumber string) bool {
	for _, l := range a.ParkingLots {
		if l.IsVehicleParked(registrationNumber) {
			return true
		}
	}
	return false
}

func (a *Attendant) Unpark(t *ticket.Ticket) error {
	for _, l := range a.ParkingLots {
		err := l.UnPark(t)
		return err
	}
	return errors.ErrTicketNotFound
}

func NewAttendant() *Attendant {
	return &Attendant{
		ParkingLots: make([]*parkinglot.ParkingLot, 0),
	}
}
