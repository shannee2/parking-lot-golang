package parkinglot

import (
	"parkinglot/errors"
	"parkinglot/slot"
)

type ParkingLot struct {
	slots []*slot.Slot
}

func NewParkingLot(size int) (*ParkingLot, error) {
	if size <= 0 {
		return nil, errors.ErrInvalidSlotSize
	}

	parkingLot := &ParkingLot{
		slots: make([]*slot.Slot, size),
	}
	for i := 0; i < size; i++ {
		parkingLot.slots[i] = slot.NewSlot()
	}
	return &ParkingLot{}, nil
}
