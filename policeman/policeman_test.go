package policeman

import (
	"github.com/stretchr/testify/mock"
	"parkinglot/parkinglot"
	"parkinglot/vehicle"
	"testing"
)

type MockPoliceMan struct {
	PoliceMan
	mock.Mock
}

func (p *MockPoliceMan) OnParkingLotFull(l *parkinglot.ParkingLot) {
	p.Called(l)
}

func (p *MockPoliceMan) OnParkingLotAvailable(l *parkinglot.ParkingLot) {
	p.Called(l)
}

func TestOwnerGetsNotifiedWhenParkingLotIsFull(t *testing.T) {
	mockPoliceman := new(MockPoliceMan)

	parkingLot, _ := parkinglot.New(2)

	parkingLot.AddObserver(mockPoliceman)

	mockPoliceman.On("OnParkingLotFull", parkingLot).Once()
	mockPoliceman.On("OnParkingLotAvailable", parkingLot).Once()

	v1 := vehicle.New("RJ-12-JI-1234", vehicle.Red)
	v2 := vehicle.New("RJ-12-JI-5678", vehicle.Blue)

	_, _ = parkingLot.Park(v1)
	ticket, _ := parkingLot.Park(v2)
	_ = parkingLot.UnPark(ticket)

	mockPoliceman.AssertExpectations(t)
}

func TestOwnerNotNotifiedWhenParkingLotIsNotFull(t *testing.T) {
	mockPoliceman := new(MockPoliceMan)
	parkingLot, _ := parkinglot.New(3)

	parkingLot.AddObserver(mockPoliceman)

	v1 := vehicle.New("RJ-12-JI-1234", vehicle.Red)
	v2 := vehicle.New("RJ-12-JI-5678", vehicle.Blue)

	_, _ = parkingLot.Park(v1)
	_, _ = parkingLot.Park(v2)

	mockPoliceman.AssertNotCalled(t, "OnParkingLotFull")

	mockPoliceman.AssertNotCalled(t, "OnParkingLotAvailable")
}
