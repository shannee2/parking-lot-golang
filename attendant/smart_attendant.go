package attendant

import (
	"parkinglot/errors"
	"parkinglot/parkinglot"
	"parkinglot/ticket"
	"parkinglot/vehicle"
)

type SmartAttendant struct {
	Attendant
}

func (a *SmartAttendant) Park(vehicle *vehicle.Vehicle) (*ticket.Ticket, error) {
	if len(a.ParkingLots) == 0 {
		return nil, errors.ErrNoParkingLotAssignedToAttendant
	}

	var selectedLot *parkinglot.ParkingLot

	for _, l := range a.ParkingLots {
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

func NewSmartAttendant() *SmartAttendant {
	return &SmartAttendant{
		Attendant: Attendant{
			ParkingLots: make([]*parkinglot.ParkingLot, 0),
		},
	}
}
