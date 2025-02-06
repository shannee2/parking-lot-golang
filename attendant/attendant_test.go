package attendant

import (
	"github.com/stretchr/testify/assert"
	"parkinglot/parkinglot"
	"testing"
)

func TestAttendantCreation(t *testing.T) {
	attendant := NewAttendant()
	assert.NotNil(t, attendant)
}

func TestAssignParkingLotToAttendant(t *testing.T) {
	att := NewAttendant()
	l, _ := parkinglot.NewParkingLot(1)
	err := att.AssignParkingLot(l)
	assert.Nil(t, err)
}
