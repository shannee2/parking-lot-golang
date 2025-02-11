package traffic_cop

import (
	"github.com/stretchr/testify/mock"
	"parkinglot/parkinglot"
	"parkinglot/vehicle"
	"testing"
)

type MockTrafficCop struct {
	TrafficCop
	mock.Mock
}

func (p *MockTrafficCop) OnParkingLotFull(l *parkinglot.ParkingLot) {
	p.Called(l)
}

func (p *MockTrafficCop) OnParkingLotAvailable(l *parkinglot.ParkingLot) {
	p.Called(l)
}

func TestOwnerGetsNotifiedWhenParkingLotIsFull(t *testing.T) {
	mockTrafficCop := new(MockTrafficCop)

	parkingLot, _ := parkinglot.New(2)

	parkingLot.AddObserver(mockTrafficCop)

	mockTrafficCop.On("OnParkingLotFull", parkingLot).Once()
	mockTrafficCop.On("OnParkingLotAvailable", parkingLot).Once()

	v1 := vehicle.New("RJ-12-JI-1234", vehicle.Red)
	v2 := vehicle.New("RJ-12-JI-5678", vehicle.Blue)

	_, _ = parkingLot.Park(v1)
	ticket, _ := parkingLot.Park(v2)
	_ = parkingLot.UnPark(ticket)

	mockTrafficCop.AssertExpectations(t)
}

func TestOwnerNotNotifiedWhenParkingLotIsNotFull(t *testing.T) {
	mockTrafficCop := new(MockTrafficCop)
	parkingLot, _ := parkinglot.New(3)

	parkingLot.AddObserver(mockTrafficCop)

	v1 := vehicle.New("RJ-12-JI-1234", vehicle.Red)
	v2 := vehicle.New("RJ-12-JI-5678", vehicle.Blue)

	_, _ = parkingLot.Park(v1)
	_, _ = parkingLot.Park(v2)

	mockTrafficCop.AssertNotCalled(t, "OnParkingLotFull")

	mockTrafficCop.AssertNotCalled(t, "OnParkingLotAvailable")
}
