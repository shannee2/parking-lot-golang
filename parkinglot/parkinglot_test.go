package parkinglot

import (
	"parkinglot/errors"
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
