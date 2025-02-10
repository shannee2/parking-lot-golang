package owner

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"parkinglot/attendant"
	"parkinglot/parkinglot"
	"parkinglot/vehicle"
	"testing"
)

func TestManagerCreation(t *testing.T) {
	m := New()
	assert.NotNil(t, m)
}

func TestAssignAttendant(t *testing.T) {
	m := New()
	a := attendant.NewAttendant()
	m.AssignAttendant(a)

	assert.Contains(t, m.attendants, a, "Expected attendant to be assigned to the owner")
}

func TestAssignParkingLotToMultipleAttendant(t *testing.T) {
	m := New()
	a1 := attendant.NewAttendant()
	a2 := attendant.NewAttendant()

	l, _ := parkinglot.New(10)

	m.AssignParkingLotToAttendant(a1, l)
	m.AssignParkingLotToAttendant(a2, l)

	r1 := "RJ-12-JI-1234"
	r2 := "RJ-12-JI-5678"
	v1 := vehicle.New(r1, vehicle.Red)
	v2 := vehicle.New(r2, vehicle.Blue)

	a1.Park(v1)
	a2.Park(v2)

	assert.True(t, l.IsVehicleParked(r1))
	assert.True(t, l.IsVehicleParked(r2))
}

func TestNotifyOwnerWhenParkingLotFull(t *testing.T) {
	owner := New()

	l, _ := parkinglot.New(2)

	l.AddObserver(owner)

	r1 := "RJ-12-JI-1234"
	r2 := "RJ-12-JI-5678"
	v1 := vehicle.New(r1, vehicle.Red)
	v2 := vehicle.New(r2, vehicle.Blue)

	_, _ = l.Park(v1)
	ticket, _ := l.Park(v2)
	l.UnPark(ticket)
}

type MockOwner struct {
	Owner
	mock.Mock
}

func (m *MockOwner) OnParkingLotFull(l *parkinglot.ParkingLot) {
	m.Called(l)
}

func (m *MockOwner) OnParkingLotAvailable(l *parkinglot.ParkingLot) {
	m.Called(l)
}

func TestOwnerGetsNotifiedWhenParkingLotIsFull(t *testing.T) {
	mockOwner := new(MockOwner)

	parkingLot, _ := parkinglot.New(2)

	parkingLot.AddObserver(mockOwner)

	mockOwner.On("OnParkingLotFull", parkingLot).Once()
	mockOwner.On("OnParkingLotAvailable", parkingLot).Once()

	v1 := vehicle.New("RJ-12-JI-1234", vehicle.Red)
	v2 := vehicle.New("RJ-12-JI-5678", vehicle.Blue)

	_, _ = parkingLot.Park(v1)
	ticket, _ := parkingLot.Park(v2)
	_ = parkingLot.UnPark(ticket)

	mockOwner.AssertExpectations(t)
}

func TestOwnerNotNotifiedWhenParkingLotIsNotFull(t *testing.T) {
	mockOwner := new(MockOwner)
	parkingLot, _ := parkinglot.New(3)

	parkingLot.AddObserver(mockOwner)

	v1 := vehicle.New("RJ-12-JI-1234", vehicle.Red)
	v2 := vehicle.New("RJ-12-JI-5678", vehicle.Blue)

	_, _ = parkingLot.Park(v1)
	_, _ = parkingLot.Park(v2)

	mockOwner.AssertNotCalled(t, "OnParkingLotFull")

	mockOwner.AssertNotCalled(t, "OnParkingLotAvailable")
}

func TestAssignParkingLotToOwner(t *testing.T) {
	owner := New()
	parkingLot, _ := parkinglot.New(10)

	owner.AssignParkingLot(parkingLot)

	assert.Contains(t, owner.parkingLots, parkingLot, "Expected parking lot to be assigned to the owner")
}

func TestOwnerCanParkVehicles(t *testing.T) {
	owner := New()
	parkingLot, _ := parkinglot.New(2)

	owner.AssignParkingLot(parkingLot)

	r1 := "RJ-12-JI-1234"
	v1 := vehicle.New(r1, vehicle.Red)
	r2 := "RJ-12-JI-5678"
	v2 := vehicle.New(r2, vehicle.Blue)

	ticket1, err1 := owner.Park(v1)
	ticket2, err2 := owner.Park(v2)

	assert.NotNil(t, ticket1, "Expected to receive a ticket for the first vehicle")
	assert.Nil(t, err1, "Expected no error for the first vehicle parking")
	assert.NotNil(t, ticket2, "Expected to receive a ticket for the second vehicle")
	assert.Nil(t, err2, "Expected no error for the second vehicle parking")
	assert.True(t, parkingLot.IsVehicleParked(r1), "Expected the first vehicle to be parked")
	assert.True(t, parkingLot.IsVehicleParked(r2), "Expected the second vehicle to be parked")
}
