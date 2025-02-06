package policeman

import (
	"fmt"
	"parkinglot/parkinglot"
)

type PoliceMan struct {
}

func NewPoliceman() *PoliceMan {
	return &PoliceMan{}
}

func (o *PoliceMan) OnParkingLotFull(l *parkinglot.ParkingLot) {
	fmt.Println("Alert! Parking lot with id=", l.Id, " is full")
}

func (o *PoliceMan) OnParkingLotAvailable(l *parkinglot.ParkingLot) {
	fmt.Println("Alert! Parking lot with id=", l.Id, " has available space")
}
