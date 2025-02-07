package parkinglot

import (
	"github.com/stretchr/testify/assert"
	"parkinglot/errors"
	"parkinglot/ticket"
	"parkinglot/vehicle"
	"testing"
)

func TestThrowException_WhenSlotSizeIsNegative(t *testing.T) {
	_, err := New(-1)

	if err != errors.ErrInvalidSlotSize {
		t.Error("Expected ErrInvalidSlotSize when slot size is negative, but nothing thrown")
	}
}

func TestThrowException_WhenSlotSizeIs0(t *testing.T) {
	_, err := New(0)

	if err != errors.ErrInvalidSlotSize {
		t.Error("Expected ErrInvalidSlotSize when slot size is 0, but nothing thrown")
	}
}

func TestParkingLotCreation(t *testing.T) {
	_, err := New(1)

	if err != nil {
		t.Error("Expected successful ParkingLot creation, but error thrown")
	}
}

func TestParkVehicle_AvailableSlot(t *testing.T) {
	parkingLot, _ := New(2)
	registrationNumber := "UJ-12-HG-3847"
	v := vehicle.New(registrationNumber, vehicle.Red)

	parkingLot.Park(v)
	isVehicleParked := parkingLot.IsVehicleParked(registrationNumber)

	if !isVehicleParked {
		t.Error("Expected vehicle to be parked, but not parked")
	}
}

func TestMulitpleVehiclePark_InAvailableSlot(t *testing.T) {
	parkingLot, _ := New(3)
	firstRegNum := "UJ-12-HG-3847"
	secondRegNum := "DJ-79-DH-2938"
	thirdRegNum := "MP-13-UH-9098"
	firstVehicle := vehicle.New(firstRegNum, vehicle.Red)
	secondVehicle := vehicle.New(secondRegNum, vehicle.Red)
	thirdVehicle := vehicle.New(thirdRegNum, vehicle.Red)

	parkingLot.Park(firstVehicle)
	parkingLot.Park(secondVehicle)
	parkingLot.Park(thirdVehicle)

	isFirstVehicleParked := parkingLot.IsVehicleParked(firstRegNum)
	isSecondVehicleParked := parkingLot.IsVehicleParked(secondRegNum)
	isThirdVehicleParked := parkingLot.IsVehicleParked(thirdRegNum)

	if !isFirstVehicleParked || !isSecondVehicleParked || !isThirdVehicleParked {
		t.Error("All three vehicles are not parked", isFirstVehicleParked, isSecondVehicleParked, isThirdVehicleParked)
	}
}

func TestThrowExceptionWhenAllSlotsOccupied(t *testing.T) {
	parkingLot, _ := New(2)
	firstRegNum := "UJ-12-HG-3847"
	secondRegNum := "DJ-79-DH-2938"
	thirdRegNum := "MP-13-UH-9098"
	firstVehicle := vehicle.New(firstRegNum, vehicle.Red)
	secondVehicle := vehicle.New(secondRegNum, vehicle.Red)
	thirdVehicle := vehicle.New(thirdRegNum, vehicle.Red)

	parkingLot.Park(firstVehicle)
	parkingLot.Park(secondVehicle)
	_, err := parkingLot.Park(thirdVehicle)

	if err != errors.ErrAllSlotsOccupied {
		t.Error("Expected to throw error when all slots occupied, but nothing thrown", err)
	}
}

func TestUnparkVehicle(t *testing.T) {
	parkingLot, _ := New(2)
	registrationNumber := "UJ-12-HG-3847"
	vehicle := vehicle.New(registrationNumber, vehicle.Red)

	ticket, _ := parkingLot.Park(vehicle)
	parkingLot.UnPark(ticket)
	isVehicleParked := parkingLot.IsVehicleParked(registrationNumber)

	if isVehicleParked {
		t.Error("Expected vehicle to be unparked, but vehicle still parked")
	}
}

func TestThrowException_WhenUnparkingInEmptySlot(t *testing.T) {
	parkingLot, _ := New(2)
	invalidTicket := ticket.New()

	err := parkingLot.UnPark(invalidTicket)

	if err != errors.ErrTicketNotFound {
		t.Error("Expected ErrTicketNotFound when unparking with an invalid ticket, but got a different error or no error")
	}
}

