package Slot

import (
	CustomErrors "ParkingLotGo"
	"ParkingLotGo/Car"
	"ParkingLotGo/Ticket"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Tests for IsEmpty() method
func TestIsEmpty(t *testing.T) {
	slot := NewSlot()

	t.Run("if slot is empty", func(t *testing.T) {
		assert.True(t, slot.IsEmpty())
	})

	t.Run("if slot is not empty", func(t *testing.T) {
		car := Car.NewCar("TS-1234", Car.RED)

		slot.Park(car)

		assert.False(t, slot.IsEmpty())
	})

	t.Run("if slot is empty after park and unpark", func(t *testing.T) {
		car := Car.NewCar("TS-1234", Car.RED)
		slot.Park(car)
		assert.False(t, slot.IsEmpty())

		_, _ = slot.Unpark(slot.ticket)

		assert.True(t, slot.IsEmpty())
	})
}

// Tests for Park() method
func TestPark(t *testing.T) {
	t.Run("park successfully", func(t *testing.T) {
		slot := NewSlot()
		car := Car.NewCar("TS-1234", Car.RED)

		ticket := slot.Park(car)

		assert.NotNil(t, ticket)
	})
}

// Tests for Unpark() method
func TestUnpark(t *testing.T) {
	slot := NewSlot()
	car := Car.NewCar("TS-1234", Car.RED)

	t.Run("if ticket is valid", func(t *testing.T) {
		ticket := slot.Park(car)

		actualCar, _ := slot.Unpark(ticket)

		assert.Equal(t, car, actualCar)
	})

	t.Run("if ticket is invalid", func(t *testing.T) {
		slot.Park(car)
		invalidTicket := Ticket.NewTicket()

		_, err := slot.Unpark(invalidTicket)

		assert.Error(t, err)
		assert.Equal(t, CustomErrors.ErrInvalidTicket, err)
	})
}

// Tests for IfCarIsParked() method
func TestIsCarParked(t *testing.T) {
	slot := NewSlot()
	car := Car.NewCar("TS-1234", Car.RED)

	t.Run("if car is parked", func(t *testing.T) {
		slot.Park(car)

		err := slot.IsCarParked(car)

		assert.Error(t, err)
		assert.Equal(t, CustomErrors.ErrCarAlreadyParked, err)
	})

	t.Run("if car is not parked", func(t *testing.T) {
		slot := NewSlot()
		car := Car.NewCar("TS-1234", Car.RED)

		err := slot.IsCarParked(car)

		assert.NoError(t, err)
	})
}

// Tests for IsCarOfColor() method
func TestIsCarOfColor(t *testing.T) {
	slot := NewSlot()
	car := Car.NewCar("TS-1234", Car.RED)

	t.Run("if car is of same color", func(t *testing.T) {
		slot.Park(car)

		assert.True(t, slot.IsCarOfColor(Car.RED))
	})

	t.Run("if car is of different color", func(t *testing.T) {
		slot.Park(car)

		assert.False(t, slot.IsCarOfColor(Car.BLUE))
	})
}

// Tests for HasSameRegistrationNumber() method
func TestHasSameRegistrationNumber(t *testing.T) {
	slot := NewSlot()
	car := Car.NewCar("TS-1234", Car.RED)

	t.Run("if car registration numbers are same", func(t *testing.T) {
		slot.Park(car)

		assert.True(t, slot.HasSameRegistrationNumber("TS-1234"))
	})

	t.Run("if car registration numbers are different", func(t *testing.T) {
		slot.Park(car)

		assert.False(t, slot.HasSameRegistrationNumber("TS-1235"))
	})
}
