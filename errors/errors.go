package errors

import (
	"errors"
)

var ErrSlotAlreadyOccupied = errors.New("slot is already occupied")

var ErrSlotEmpty = errors.New("slot is empty")

var ErrInvalidSlotSize = errors.New("slot size must be greater than 0")

var ErrAllSlotsOccupied = errors.New("all slots are occupied")

var ErrTicketNotFound = errors.New("ticket does not found")

var ErrInvalidParkingLot = errors.New("invalid parking lot")

var ErrAllLotsAreFull = errors.New("all lots are full")

var ErrNoParkingLotAssignedToAttendant = errors.New("no parking lot assigned to attendant")