func TestParkingVehicle_WhenAnotherVehicleUnparks(t *testing.T) {
	parkingLot, _ := New(2)
	firstVehicle := vehicle.New("UJ-12-HG-3847", vehicle.Red)
	secondVehicle := vehicle.New("DJ-79-DH-2938", vehicle.Blue)
	thirdVehicle := vehicle.New("MP-13-UH-9098", vehicle.Green)

	firstTicket, _ := parkingLot.Park(firstVehicle)
	parkingLot.Park(secondVehicle)

	parkingLot.UnPark(firstTicket)

	_, err := parkingLot.Park(thirdVehicle)
	if err != nil {
		t.Errorf("Expected to park the third vehicle after unparking the first vehicle, but got error: %v", err)
	}

	isThirdVehicleParked := parkingLot.IsVehicleParked("MP-13-UH-9098")
	if !isThirdVehicleParked {
		t.Error("Expected the third vehicle to be parked, but it is not parked")
	}
}

func TestCountVehiclesWithRedColor(t *testing.T) {
	parkingLot, _ := New(5)

	firstVehicle := vehicle.New("UJ-12-HG-3847", vehicle.Red)
	secondVehicle := vehicle.New("DJ-79-DH-2938", vehicle.Blue)
	thirdVehicle := vehicle.New("MP-13-UH-9098", vehicle.Red)
	fourthVehicle := vehicle.New("KA-05-MH-1234", vehicle.Green)
	fifthVehicle := vehicle.New("TN-22-XY-5678", vehicle.Red)

	parkingLot.Park(firstVehicle)
	parkingLot.Park(secondVehicle)
	parkingLot.Park(thirdVehicle)
	parkingLot.Park(fourthVehicle)
	parkingLot.Park(fifthVehicle)

	redCount := parkingLot.countVehiclesWithColor(vehicle.Red)
	blueCount := parkingLot.countVehiclesWithColor(vehicle.Blue)
	greenCount := parkingLot.countVehiclesWithColor(vehicle.Green)

	if redCount != 3 {
		t.Errorf("Expected 3 red vehicles, but got %d", redCount)
	}

	if blueCount != 1 {
		t.Errorf("Expected 1 blue vehicle, but got %d", blueCount)
	}

	if greenCount != 1 {
		t.Errorf("Expected 1 green vehicle, but got %d", greenCount)
	}
}

func TestCountParkedVehicles_When2ParkedVehicles(t *testing.T) {
	parkingLot, _ := New(5)

	firstVehicle := vehicle.New("UJ-12-HG-3847", vehicle.Red)
	secondVehicle := vehicle.New("DJ-79-DH-2938", vehicle.Blue)

	parkingLot.Park(firstVehicle)
	parkingLot.Park(secondVehicle)

	count := parkingLot.CountParkedVehicles()
	assert.Equal(t, 2, count, "Expected 2 parked vehicles, but got %d", count)
}

func TestCountParkedVehicles_When5ParkedVehicles(t *testing.T) {
	parkingLot, _ := New(5)

	firstVehicle := vehicle.New("UJ-12-HG-3847", vehicle.Red)
	secondVehicle := vehicle.New("DJ-79-DH-2938", vehicle.Blue)
	thirdVehicle := vehicle.New("MP-13-UH-9098", vehicle.Green)
	fourthVehicle := vehicle.New("KA-05-MH-1234", vehicle.Red)
	fifthVehicle := vehicle.New("TN-22-XY-5678", vehicle.Blue)

	parkingLot.Park(firstVehicle)
	parkingLot.Park(secondVehicle)
	parkingLot.Park(thirdVehicle)
	parkingLot.Park(fourthVehicle)
	parkingLot.Park(fifthVehicle)

	count := parkingLot.CountParkedVehicles()
	assert.Equal(t, 5, count, "Expected 5 parked vehicles, but got %d", count)
}

func TestIsFull_WhenParkingLotIsFull(t *testing.T) {
	parkingLot, _ := New(2)
	firstVehicle := vehicle.New("UJ-12-HG-3847", vehicle.Red)
	secondVehicle := vehicle.New("DJ-79-DH-2938", vehicle.Blue)

	parkingLot.Park(firstVehicle)
	parkingLot.Park(secondVehicle)

	isFull := parkingLot.IsFull()
	assert.True(t, isFull, "Expected parking lot to be full, but it is not")
}

func TestIsFull_WhenParkingLotIsNotFull(t *testing.T) {
	parkingLot, _ := New(2)
	firstVehicle := vehicle.New("UJ-12-HG-3847", vehicle.Red)

	parkingLot.Park(firstVehicle)

	isFull := parkingLot.IsFull()
	assert.False(t, isFull, "Expected parking lot to not be full, but it is")
}
