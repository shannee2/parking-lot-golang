package attendant

import (
	"github.com/stretchr/testify/assert"
	"parkinglot/parkinglot"
	"parkinglot/vehicle"
	"testing"
)

func TestEvenDistributionOfVehicles_When2ParkingLots_And2Vehicles(t *testing.T) {
	att := NewSmartAttendant()
	l1, _ := parkinglot.NewParkingLot(2)
	l2, _ := parkinglot.NewParkingLot(2)
	att.AssignParkingLot(l1)
	att.AssignParkingLot(l2)

	v1 := vehicle.NewVehicle("RJ-12-JI-1234", vehicle.Red)
	v2 := vehicle.NewVehicle("RJ-12-JI-5678", vehicle.Blue)

	att.Park(v1)
	att.Park(v2)

	lot1Vehicles := att.ParkingLots[0].CountParkedVehicles()
	lot2Vehicles := att.ParkingLots[1].CountParkedVehicles()

	assert.Equal(t, 1, lot1Vehicles)
	assert.Equal(t, 1, lot2Vehicles)

}

func TestEvenDistributionOfVehicles_When3ParkingLots_And9Vehicles(t *testing.T) {
	att := NewSmartAttendant()
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

	lot1Vehicles := att.ParkingLots[0].CountParkedVehicles()
	lot2Vehicles := att.ParkingLots[1].CountParkedVehicles()
	lot3Vehicles := att.ParkingLots[2].CountParkedVehicles()

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

	lot1Vehicles = att.ParkingLots[0].CountParkedVehicles()
	lot2Vehicles = att.ParkingLots[1].CountParkedVehicles()
	lot3Vehicles = att.ParkingLots[2].CountParkedVehicles()

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

	lot1Vehicles = att.ParkingLots[0].CountParkedVehicles()
	lot2Vehicles = att.ParkingLots[1].CountParkedVehicles()
	lot3Vehicles = att.ParkingLots[2].CountParkedVehicles()

	assert.Equal(t, 4, lot1Vehicles)
	assert.Equal(t, 3, lot2Vehicles)
	assert.Equal(t, 2, lot3Vehicles)
}
