package attendant

import (
	"fmt"
	"parkinglot/errors"
	"parkinglot/parkinglot"
	"parkinglot/ticket"
	"parkinglot/vehicle"
)

type Attendant struct {
	ParkingLots     []*parkinglot.ParkingLot
	ParkingStrategy ParkingStrategy
}

func (a *Attendant) AssignParkingLot(lot *parkinglot.ParkingLot) {
	a.ParkingLots = append(a.ParkingLots, lot)
}

func (a *Attendant) ChangeStrategy(strategy ParkingStrategy) {
	a.ParkingStrategy = strategy
}

func (a *Attendant) Park(vehicle *vehicle.Vehicle) (*ticket.Ticket, error) {
	if len(a.ParkingLots) == 0 {
		return nil, errors.ErrNoParkingLotAssignedToAttendant
	}

	selectedLot, err := a.ParkingStrategy.selectLot(a.ParkingLots)
	if err != nil {
		return nil, err
	}

	return selectedLot.Park(vehicle)
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
		if err == nil {
			return nil
		}
	}
	fmt.Println("never seen this ticket")
	return errors.ErrTicketNotFound
}

func NewAttendant(strategy ParkingStrategy) *Attendant {
	return &Attendant{
		ParkingLots:     make([]*parkinglot.ParkingLot, 0),
		ParkingStrategy: strategy,
	}
}
