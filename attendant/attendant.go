package attendant

import (
	"parkinglot/parkinglot"
)

type Attendant struct {
	parkingLots []*parkinglot.ParkingLot
}

func (a *Attendant) AssignParkingLot(lot *parkinglot.ParkingLot) error {
	a.parkingLots = append(a.parkingLots, lot)
	return nil
}

func NewAttendant() *Attendant {
	return &Attendant{
		parkingLots: make([]*parkinglot.ParkingLot, 0),
	}
}
