package ParkingLot

import (
	CustomErrors "ParkingLotGo"
	"ParkingLotGo/Car"
	"ParkingLotGo/Ticket"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Tests for ParkingLot class
func TestNewParkingLot(t *testing.T) {
	t.Run("capacity 0 throws error", func(t *testing.T) {
		_, err := NewParkingLot(0)

		assert.ErrorIs(t, err, CustomErrors.ErrCannotCreateParkingLot)
	})

	t.Run("valid parking lot with 5 slots", func(t *testing.T) {
		parkingLot, err := NewParkingLot(5)

		assert.NotNil(t, parkingLot)
		assert.NoError(t, err)
	})
}

// Tests for Park() method
func TestPark(t *testing.T) {
	t.Run("park car when parking lot is empty", func(t *testing.T) {
		parkingLot, _ := NewParkingLot(2)
		firstCar := Car.NewCar("TS-1234", Car.RED)

		ticket, err := parkingLot.Park(firstCar)

		assert.NotNil(t, ticket)
		assert.NoError(t, err)
	})

	t.Run("park car when parking lot is not empty", func(t *testing.T) {
		parkingLot, _ := NewParkingLot(2)
		firstCar := Car.NewCar("TS-1234", Car.RED)
		parkingLot.Park(firstCar)
		secondCar := Car.NewCar("TS-1235", Car.RED)

		ticket, err := parkingLot.Park(secondCar)

		assert.NotNil(t, ticket)
		assert.NoError(t, err)
	})

	t.Run("park car when parking lot is full", func(t *testing.T) {
		parkingLot, _ := NewParkingLot(1)
		firstCar := Car.NewCar("TS-1234", Car.RED)
		secondCar := Car.NewCar("TS-1235", Car.BLUE)

		parkingLot.Park(firstCar)

		_, err := parkingLot.Park(secondCar)
		assert.ErrorIs(t, err, CustomErrors.ErrParkingLotFull)
	})

	t.Run("park car when car is already parked", func(t *testing.T) {
		parkingLot, _ := NewParkingLot(2)
		firstCar := Car.NewCar("TS-1234", Car.RED)
		parkingLot.Park(firstCar)

		_, err := parkingLot.Park(firstCar)

		assert.ErrorIs(t, err, CustomErrors.ErrCarAlreadyParked)
	})
}

// Tests for Unpark() method
func TestUnpark(t *testing.T) {
	t.Run("unpark car with valid ticket", func(t *testing.T) {
		parkingLot, _ := NewParkingLot(2)
		firstCar := Car.NewCar("TS-1234", Car.RED)
		ticket, _ := parkingLot.Park(firstCar)

		_, err := parkingLot.Unpark(ticket)

		assert.NoError(t, err)
	})

	t.Run("unpark returns the parked car", func(t *testing.T) {
		parkingLot, _ := NewParkingLot(2)
		firstCar := Car.NewCar("TS-1234", Car.RED)
		ticket, _ := parkingLot.Park(firstCar)

		returnedCar, _ := parkingLot.Unpark(ticket)

		assert.Equal(t, firstCar, returnedCar)
	})

	t.Run("slot is empty after unpark", func(t *testing.T) {
		parkingLot, _ := NewParkingLot(1)
		firstCar := Car.NewCar("TS-1234", Car.RED)
		ticket, _ := parkingLot.Park(firstCar)

		_, _ = parkingLot.Unpark(ticket)

		assert.False(t, parkingLot.IsFull())
	})

	t.Run("unpark with invalid ticket throws error", func(t *testing.T) {
		parkingLot, _ := NewParkingLot(2)
		firstCar := Car.NewCar("TS-1234", Car.RED)
		parkingLot.Park(firstCar)
		invalidTicket := &Ticket.Ticket{}

		_, err := parkingLot.Unpark(invalidTicket)

		assert.ErrorIs(t, err, CustomErrors.ErrInvalidTicket)
	})
}

// Tests for CountCarsByColor() method
func TestCountCarsByColor(t *testing.T) {
	t.Run("count cars by color when no car is parked", func(t *testing.T) {
		parkingLot, _ := NewParkingLot(5)

		count := parkingLot.CountCarsByColor(Car.RED)

		assert.Equal(t, 0, count)
	})

	t.Run("count cars by color when cars are parked", func(t *testing.T) {
		parkingLot, _ := NewParkingLot(5)
		firstCar := Car.NewCar("TS-1234", Car.RED)
		secondCar := Car.NewCar("TS-1235", Car.BLUE)
		thirdCar := Car.NewCar("TS-1236", Car.RED)
		parkingLot.Park(firstCar)
		parkingLot.Park(secondCar)
		parkingLot.Park(thirdCar)

		count := parkingLot.CountCarsByColor(Car.RED)

		assert.Equal(t, 2, count)
	})
}

// Tests for GetSlotNumberByRegistrationNumber() method
func TestGetCarSlotNumberByRegistrationNumber(t *testing.T) {
	t.Run("get car slot number when parking lot is empty throws error", func(t *testing.T) {
		parkingLot, _ := NewParkingLot(5)

		_, err := parkingLot.GetCarSlotNumberByRegistrationNumber("TS-1234")

		assert.ErrorIs(t, err, CustomErrors.ErrCarNotFound)
	})

	t.Run("get car slot number when car is parked", func(t *testing.T) {
		parkingLot, _ := NewParkingLot(5)
		firstCar := Car.NewCar("TS-1234", Car.RED)
		parkingLot.Park(firstCar)

		slotNumber, _ := parkingLot.GetCarSlotNumberByRegistrationNumber("TS-1234")

		assert.Equal(t, 1, slotNumber)
	})

	t.Run("get car slot number when car is parked second", func(t *testing.T) {
		parkingLot, _ := NewParkingLot(5)
		firstCar := Car.NewCar("TS-1234", Car.RED)
		secondCar := Car.NewCar("TS-1235", Car.BLUE)
		parkingLot.Park(firstCar)
		parkingLot.Park(secondCar)

		slotNumber, _ := parkingLot.GetCarSlotNumberByRegistrationNumber("TS-1235")

		assert.Equal(t, 2, slotNumber)
	})

	t.Run("get car slot number after car is unparked throws error", func(t *testing.T) {
		parkingLot, _ := NewParkingLot(1)
		firstCar := Car.NewCar("TS-1234", Car.RED)
		ticket, _ := parkingLot.Park(firstCar)
		_, _ = parkingLot.Unpark(ticket)

		_, err := parkingLot.GetCarSlotNumberByRegistrationNumber("TS-1234")

		assert.ErrorIs(t, err, CustomErrors.ErrCarNotFound)
	})

	t.Run("get car slot number after car is not present throws error", func(t *testing.T) {
		parkingLot, _ := NewParkingLot(5)
		firstCar := Car.NewCar("TS-1234", Car.RED)
		secondCar := Car.NewCar("TS-1235", Car.RED)
		parkingLot.Park(firstCar)
		parkingLot.Park(secondCar)

		_, err := parkingLot.GetCarSlotNumberByRegistrationNumber("TS-1236")

		assert.ErrorIs(t, err, CustomErrors.ErrCarNotFound)
	})
}
