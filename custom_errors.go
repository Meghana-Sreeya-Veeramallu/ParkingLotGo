package CustomError

import "errors"

var (
	ErrInvalidTicket          = errors.New("ticket is invalid")
	ErrCarAlreadyParked       = errors.New("car is already parked")
	ErrCannotCreateParkingLot = errors.New("cannot create parking lot")
	ErrParkingLotFull         = errors.New("parking lot is full")
	ErrCarNotFound            = errors.New("car not found")
)
