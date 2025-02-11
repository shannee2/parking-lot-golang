package attendant

import (
	"math/rand"
	"parkinglot/errors"
	"parkinglot/parkinglot"
	"time"
)

type ParkingStrategy interface {
	selectLot(parkingLots []*parkinglot.ParkingLot) (*parkinglot.ParkingLot, error)
}

type SequentialStrategy struct{}
type LeastOccupiedStrategy struct{}
type RandomStrategy struct{}
type MostOccupiedStrategy struct{}
type CircularStrategy struct {
	selectedIndex int
}

func (s *SequentialStrategy) selectLot(parkingLots []*parkinglot.ParkingLot) (*parkinglot.ParkingLot, error) {
	for _, lot := range parkingLots {
		if !lot.IsFull() {
			return lot, nil
		}
	}
	return nil, errors.ErrAllLotsAreFull
}

func (s *LeastOccupiedStrategy) selectLot(parkingLots []*parkinglot.ParkingLot) (*parkinglot.ParkingLot, error) {
	var selectedLot *parkinglot.ParkingLot
	for _, lot := range parkingLots {
		if !lot.IsFull() {
			if selectedLot == nil {
				selectedLot = lot
			} else {
				selectedLot = selectedLot.CompareLessOccupied(lot)
			}
		}
	}
	if selectedLot == nil {
		return nil, errors.ErrAllLotsAreFull
	}
	return selectedLot, nil
}

func (r *RandomStrategy) selectLot(parkingLots []*parkinglot.ParkingLot) (*parkinglot.ParkingLot, error) {
	if len(parkingLots) == 0 {
		return nil, errors.ErrAllLotsAreFull
	}
	for {
		selectedLot := parkingLots[GenerateRandomNumber(len(parkingLots))]
		if !selectedLot.IsFull() {
			return selectedLot, nil
		}
	}
}

func (m *MostOccupiedStrategy) selectLot(parkingLots []*parkinglot.ParkingLot) (*parkinglot.ParkingLot, error) {
	var selectedLot *parkinglot.ParkingLot
	for _, lot := range parkingLots {
		if !lot.IsFull() {
			if selectedLot == nil {
				selectedLot = lot
			} else {
				selectedLot = selectedLot.CompareMoreOccupied(lot)
			}
		}
	}
	if selectedLot == nil {
		return nil, errors.ErrAllLotsAreFull
	}
	return selectedLot, nil
}

func (c *CircularStrategy) selectLot(parkingLots []*parkinglot.ParkingLot) (*parkinglot.ParkingLot, error) {
	if len(parkingLots) == 0 {
		return nil, errors.ErrAllLotsAreFull
	}

	for {
		if !parkingLots[c.selectedIndex].IsFull() {
			selectedLot := parkingLots[c.selectedIndex]
			c.selectedIndex = (c.selectedIndex + 1) % len(parkingLots)
			return selectedLot, nil
		}
		c.selectedIndex = (c.selectedIndex + 1) % len(parkingLots)
	}
}

func GenerateRandomNumber(n int) int {
	if n <= 0 {
		return 0
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(n)
}
