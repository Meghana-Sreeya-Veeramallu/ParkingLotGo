package Attendant

import (
	CustomErrors "ParkingLotGo"
	"ParkingLotGo/Car"
	"ParkingLotGo/ParkingLot"
	"ParkingLotGo/Ticket"
)

type Attendant struct {
	nextLotStrategy NextLotStrategy
	ParkingLots     []*ParkingLot.ParkingLot
}

func NewAttendant(strategy NextLotStrategy) *Attendant {
	return &Attendant{
		nextLotStrategy: strategy,
		ParkingLots:     []*ParkingLot.ParkingLot{},
	}
}

func (attendant *Attendant) Assign(parkingLot *ParkingLot.ParkingLot) error {
	for _, lot := range attendant.ParkingLots {
		if lot == parkingLot {
			return CustomErrors.ErrParkingLotAlreadyAssigned
		}
	}

	attendant.ParkingLots = append(attendant.ParkingLots, parkingLot)
	return nil
}

func (attendant *Attendant) CheckIfCarIsParked(car *Car.Car) error {
	for _, parkingLot := range attendant.ParkingLots {
		err := parkingLot.CheckIfCarIsParked(car)
		if err != nil {
			return err
		}
	}
	return nil
}

func (attendant *Attendant) Park(car *Car.Car) (*Ticket.Ticket, error) {
	if len(attendant.ParkingLots) == 0 {
		return nil, CustomErrors.ErrNoParkingLotAssigned
	}

	err := attendant.CheckIfCarIsParked(car)
	if err != nil {
		return nil, err
	}

	selectedLot, err := attendant.nextLotStrategy.GetNextLot(attendant.ParkingLots)
	if err != nil {
		return nil, err
	}

	return selectedLot.Park(car)
}

func (attendant *Attendant) Unpark(ticket *Ticket.Ticket) (*Car.Car, error) {
	for _, parkingLot := range attendant.ParkingLots {
		car, err := parkingLot.Unpark(ticket)
		if err == nil {
			return car, nil
		}
	}
	return nil, CustomErrors.ErrInvalidTicket
}
