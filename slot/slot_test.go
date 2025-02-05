package slot

import (
	"parkinglot/errors"
	"parkinglot/vehicle"
	"testing"
)

func TestSlotCreation(t *testing.T) {
	slot := NewSlot()

	if slot == nil {
		t.Error("Expected slot to be initialized, but it is nil")
	}
}

func TestSlotUnoccupiedInitially(t *testing.T) {
	slot := NewSlot()

	if slot.IsOccupied() {
		t.Error("Expected slot to be unoccupied when not parked a vehicle, but it is occupied")
	}
}

func TestSlotOccupiedWhenParkedVehicle(t *testing.T) {
	slot := NewSlot()

	slot.Park(vehicle.NewVehicle("UJ-12-HG-3847", vehicle.Red))

	if !slot.IsOccupied() {
		t.Error("Expected slot to be occupied, but it is unoccupied")
	}
}

func TestVehicleParked_WithRegistrationNumber(t *testing.T) {
	slot := NewSlot()
	registrationNumber := "UJ-12-HG-3847"

	slot.Park(vehicle.NewVehicle(registrationNumber, vehicle.Red))

	if !slot.IsVehicleParked(registrationNumber) {
		t.Error("Expected registration number to be parked, but it not found")
	}
}

func TestVehicleNotParked_WithRegistrationNumber(t *testing.T) {
	slot := NewSlot()

	slot.Park(vehicle.NewVehicle("UJ-12-HG-3847", vehicle.Red))

	if slot.IsVehicleParked("UJ-12-HG-1234") {
		t.Error("Expected registration number to be not parked, but it is parked")
	}
}

func TestUnparkVehicle(t *testing.T) {
	slot := NewSlot()

	slot.Park(vehicle.NewVehicle("UJ-12-HG-3847", vehicle.Red))
	slot.UnPark()

	if slot.IsOccupied() {
		t.Error("Expected slot to be unoccupied when unparked a vehicle, but it is occupied")
	}
}

func TestThrowError_WhenParkingInOccupiedSlot(t *testing.T) {
	slot := NewSlot()
	firstVehicle := vehicle.NewVehicle("UJ-12-HG-3847", vehicle.Red)
	secondVehicle := vehicle.NewVehicle("UJ-12-HG-1234", vehicle.Blue)

	_ = slot.Park(firstVehicle)
	err := slot.Park(secondVehicle)

	if err != errors.ErrSlotAlreadyOccupied {
		t.Error("Expected slot to throw exception when parking in occupied slot, but nothing thrown")
	}
}

func TestThrowError_WhenUnParkingInEmptySlot(t *testing.T) {
	slot := NewSlot()

	err := slot.UnPark()

	if err != errors.ErrSlotEmpty {
		t.Error("Expected slot to throw exception when parking in occupied slot, but nothing thrown")
	}
}

func TestSlotHasVehicleWithRedColor(t *testing.T) {
	slot := NewSlot()
	vechicleWithRedColor := vehicle.NewVehicle("UJ-12-HG-3847", vehicle.Red)
	slot.Park(vechicleWithRedColor)

	hasRedColorVehicle := slot.HasVehicleColor(vehicle.Red)

	if hasRedColorVehicle == false {
		t.Error("Expected vehicle color to be red, but not found")
	}
}
