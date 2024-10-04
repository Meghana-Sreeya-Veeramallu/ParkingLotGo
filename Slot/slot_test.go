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

	t.Run("IfSlotIsEmpty", func(t *testing.T) {
		assert.True(t, slot.IsEmpty())
	})

	t.Run("IfSlotIsNotEmpty", func(t *testing.T) {
		car := Car.NewCar("TS-1234", Car.RED)
		slot.Park(car)
		assert.False(t, slot.IsEmpty())
	})
}

// Tests for Park() method
func TestPark(t *testing.T) {
	t.Run("ParkSuccessfully", func(t *testing.T) {
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

	t.Run("IfTicketIsValid", func(t *testing.T) {
		ticket := slot.Park(car)

		actualCar, _ := slot.Unpark(ticket)
		assert.Equal(t, car, actualCar)
	})

	t.Run("IfTicketIsInvalid", func(t *testing.T) {
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

	t.Run("IfCarIsParked", func(t *testing.T) {
		slot.Park(car)

		err := slot.IsCarIsParked(car)

		assert.Error(t, err)
		assert.Equal(t, CustomErrors.ErrCarAlreadyParked, err)
	})

	t.Run("IfCarIsNotParked", func(t *testing.T) {
		slot := NewSlot()
		car := Car.NewCar("TS-1234", Car.RED)

		err := slot.IsCarIsParked(car)

		assert.NoError(t, err)
	})
}

// Tests for IsCarOfColor() method
func TestIsCarOfColor(t *testing.T) {
	slot := NewSlot()
	car := Car.NewCar("TS-1234", Car.RED)

	t.Run("IfCarIsOfSameColor", func(t *testing.T) {
		slot.Park(car)

		assert.True(t, slot.IsCarOfColor(Car.RED))
	})

	t.Run("IfCarIsOfDifferentColor", func(t *testing.T) {
		slot.Park(car)

		assert.False(t, slot.IsCarOfColor(Car.BLUE))
	})
}

// Tests for HasSameRegistrationNumber() method
func TestHasSameRegistrationNumber(t *testing.T) {
	slot := NewSlot()
	car := Car.NewCar("TS-1234", Car.RED)

	t.Run("IfCarRegistrationNumbersAreSame", func(t *testing.T) {
		slot.Park(car)

		assert.True(t, slot.HasSameRegistrationNumber("TS-1234"))
	})

	t.Run("IfCarRegistrationNumbersAreDifferent", func(t *testing.T) {
		slot.Park(car)

		assert.False(t, slot.HasSameRegistrationNumber("TS-1235"))
	})
}
