package attendant

import (
	"fmt"
	"math/rand"
	"parkinglot/errors"
	"parkinglot/parkinglot"
	"parkinglot/ticket"
	"parkinglot/vehicle"
	"time"
)

type RandomAttendant struct {
	Attendant
}

func GenerateRandomNumber(n int) int {
	if n <= 0 {
		return 0
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(n)
}

func (a *RandomAttendant) Park(vehicle *vehicle.Vehicle) (*ticket.Ticket, error) {
	totalLots := len(a.ParkingLots)

	if totalLots == 0 {
		return nil, errors.ErrNoParkingLotAssignedToAttendant
	}

	if a.allLotsOccupied() {
		return nil, errors.ErrAllLotsAreFull
	}

	selectedLot, err := a.findAvailableLot()
	if err != nil {
		return nil, err
	}

	return a.parkAtRandomSlot(selectedLot, vehicle)
}

func (a *RandomAttendant) findAvailableLot() (*parkinglot.ParkingLot, error) {
	totalLots := len(a.ParkingLots)
	for {
		selectedLotIndex := GenerateRandomNumber(totalLots)
		selectedLot := a.ParkingLots[selectedLotIndex]
		if !selectedLot.IsFull() {
			return selectedLot, nil
		}
	}
}

func (a *RandomAttendant) parkAtRandomSlot(lot *parkinglot.ParkingLot, vehicle *vehicle.Vehicle) (*ticket.Ticket, error) {
	for {
		selectedSlotIndex := GenerateRandomNumber(lot.TotalSlots())
		t, err := lot.ParkInSlot(vehicle, selectedSlotIndex)
		if err == nil {
			return t, nil
		}
		if err != errors.ErrSlotAlreadyOccupied {
			return nil, err
		}
	}
}

func (a *RandomAttendant) PrintSlots() {
	for i, lot := range a.ParkingLots {
		fmt.Println("Lot", i+1)
		lot.Display()
		fmt.Println()
	}
}

func (a *RandomAttendant) allLotsOccupied() bool {
	for _, lot := range a.ParkingLots {
		if !lot.IsFull() {
			return false
		}
	}
	return true
}

func NewRandomAttendant() *RandomAttendant {
	return &RandomAttendant{
		Attendant: Attendant{
			ParkingLots: make([]*parkinglot.ParkingLot, 0),
		},
	}
}
