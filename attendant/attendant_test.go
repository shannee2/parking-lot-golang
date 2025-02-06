package attendant

import (
	"github.com/stretchr/testify/assert"
	"parkinglot/errors"
	"parkinglot/parkinglot"
	"parkinglot/vehicle"
	"testing"
)

func TestAttendantCreation(t *testing.T) {
	attendant := NewAttendant()
	assert.NotNil(t, attendant)
}

func TestAssignParkingLotToAttendant(t *testing.T) {
	att := NewAttendant()
	l, err := parkinglot.NewParkingLot(1)
	assert.NoError(t, err)
	att.AssignParkingLot(l)
	assert.Contains(t, att.parkingLots, l)
}

func TestParkVehicle(t *testing.T) {
	att := NewAttendant()
	l, err := parkinglot.NewParkingLot(1)
	assert.NoError(t, err)
	att.AssignParkingLot(l)
	registrationNumber := "RJ-12-JI-1234"
	att.Park(vehicle.NewVehicle(registrationNumber, vehicle.Red))
	parked := att.IsParked(registrationNumber)
	assert.True(t, parked)
}

func TestParkMultipleVehicle(t *testing.T) {
	att := NewAttendant()
	l, err := parkinglot.NewParkingLot(2)
	assert.NoError(t, err)
	att.AssignParkingLot(l)

	firstRegistrationNumber := "RJ-12-JI-1234"
	secondRegistrationNumber := "RJ-12-JI-5678"

	vehicle1 := vehicle.NewVehicle(firstRegistrationNumber, vehicle.Red)
	vehicle2 := vehicle.NewVehicle(secondRegistrationNumber, vehicle.Blue)

	ticket1, err := att.Park(vehicle1)
	assert.NoError(t, err)
	assert.NotNil(t, ticket1)

	ticket2, err := att.Park(vehicle2)
	assert.NoError(t, err)
	assert.NotNil(t, ticket2)

	assert.True(t, att.IsParked(firstRegistrationNumber))
	assert.True(t, att.IsParked(secondRegistrationNumber))
}

func TestParkVehicleInMultipleParkingLots(t *testing.T) {
	att := NewAttendant()
	l1, err := parkinglot.NewParkingLot(2)
	assert.NoError(t, err)
	l2, err := parkinglot.NewParkingLot(1)
	assert.NoError(t, err)
	att.AssignParkingLot(l1)
	att.AssignParkingLot(l2)

	vehicle1 := vehicle.NewVehicle("RJ-12-JI-1234", vehicle.Red)
	vehicle2 := vehicle.NewVehicle("RJ-12-JI-5678", vehicle.Blue)
	vehicle3 := vehicle.NewVehicle("RJ-12-JI-9101", vehicle.Green)

	ticket1, err := att.Park(vehicle1)
	assert.NoError(t, err)
	assert.NotNil(t, ticket1)

	ticket2, err := att.Park(vehicle2)
	assert.NoError(t, err)
	assert.NotNil(t, ticket2)

	ticket3, err := att.Park(vehicle3)
	assert.NoError(t, err)
	assert.NotNil(t, ticket3)

	assert.True(t, att.IsParked("RJ-12-JI-1234"))
	assert.True(t, att.IsParked("RJ-12-JI-5678"))
	assert.True(t, att.IsParked("RJ-12-JI-9101"))
}

func TestThrowError_WhenAllLotsAreFull(t *testing.T) {
	att := NewAttendant()
	l1, err := parkinglot.NewParkingLot(1)
	assert.NoError(t, err)
	l2, err := parkinglot.NewParkingLot(1)
	assert.NoError(t, err)
	att.AssignParkingLot(l1)
	att.AssignParkingLot(l2)

	vehicle1 := vehicle.NewVehicle("RJ-12-JI-1234", vehicle.Red)
	vehicle2 := vehicle.NewVehicle("RJ-12-JI-5678", vehicle.Blue)
	vehicle3 := vehicle.NewVehicle("RJ-12-JI-9101", vehicle.Green)

	_, err = att.Park(vehicle1)
	assert.NoError(t, err)

	_, err = att.Park(vehicle2)
	assert.NoError(t, err)

	_, err = att.Park(vehicle3)
	assert.Error(t, err)
	assert.Equal(t, errors.ErrAllLotsAreFull, err)
}

func TestUnparkVehicle(t *testing.T) {
	att := NewAttendant()
	l, err := parkinglot.NewParkingLot(2)
	assert.NoError(t, err)
	att.AssignParkingLot(l)

	registrationNumber := "RJ-12-JI-1234"
	vehicle := vehicle.NewVehicle(registrationNumber, vehicle.Red)
	ticket, err := att.Park(vehicle)
	assert.NoError(t, err)
	assert.NotNil(t, ticket)

	err = att.Unpark(ticket)
	assert.NoError(t, err)
	assert.False(t, att.IsParked(registrationNumber))
}
