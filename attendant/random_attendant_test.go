package attendant

import (
	"parkinglot/parkinglot"
	"parkinglot/vehicle"
	"testing"
)

func TestRandomAttendant_Park(t *testing.T) {
	att := NewRandomAttendant()
	l1, _ := parkinglot.New(3)
	l2, _ := parkinglot.New(4)
	att.AssignParkingLot(l1)
	att.AssignParkingLot(l2)

	v1 := vehicle.New("RJ-12-JI-1234", vehicle.Red)
	v2 := vehicle.New("RJ-12-JI-5678", vehicle.Blue)
	v3 := vehicle.New("RJ-12-JI-9012", vehicle.Green)

	att.Park(v1)
	att.Park(v2)
	att.Park(v3)
	att.PrintSlots()
}
