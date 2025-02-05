package errors

import (
	"errors"
)

var ErrSlotAlreadyOccupied = errors.New("slot is already occupied")

var ErrSlotEmpty = errors.New("slot is empty")

var ErrInvalidSlotSize = errors.New("slot size must be greater than 0")
