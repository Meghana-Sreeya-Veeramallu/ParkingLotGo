package Attendant

import (
	CustomErrors "ParkingLotGo"
	"ParkingLotGo/ParkingLot"
)

type SmartNextLotStrategy struct{}

func (s *SmartNextLotStrategy) GetNextLot(parkingLots []*ParkingLot.ParkingLot) (*ParkingLot.ParkingLot, error) {
	var selectedLot *ParkingLot.ParkingLot
	maxCapacityLeft := -1

	for _, parkingLot := range parkingLots {
		availableSlots := parkingLot.CountAvailableSlots()
		if availableSlots > maxCapacityLeft {
			maxCapacityLeft = availableSlots
			selectedLot = parkingLot
		}
	}

	if selectedLot == nil || selectedLot.IsFull() {
		return nil, CustomErrors.ErrParkingLotFull
	}
	return selectedLot, nil
}
