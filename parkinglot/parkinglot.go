package parkinglot

import (
	"parkinglot/errors"
)

type ParkingLot struct {
}

func NewParkingLot(size int) (*ParkingLot, error) {
	if size <= 0 {
		return nil, errors.ErrInvalidSlotSize
	}
	return &ParkingLot{}, nil
}
