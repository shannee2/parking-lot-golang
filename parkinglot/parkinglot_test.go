package parkinglot

import (
	"parkinglot/errors"
	"parkinglot/vehicle"
	"testing"
)

func TestThrowException_WhenSlotSizeIsNegative(t *testing.T) {
	_, err := NewParkingLot(-1)

	if err != errors.ErrInvalidSlotSize {
		t.Error("Expected ErrInvalidSlotSize when slot size is negative, but nothing thrown")
	}
}

func TestThrowException_WhenSlotSizeIs0(t *testing.T) {
	_, err := NewParkingLot(0)

	if err != errors.ErrInvalidSlotSize {
		t.Error("Expected ErrInvalidSlotSize when slot size is 0, but nothing thrown")
	}
}

func TestParkingLotCreation(t *testing.T) {
	_, err := NewParkingLot(1)

	if err != nil {
		t.Error("Expected successful ParkingLot creation, but error thrown")
	}
}

func TestParkVehicle_AvailableSlot(t *testing.T) {
	parkingLot, _ := NewParkingLot(2)
	registrationNumber := "UJ-12-HG-3847"
	vehicle := vehicle.NewVehicle(registrationNumber, vehicle.Red)

	parkingLot.Park(vehicle)
	isVehicleParked := parkingLot.IsVehicleParked(registrationNumber)

	if !isVehicleParked {
		t.Error("Expected vehicle to be parked, but not parked")
	}
}

func TestMulitpleVehiclePark_InAvailableSlot(t *testing.T) {
	parkingLot, _ := NewParkingLot(3)
	firstRegNum := "UJ-12-HG-3847"
	secondRegNum := "DJ-79-DH-2938"
	thirdRegNum := "MP-13-UH-9098"
	firstVehicle := vehicle.NewVehicle(firstRegNum, vehicle.Red)
	secondVehicle := vehicle.NewVehicle(secondRegNum, vehicle.Red)
	thirdVehicle := vehicle.NewVehicle(thirdRegNum, vehicle.Red)

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
	parkingLot, _ := NewParkingLot(2)
	firstRegNum := "UJ-12-HG-3847"
	secondRegNum := "DJ-79-DH-2938"
	thirdRegNum := "MP-13-UH-9098"
	firstVehicle := vehicle.NewVehicle(firstRegNum, vehicle.Red)
	secondVehicle := vehicle.NewVehicle(secondRegNum, vehicle.Red)
	thirdVehicle := vehicle.NewVehicle(thirdRegNum, vehicle.Red)

	parkingLot.Park(firstVehicle)
	parkingLot.Park(secondVehicle)
	_, err := parkingLot.Park(thirdVehicle)

	if err != errors.ErrAllSlotsOccupied {
		t.Error("Expected to throw error when all slots occupied, but nothing thrown", err)
	}
}
