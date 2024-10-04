package Slot

import (
	CustomErrors "ParkingLotGo"
	"ParkingLotGo/Car"
	"ParkingLotGo/Ticket"
)

type Slot struct {
	car    *Car.Car
	ticket *Ticket.Ticket
}

func NewSlot() *Slot {
	return &Slot{
		car:    nil,
		ticket: nil,
	}
}

func (slot *Slot) IsEmpty() bool {
	return slot.car == nil
}

func (slot *Slot) Park(car *Car.Car) *Ticket.Ticket {
	slot.car = car
	slot.ticket = Ticket.NewTicket()
	return slot.ticket
}

func (slot *Slot) Unpark(ticket *Ticket.Ticket) (*Car.Car, error) {
	if slot.ticket != nil && slot.ticket.IsSameTicket(ticket) {
		car := slot.car
		slot.car = nil
		slot.ticket = nil
		return car, nil
	}
	return nil, CustomErrors.ErrInvalidTicket
}

func (slot *Slot) IsCarIsParked(car *Car.Car) error {
	if !slot.IsEmpty() && slot.car.IsSameCar(car) {
		return CustomErrors.ErrCarAlreadyParked
	}
	return nil
}

func (slot *Slot) IsCarOfColor(color Car.CarColor) bool {
	return !slot.IsEmpty() && slot.car.IsSameColor(color)
}

func (slot *Slot) HasSameRegistrationNumber(registrationNumber string) bool {
	return !slot.IsEmpty() && slot.car.HasSameRegistrationNumber(registrationNumber)
}
