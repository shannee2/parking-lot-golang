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

func TestThrowErrorWhenParking_IfNoParkingLotAssigned(t *testing.T) {
	att := NewAttendant()

	_, err := att.Park(vehicle.NewVehicle("RJ-12-JI-1234", vehicle.Red))
	assert.Equal(t, errors.ErrNoParkingLotAssignedToAttendant, err)
}

func TestDoesNotThrowErrorWhenParking_IfParkingLotAssigned(t *testing.T) {
	att := NewAttendant()
	l, _ := parkinglot.NewParkingLot(1)
	att.AssignParkingLot(l)

	_, err := att.Park(vehicle.NewVehicle("RJ-12-JI-1234", vehicle.Red))
	assert.NoError(t, err)
}

func TestEvenDistributionOfVehicles_When2ParkingLots_And2Vehicles(t *testing.T) {
	att := NewAttendant()
	l1, _ := parkinglot.NewParkingLot(2)
	l2, _ := parkinglot.NewParkingLot(2)
	att.AssignParkingLot(l1)
	att.AssignParkingLot(l2)

	v1 := vehicle.NewVehicle("RJ-12-JI-1234", vehicle.Red)
	v2 := vehicle.NewVehicle("RJ-12-JI-5678", vehicle.Blue)

	att.Park(v1)
	att.Park(v2)

	lot1Vehicles := att.parkingLots[0].CountParkedVehicles()
	lot2Vehicles := att.parkingLots[1].CountParkedVehicles()

	assert.Equal(t, 1, lot1Vehicles)
	assert.Equal(t, 1, lot2Vehicles)

}

func TestEvenDistributionOfVehicles_When3ParkingLots_And9Vehicles(t *testing.T) {
	att := NewAttendant()
	l1, _ := parkinglot.NewParkingLot(5)
	l2, _ := parkinglot.NewParkingLot(3)
	l3, _ := parkinglot.NewParkingLot(2)
	att.AssignParkingLot(l1)
	att.AssignParkingLot(l2)
	att.AssignParkingLot(l3)

	// park 3 vehicles

	v1 := vehicle.NewVehicle("RJ-12-JI-1234", vehicle.Red)
	v2 := vehicle.NewVehicle("RJ-12-JI-5678", vehicle.Blue)
	v3 := vehicle.NewVehicle("RJ-12-JI-5008", vehicle.Blue)

	att.Park(v1)
	att.Park(v2)
	att.Park(v3)

	lot1Vehicles := att.parkingLots[0].CountParkedVehicles()
	lot2Vehicles := att.parkingLots[1].CountParkedVehicles()
	lot3Vehicles := att.parkingLots[2].CountParkedVehicles()

	assert.Equal(t, 1, lot1Vehicles)
	assert.Equal(t, 1, lot2Vehicles)
	assert.Equal(t, 1, lot3Vehicles)

	// park 3 more vehicles

	v4 := vehicle.NewVehicle("RJ-12-JI-6009", vehicle.Green)
	v5 := vehicle.NewVehicle("RJ-12-JI-7001", vehicle.Blue)
	v6 := vehicle.NewVehicle("RJ-12-JI-8002", vehicle.Red)

	att.Park(v4)
	att.Park(v5)
	att.Park(v6)

	lot1Vehicles = att.parkingLots[0].CountParkedVehicles()
	lot2Vehicles = att.parkingLots[1].CountParkedVehicles()
	lot3Vehicles = att.parkingLots[2].CountParkedVehicles()

	assert.Equal(t, 2, lot1Vehicles)
	assert.Equal(t, 2, lot2Vehicles)
	assert.Equal(t, 2, lot3Vehicles)

	// park another 3 vehicles

	v7 := vehicle.NewVehicle("RJ-12-JI-9003", vehicle.Red)
	v8 := vehicle.NewVehicle("RJ-12-JI-1004", vehicle.Green)
	v9 := vehicle.NewVehicle("RJ-12-JI-1105", vehicle.Blue)

	att.Park(v7)
	att.Park(v8)
	att.Park(v9)

	lot1Vehicles = att.parkingLots[0].CountParkedVehicles()
	lot2Vehicles = att.parkingLots[1].CountParkedVehicles()
	lot3Vehicles = att.parkingLots[2].CountParkedVehicles()

	assert.Equal(t, 4, lot1Vehicles)
	assert.Equal(t, 3, lot2Vehicles)
	assert.Equal(t, 2, lot3Vehicles)
}
