package Attendant

import (
	"ParkingLotGo/ParkingLot"
)

type NextLotStrategy interface {
	GetNextLot(parkingLots []*ParkingLot.ParkingLot) (*ParkingLot.ParkingLot, error)
}
