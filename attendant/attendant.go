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
	for _, l := range a.parkingLots {
		t, _ := l.Park(vehicle)
		if t != nil {
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
