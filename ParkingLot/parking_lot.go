package ParkingLot

import (
	CustomErrors "ParkingLotGo"
	"ParkingLotGo/Car"
	"ParkingLotGo/Slot"
	"ParkingLotGo/Ticket"
	"fmt"
)

type ParkingLot struct {
	slots []*Slot.Slot
}

func NewParkingLot(capacity int) (*ParkingLot, error) {
	if capacity <= 0 {
		return nil, fmt.Errorf("%w: capacity should be greater than 0", CustomErrors.ErrCannotCreateParkingLot)
	}

	slots := make([]*Slot.Slot, capacity)
	for i := 0; i < capacity; i++ {
		slots[i] = &Slot.Slot{}
	}

	parkingLot := &ParkingLot{
		slots: slots,
	}

	return parkingLot, nil
}

func (parkingLot *ParkingLot) IsFull() bool {
	for i := 0; i < len(parkingLot.slots); i++ {
		if parkingLot.slots[i].IsEmpty() {
			return false
		}
	}
	return true
}

func (parkingLot *ParkingLot) CheckIfCarIsParked(car *Car.Car) error {
	for i := 0; i < len(parkingLot.slots); i++ {
		err := parkingLot.slots[i].IsCarParked(car)
		if err != nil {
			return err
		}
	}
	return nil
}

func (parkingLot *ParkingLot) Park(car *Car.Car) (*Ticket.Ticket, error) {
	if parkingLot.IsFull() {
		return nil, CustomErrors.ErrParkingLotFull
	}

	err := parkingLot.CheckIfCarIsParked(car)
	if err != nil {
		return nil, err
	}

	slot, err := parkingLot.getNearestSlot()
	if err != nil {
		return nil, err
	}

	ticket := slot.Park(car)
	return ticket, nil
}

func (parkingLot *ParkingLot) Unpark(ticket *Ticket.Ticket) (*Car.Car, error) {
	for i := 0; i < len(parkingLot.slots); i++ {
		car, err := parkingLot.slots[i].Unpark(ticket)
		if err == nil {
			return car, nil
		}
	}
	return nil, CustomErrors.ErrInvalidTicket
}

func (parkingLot *ParkingLot) getNearestSlot() (*Slot.Slot, error) {
	for i := 0; i < len(parkingLot.slots); i++ {
		if parkingLot.slots[i].IsEmpty() {
			return parkingLot.slots[i], nil
		}
	}
	return nil, CustomErrors.ErrParkingLotFull
}

func (parkingLot *ParkingLot) GetCarSlotNumberByRegistrationNumber(registrationNumber string) (int, error) {
	for i := 0; i < len(parkingLot.slots); i++ {
		if !parkingLot.slots[i].IsEmpty() && parkingLot.slots[i].HasSameRegistrationNumber(registrationNumber) {
			return i + 1, nil
		}
	}
	return 0, CustomErrors.ErrCarNotFound
}

func (parkingLot *ParkingLot) CountCarsByColor(color Car.CarColor) int {
	count := 0
	for i := 0; i < len(parkingLot.slots); i++ {
		if parkingLot.slots[i].IsCarOfColor(color) {
			count++
		}
	}
	return count
}

func (parkingLot *ParkingLot) CountAvailableSlots() int {
	count := 0
	for i := 0; i < len(parkingLot.slots); i++ {
		if parkingLot.slots[i].IsEmpty() {
			count++
		}
	}
	return count
}
