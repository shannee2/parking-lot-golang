package policeman

import (
	"parkinglot/parkinglot"
	"parkinglot/vehicle"
	"testing"
)

func TestNotifyOwnerWhenParkingLotFull(t *testing.T) {
	p := NewPoliceman()

	l, _ := parkinglot.NewParkingLot(2)

	l.AddObserver(p)

	r1 := "RJ-12-JI-1234"
	r2 := "RJ-12-JI-5678"
	v1 := vehicle.NewVehicle(r1, vehicle.Red)
	v2 := vehicle.NewVehicle(r2, vehicle.Blue)

	_, _ = l.Park(v1)
	ticket, _ := l.Park(v2)
	l.UnPark(ticket)
}
