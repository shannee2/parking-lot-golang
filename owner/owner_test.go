package owner

import (
	"github.com/stretchr/testify/assert"
	"parkinglot/attendant"
	"parkinglot/parkinglot"
	"parkinglot/vehicle"
	"testing"
)

func TestManagerCreation(t *testing.T) {
	m := NewOwner()
	assert.NotNil(t, m)
}

func TestAssignAttendant(t *testing.T) {
	m := NewOwner()
	a := attendant.NewAttendant()
	m.AssignAttendant(a)

	assert.Contains(t, m.attendants, a, "Expected attendant to be assigned to the owner")
}

func TestAssignParkingLotToMultipleAttendant(t *testing.T) {
	m := NewOwner()
	a1 := attendant.NewAttendant()
	a2 := attendant.NewAttendant()

	l, _ := parkinglot.NewParkingLot(10)

	m.AssignParkingLotToAttendant(a1, l)
	m.AssignParkingLotToAttendant(a2, l)

	r1 := "RJ-12-JI-1234"
	r2 := "RJ-12-JI-5678"
	v1 := vehicle.NewVehicle(r1, vehicle.Red)
	v2 := vehicle.NewVehicle(r2, vehicle.Blue)

	a1.Park(v1)
	a2.Park(v2)

	assert.True(t, l.IsVehicleParked(r1))
	assert.True(t, l.IsVehicleParked(r2))
}

func TestNotifyOwnerWhenParkingLotFull(t *testing.T) {
	owner := NewOwner()

	l, _ := parkinglot.NewParkingLot(2)

	l.AddObserver(owner)

	r1 := "RJ-12-JI-1234"
	r2 := "RJ-12-JI-5678"
	v1 := vehicle.NewVehicle(r1, vehicle.Red)
	v2 := vehicle.NewVehicle(r2, vehicle.Blue)

	_, _ = l.Park(v1)
	ticket, _ := l.Park(v2)
	l.UnPark(ticket)
}
