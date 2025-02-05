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

	if slot.isOccupied() {
		t.Error("Expected slot to be unoccupied when not parked a vehicle, but it is occupied")
	}
}

func TestSlotOccupiedWhenParkedVehicle(t *testing.T) {
	slot := NewSlot()

	slot.Park(vehicle.Vehicle{RegistrationNumber: "UJ-12-HG-3847", Color: vehicle.Red})

	if !slot.isOccupied() {
		t.Error("Expected slot to be occupied, but it is unoccupied")
	}
}

func TestUnparkVehicle(t *testing.T) {
	slot := NewSlot()

	slot.Park(vehicle.Vehicle{RegistrationNumber: "UJ-12-HG-3847", Color: vehicle.Red})
	slot.UnPark()

	if slot.isOccupied() {
		t.Error("Expected slot to be unoccupied when unparked a vehicle, but it is occupied")
	}
}

func TestThrowError_WhenParkingInOccupiedSlot(t *testing.T) {
	slot := NewSlot()
	firstVehicle := vehicle.Vehicle{RegistrationNumber: "UJ-12-HG-3847", Color: vehicle.Red}
	secondVehicle := vehicle.Vehicle{RegistrationNumber: "UJ-12-HG-1234", Color: vehicle.Blue}

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
	vechicleWithRedColor := vehicle.Vehicle{RegistrationNumber: "UJ-12-HG-3847", Color: vehicle.Red}
	slot.Park(vechicleWithRedColor)

	hasRedColorVehicle := slot.HasVehicleColor(vehicle.Red)

	if hasRedColorVehicle == false {
		t.Error("Expected vehicle color to be red, but not found")
	}
}
