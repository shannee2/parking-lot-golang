package attendant

import (
	"math/rand"
	"parkinglot/errors"
	"parkinglot/parkinglot"
	"time"
)

type ParkingStrategy interface {
	SelectLot(parkingLots []*parkinglot.ParkingLot) (*parkinglot.ParkingLot, error)
}

type RandomStrategy struct{}
type SmartStrategy struct{}
type NormalStrategy struct{}

func (s *SmartStrategy) SelectLot(parkingLots []*parkinglot.ParkingLot) (*parkinglot.ParkingLot, error) {
	var selectedLot *parkinglot.ParkingLot
	for _, lot := range parkingLots {
		if !lot.IsFull() {
			if selectedLot == nil || lot.CountParkedVehicles() < selectedLot.CountParkedVehicles() {
				selectedLot = lot
			}
		}
	}
	if selectedLot == nil {
		return nil, errors.ErrAllLotsAreFull
	}
	return selectedLot, nil
}

func (r *RandomStrategy) SelectLot(parkingLots []*parkinglot.ParkingLot) (*parkinglot.ParkingLot, error) {
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

func (n *NormalStrategy) SelectLot(parkingLots []*parkinglot.ParkingLot) (*parkinglot.ParkingLot, error) {
	for _, lot := range parkingLots {
		if !lot.IsFull() {
			return lot, nil
		}
	}
	return nil, errors.ErrAllLotsAreFull
}

func GenerateRandomNumber(n int) int {
	if n <= 0 {
		return 0
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(n)
}
