package attendant

import (
	"fmt"
	"parkinglot/errors"
	"parkinglot/parkinglot"
	"parkinglot/ticket"
	"parkinglot/vehicle"
)

type Attendant struct {
	parkingLots     []*parkinglot.ParkingLot
	parkingStrategy ParkingStrategy
}

func (a *Attendant) AssignParkingLot(lot *parkinglot.ParkingLot) {
	a.parkingLots = append(a.parkingLots, lot)
}

func (a *Attendant) ChangeStrategy(strategy ParkingStrategy) {
	a.parkingStrategy = strategy
}

func (a *Attendant) Park(vehicle *vehicle.Vehicle) (*ticket.Ticket, error) {
	if len(a.parkingLots) == 0 {
		return nil, errors.ErrNoParkingLotAssignedToAttendant
	}

	selectedLot, err := a.parkingStrategy.selectLot(a.parkingLots)
	if err != nil {
		return nil, err
	}

	return selectedLot.Park(vehicle)
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
		if err == nil {
			return nil
		}
	}
	fmt.Println("never seen this ticket")
	return errors.ErrTicketNotFound
}

func NewAttendant(strategy ParkingStrategy) *Attendant {
	return &Attendant{
		parkingLots:     make([]*parkinglot.ParkingLot, 0),
		parkingStrategy: strategy,
	}
}
