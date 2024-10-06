package Attendant

import (
	"ParkingLotGo/Car"
	"ParkingLotGo/Ticket"
)

type Attendable interface {
	CheckIfCarIsParked(car Car.Car)
	Park(car Car.Car) Ticket.Ticket
	Unpark(ticket Ticket.Ticket) Car.Car
}
