package CustomError

import "errors"

var (
	ErrInvalidTicket    = errors.New("ticket is invalid")
	ErrCarAlreadyParked = errors.New("car is already parked")
)
