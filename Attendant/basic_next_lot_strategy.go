package Attendant

import (
	CustomErrors "ParkingLotGo"
	"ParkingLotGo/ParkingLot"
)

type BasicNextLotStrategy struct{}

func (strategy *BasicNextLotStrategy) GetNextLot(parkingLots []*ParkingLot.ParkingLot) (*ParkingLot.ParkingLot, error) {
	var selectedLot *ParkingLot.ParkingLot
	for _, parkingLot := range parkingLots {
		if !parkingLot.IsFull() {
			selectedLot = parkingLot
			break
		}
	}
	if selectedLot == nil {
		return nil, CustomErrors.ErrParkingLotFull
	}
	return selectedLot, nil
}
