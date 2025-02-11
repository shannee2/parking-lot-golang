package traffic_cop

import (
	"fmt"
	"parkinglot/parkinglot"
)

type TrafficCop struct {
}

func New() *TrafficCop {
	return &TrafficCop{}
}

func (o *TrafficCop) OnParkingLotFull(l *parkinglot.ParkingLot) {
	fmt.Println("Alert! Parking lot with id=", l.Id, " is full")
}

func (o *TrafficCop) OnParkingLotAvailable(l *parkinglot.ParkingLot) {
	fmt.Println("Alert! Parking lot with id=", l.Id, " has available space")
}
