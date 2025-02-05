package errors

import (
	"errors"
)

var ErrSlotAlreadyOccupied = errors.New("slot is already occupied")
var ErrSlotEmpty = errors.New("slot is already occupied")
