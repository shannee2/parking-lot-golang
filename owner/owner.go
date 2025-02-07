package owner

import (
	"fmt"
	"parkinglot/attendant"
	"parkinglot/parkinglot"
)

type Owner struct {
	attendant.Attendant
	attendants []*attendant.Attendant
}

func (o *Owner) AssignAttendant(a *attendant.Attendant) {
	o.attendants = append(o.attendants, a)
}

func (o *Owner) AssignParkingLotToAttendant(a1 *attendant.Attendant, p *parkinglot.ParkingLot) {
	a1.AssignParkingLot(p)
}

func (o *Owner) OnParkingLotFull(l *parkinglot.ParkingLot) {
	fmt.Println("Alert! Parking lot with id=", l.Id, " is full")
}

func (o *Owner) OnParkingLotAvailable(l *parkinglot.ParkingLot) {
	fmt.Println("Alert! Parking lot with id=", l.Id, " has available space")
}

func New() *Owner {
	return &Owner{
		Attendant: attendant.Attendant{
			ParkingLots: make([]*parkinglot.ParkingLot, 0),
		},
		attendants: make([]*attendant.Attendant, 0),
	}
}
