package Attendant

import (
	CustomErrors "ParkingLotGo"
	"ParkingLotGo/Car"
	"ParkingLotGo/ParkingLot"
	"github.com/stretchr/testify/assert"
	"testing"
)

// tests for Assign() method
func TestAssignParkingLot(t *testing.T) {
	t.Run("parking lot to an attendant", func(t *testing.T) {
		attendant := NewAttendant(&BasicNextLotStrategy{})
		parkingLot, _ := ParkingLot.NewParkingLot(5)

		assert.NoError(t, attendant.Assign(parkingLot))
	})

	t.Run("two parking lots to an attendant", func(t *testing.T) {
		attendant := NewAttendant(&BasicNextLotStrategy{})
		firstParkingLot, _ := ParkingLot.NewParkingLot(5)
		secondParkingLot, _ := ParkingLot.NewParkingLot(5)

		assert.NoError(t, attendant.Assign(firstParkingLot))
		assert.NoError(t, attendant.Assign(secondParkingLot))
	})

	t.Run("can not assign the same parking lot twice", func(t *testing.T) {
		attendant := NewAttendant(&BasicNextLotStrategy{})
		parkingLot, _ := ParkingLot.NewParkingLot(5)

		assert.NoError(t, attendant.Assign(parkingLot))
		assert.Error(t, attendant.Assign(parkingLot), CustomErrors.ErrParkingLotAlreadyAssigned)
	})

	t.Run("assign parking lot to smart attendant", func(t *testing.T) {
		smartAttendant := NewAttendant(&SmartNextLotStrategy{})
		parkingLot, _ := ParkingLot.NewParkingLot(5)

		assert.NoError(t, smartAttendant.Assign(parkingLot))
	})

	// Smart Attendant
	t.Run("assign two parking lots to smart attendant", func(t *testing.T) {
		smartAttendant := NewAttendant(&SmartNextLotStrategy{})
		firstParkingLot, _ := ParkingLot.NewParkingLot(5)
		secondParkingLot, _ := ParkingLot.NewParkingLot(5)

		assert.NoError(t, smartAttendant.Assign(firstParkingLot))
		assert.NoError(t, smartAttendant.Assign(secondParkingLot))
	})

	t.Run("can not assign the same parking lot twice to smart attendant", func(t *testing.T) {
		smartAttendant := NewAttendant(&SmartNextLotStrategy{})
		parkingLot, _ := ParkingLot.NewParkingLot(5)

		assert.NoError(t, smartAttendant.Assign(parkingLot))
		assert.Error(t, smartAttendant.Assign(parkingLot), CustomErrors.ErrParkingLotAlreadyAssigned)
	})
}

// Tests for Park() method
func TestPark(t *testing.T) {
	t.Run("if no parking lot is assigned", func(t *testing.T) {
		attendant := NewAttendant(&BasicNextLotStrategy{})
		car := Car.NewCar("TS-1234", Car.RED)

		_, err := attendant.Park(car)

		assert.Error(t, err, CustomErrors.ErrNoParkingLotAssigned)
	})

	t.Run("if parking lot is not full", func(t *testing.T) {
		attendant := NewAttendant(&BasicNextLotStrategy{})
		parkingLot, _ := ParkingLot.NewParkingLot(5)
		attendant.Assign(parkingLot)
		car := Car.NewCar("TS-1234", Car.RED)
		_, err := attendant.Park(car)
		assert.NoError(t, err)
	})

	t.Run("should return ticket", func(t *testing.T) {
		attendant := NewAttendant(&BasicNextLotStrategy{})
		parkingLot, _ := ParkingLot.NewParkingLot(5)
		attendant.Assign(parkingLot)
		car := Car.NewCar("TS-1234", Car.RED)
		ticket, err := attendant.Park(car)
		assert.NoError(t, err)
		assert.NotNil(t, ticket)
	})

	t.Run("if first parking lot is full", func(t *testing.T) {
		attendant := NewAttendant(&BasicNextLotStrategy{})
		firstParkingLot, _ := ParkingLot.NewParkingLot(1)
		secondParkingLot, _ := ParkingLot.NewParkingLot(1)
		attendant.Assign(firstParkingLot)
		attendant.Assign(secondParkingLot)
		firstCar := Car.NewCar("TS-1234", Car.RED)
		secondCar := Car.NewCar("TS-1235", Car.BLUE)

		_, err := attendant.Park(firstCar)
		assert.NoError(t, err)
		_, err = attendant.Park(secondCar)
		assert.NoError(t, err)
	})

	t.Run("if all parking lots are full", func(t *testing.T) {
		attendant := NewAttendant(&BasicNextLotStrategy{})
		firstParkingLot, _ := ParkingLot.NewParkingLot(1)
		secondParkingLot, _ := ParkingLot.NewParkingLot(1)
		attendant.Assign(firstParkingLot)
		attendant.Assign(secondParkingLot)
		firstCar := Car.NewCar("TS-1234", Car.RED)
		secondCar := Car.NewCar("TS-1235", Car.BLUE)
		thirdCar := Car.NewCar("TS-1236", Car.RED)
		attendant.Park(firstCar)
		attendant.Park(secondCar)

		_, err := attendant.Park(thirdCar)
		assert.Error(t, err, CustomErrors.ErrParkingLotFull)
	})

	t.Run("if car is already parked", func(t *testing.T) {
		attendant := NewAttendant(&BasicNextLotStrategy{})
		parkingLot, _ := ParkingLot.NewParkingLot(5)
		attendant.Assign(parkingLot)
		car := Car.NewCar("TS-1234", Car.RED)
		attendant.Park(car)

		_, err := attendant.Park(car)
		assert.Error(t, err, CustomErrors.ErrCarAlreadyParked)
	})

	t.Run("if car is already parked in another parking lot", func(t *testing.T) {
		attendant := NewAttendant(&BasicNextLotStrategy{})
		firstParkingLot, _ := ParkingLot.NewParkingLot(1)
		secondParkingLot, _ := ParkingLot.NewParkingLot(5)
		attendant.Assign(firstParkingLot)
		attendant.Assign(secondParkingLot)
		firstCar := Car.NewCar("TS-1234", Car.RED)
		secondCar := Car.NewCar("TS-1235", Car.BLUE)
		attendant.Park(firstCar)
		attendant.Park(secondCar)

		_, err := attendant.Park(firstCar)
		assert.Error(t, err, CustomErrors.ErrCarAlreadyParked)
	})

	// Smart Attendant
	t.Run("when second parking lot has more available slots", func(t *testing.T) {
		smartAttendant := NewAttendant(&SmartNextLotStrategy{})
		firstParkingLot, _ := ParkingLot.NewParkingLot(1)
		secondParkingLot, _ := ParkingLot.NewParkingLot(2)
		smartAttendant.Assign(firstParkingLot)
		smartAttendant.Assign(secondParkingLot)
		firstCar := Car.NewCar("TS-1234", Car.RED)
		smartAttendant.Park(firstCar)

		assert.False(t, firstParkingLot.IsFull())
	})

	t.Run("when both parking lots have equal capacity", func(t *testing.T) {
		smartAttendant := NewAttendant(&SmartNextLotStrategy{})
		firstParkingLot, _ := ParkingLot.NewParkingLot(1)
		secondParkingLot, _ := ParkingLot.NewParkingLot(1)
		smartAttendant.Assign(firstParkingLot)
		smartAttendant.Assign(secondParkingLot)
		firstCar := Car.NewCar("TS-1234", Car.RED)
		smartAttendant.Park(firstCar)

		assert.True(t, firstParkingLot.IsFull())
		assert.False(t, secondParkingLot.IsFull())
	})

	t.Run("when both parking lots have equal capacity after a car is parked", func(t *testing.T) {
		smartAttendant := NewAttendant(&SmartNextLotStrategy{})
		firstParkingLot, _ := ParkingLot.NewParkingLot(2)
		secondParkingLot, _ := ParkingLot.NewParkingLot(1)
		smartAttendant.Assign(firstParkingLot)
		smartAttendant.Assign(secondParkingLot)
		firstCar := Car.NewCar("TS-1234", Car.RED)
		secondCar := Car.NewCar("TS-1235", Car.BLUE)
		smartAttendant.Park(firstCar)
		smartAttendant.Park(secondCar)

		assert.True(t, firstParkingLot.IsFull())
		assert.False(t, secondParkingLot.IsFull())
	})
}

// Tests for unpark() method
func TestUnpark(t *testing.T) {
	t.Run("if ticket is valid for first car", func(t *testing.T) {
		attendant := NewAttendant(&BasicNextLotStrategy{})
		parkingLot, _ := ParkingLot.NewParkingLot(5)
		attendant.Assign(parkingLot)
		car := Car.NewCar("TS-1234", Car.RED)
		ticket, _ := attendant.Park(car)

		_, err := attendant.Unpark(ticket)

		assert.NoError(t, err)
	})

	t.Run("if ticket is valid for second car", func(t *testing.T) {
		attendant := NewAttendant(&BasicNextLotStrategy{})
		parkingLot, _ := ParkingLot.NewParkingLot(5)
		attendant.Assign(parkingLot)
		firstCar := Car.NewCar("TS-1234", Car.RED)
		secondCar := Car.NewCar("TS-1235", Car.BLUE)
		attendant.Park(firstCar)
		ticket, _ := attendant.Park(secondCar)

		unparkedCar, err := attendant.Unpark(ticket)

		assert.NoError(t, err)
		assert.Equal(t, secondCar, unparkedCar)
	})

	t.Run("if ticket is valid for second parking lot", func(t *testing.T) {
		attendant := NewAttendant(&BasicNextLotStrategy{})
		firstParkingLot, _ := ParkingLot.NewParkingLot(1)
		secondParkingLot, _ := ParkingLot.NewParkingLot(5)
		attendant.Assign(firstParkingLot)
		attendant.Assign(secondParkingLot)
		firstCar := Car.NewCar("TS-1234", Car.RED)
		secondCar := Car.NewCar("TS-1235", Car.BLUE)
		attendant.Park(firstCar)
		ticket, _ := attendant.Park(secondCar)

		unparkedCar, err := attendant.Unpark(ticket)

		assert.NoError(t, err)
		assert.Equal(t, secondCar, unparkedCar)
	})

	t.Run("if ticket is invalid", func(t *testing.T) {
		attendant := NewAttendant(&BasicNextLotStrategy{})
		parkingLot, _ := ParkingLot.NewParkingLot(5)
		attendant.Assign(parkingLot)
		firstCar := Car.NewCar("TS-1234", Car.RED)
		secondCar := Car.NewCar("TS-1235", Car.BLUE)
		attendant.Park(firstCar)
		ticket, _ := attendant.Park(secondCar)
		attendant.Unpark(ticket)

		_, err := attendant.Unpark(ticket)

		assert.Error(t, err, CustomErrors.ErrInvalidTicket)
	})

	t.Run("if ticket is invalid for second parking lot", func(t *testing.T) {
		attendant := NewAttendant(&BasicNextLotStrategy{})
		firstParkingLot, _ := ParkingLot.NewParkingLot(1)
		secondParkingLot, _ := ParkingLot.NewParkingLot(5)
		attendant.Assign(firstParkingLot)
		attendant.Assign(secondParkingLot)
		firstCar := Car.NewCar("TS-1234", Car.RED)
		secondCar := Car.NewCar("TS-1235", Car.BLUE)
		attendant.Park(firstCar)
		ticket, _ := attendant.Park(secondCar)
		attendant.Unpark(ticket)

		_, err := attendant.Unpark(ticket)

		assert.Error(t, err, CustomErrors.ErrInvalidTicket)
	})

	t.Run("unpark car and allow new car to park", func(t *testing.T) {
		attendant := NewAttendant(&BasicNextLotStrategy{})
		firstParkingLot, _ := ParkingLot.NewParkingLot(1)
		secondParkingLot, _ := ParkingLot.NewParkingLot(1)
		attendant.Assign(firstParkingLot)
		attendant.Assign(secondParkingLot)
		firstCar := Car.NewCar("TS-1234", Car.RED)
		secondCar := Car.NewCar("TS-1235", Car.BLUE)
		thirdCar := Car.NewCar("TS-1236", Car.RED)
		ticket, _ := attendant.Park(firstCar)
		attendant.Park(secondCar)
		attendant.Unpark(ticket)

		_, err := attendant.Park(thirdCar)

		assert.NoError(t, err)
	})

	// Smart Attendant
	t.Run("if ticket is valid for first car to smart attendant", func(t *testing.T) {
		smartAttendant := NewAttendant(&SmartNextLotStrategy{})
		parkingLot, _ := ParkingLot.NewParkingLot(5)
		smartAttendant.Assign(parkingLot)
		car := Car.NewCar("TS-1234", Car.RED)
		ticket, _ := smartAttendant.Park(car)

		_, err := smartAttendant.Unpark(ticket)

		assert.NoError(t, err)
	})

	t.Run("if ticket is valid for second car to smart attendant", func(t *testing.T) {
		smartAttendant := NewAttendant(&BasicNextLotStrategy{})
		parkingLot, _ := ParkingLot.NewParkingLot(5)
		smartAttendant.Assign(parkingLot)
		firstCar := Car.NewCar("TS-1234", Car.RED)
		secondCar := Car.NewCar("TS-1235", Car.BLUE)
		smartAttendant.Park(firstCar)
		ticket, _ := smartAttendant.Park(secondCar)

		unparkedCar, err := smartAttendant.Unpark(ticket)

		assert.NoError(t, err)
		assert.Equal(t, secondCar, unparkedCar)
	})

	t.Run("if ticket is valid for second parking lot to smart attendant", func(t *testing.T) {
		smartAttendant := NewAttendant(&BasicNextLotStrategy{})
		firstParkingLot, _ := ParkingLot.NewParkingLot(1)
		secondParkingLot, _ := ParkingLot.NewParkingLot(5)
		smartAttendant.Assign(firstParkingLot)
		smartAttendant.Assign(secondParkingLot)
		firstCar := Car.NewCar("TS-1234", Car.RED)
		secondCar := Car.NewCar("TS-1235", Car.BLUE)
		smartAttendant.Park(firstCar)
		ticket, _ := smartAttendant.Park(secondCar)

		unparkedCar, err := smartAttendant.Unpark(ticket)

		assert.NoError(t, err)
		assert.Equal(t, secondCar, unparkedCar)
	})

	t.Run("if ticket is invalid to smart attendant", func(t *testing.T) {
		smartAttendant := NewAttendant(&BasicNextLotStrategy{})
		parkingLot, _ := ParkingLot.NewParkingLot(5)
		smartAttendant.Assign(parkingLot)
		firstCar := Car.NewCar("TS-1234", Car.RED)
		secondCar := Car.NewCar("TS-1235", Car.BLUE)
		smartAttendant.Park(firstCar)
		ticket, _ := smartAttendant.Park(secondCar)
		smartAttendant.Unpark(ticket)

		_, err := smartAttendant.Unpark(ticket)

		assert.Error(t, err, CustomErrors.ErrInvalidTicket)
	})

	t.Run("if ticket is invalid for second parking lot to smart attendant", func(t *testing.T) {
		smartAttendant := NewAttendant(&BasicNextLotStrategy{})
		firstParkingLot, _ := ParkingLot.NewParkingLot(1)
		secondParkingLot, _ := ParkingLot.NewParkingLot(5)
		smartAttendant.Assign(firstParkingLot)
		smartAttendant.Assign(secondParkingLot)
		firstCar := Car.NewCar("TS-1234", Car.RED)
		secondCar := Car.NewCar("TS-1235", Car.BLUE)
		smartAttendant.Park(firstCar)
		ticket, _ := smartAttendant.Park(secondCar)
		smartAttendant.Unpark(ticket)

		_, err := smartAttendant.Unpark(ticket)

		assert.Error(t, err, CustomErrors.ErrInvalidTicket)
	})

	t.Run("unpark car and allow new car to park to smart attendant", func(t *testing.T) {
		smartAttendant := NewAttendant(&BasicNextLotStrategy{})
		firstParkingLot, _ := ParkingLot.NewParkingLot(1)
		secondParkingLot, _ := ParkingLot.NewParkingLot(1)
		smartAttendant.Assign(firstParkingLot)
		smartAttendant.Assign(secondParkingLot)
		firstCar := Car.NewCar("TS-1234", Car.RED)
		secondCar := Car.NewCar("TS-1235", Car.BLUE)
		thirdCar := Car.NewCar("TS-1236", Car.RED)
		ticket, _ := smartAttendant.Park(firstCar)
		smartAttendant.Park(secondCar)
		smartAttendant.Unpark(ticket)

		_, err := smartAttendant.Park(thirdCar)

		assert.NoError(t, err)
	})
}

// Tests to cover the mix of attendant and smart attendant
func TestAttendantAndSmartAttendantParkConsecutively(t *testing.T) {
	firstCar := Car.NewCar("TS-1231", Car.RED)
	secondCar := Car.NewCar("TS-1232", Car.BLUE)
	thirdCar := Car.NewCar("TS-1233", Car.GREEN)
	fourthCar := Car.NewCar("TS-1234", Car.BLACK)
	fifthCar := Car.NewCar("TS-1235", Car.YELLOW)

	t.Run("attendant is first", func(t *testing.T) {
		firstParkingLot, _ := ParkingLot.NewParkingLot(2)
		secondParkingLot, _ := ParkingLot.NewParkingLot(2)
		thirdParkingLot, _ := ParkingLot.NewParkingLot(2)
		attendant := NewAttendant(&BasicNextLotStrategy{})
		smartAttendant := NewAttendant(&SmartNextLotStrategy{})

		attendant.Assign(firstParkingLot)
		attendant.Assign(secondParkingLot)
		smartAttendant.Assign(secondParkingLot)
		smartAttendant.Assign(thirdParkingLot)

		attendant.Park(firstCar)
		assert.False(t, firstParkingLot.IsFull())

		smartAttendant.Park(secondCar)
		assert.False(t, firstParkingLot.IsFull())

		attendant.Park(thirdCar)
		assert.True(t, firstParkingLot.IsFull())

		smartAttendant.Park(fourthCar)
		assert.False(t, secondParkingLot.IsFull())

		attendant.Park(fifthCar)
		assert.True(t, secondParkingLot.IsFull())
		assert.False(t, thirdParkingLot.IsFull())
	})

	t.Run("smart attendant is first", func(t *testing.T) {
		firstParkingLot, _ := ParkingLot.NewParkingLot(2)
		secondParkingLot, _ := ParkingLot.NewParkingLot(2)
		thirdParkingLot, _ := ParkingLot.NewParkingLot(2)
		attendant := NewAttendant(&BasicNextLotStrategy{})
		smartAttendant := NewAttendant(&SmartNextLotStrategy{})

		attendant.Assign(firstParkingLot)
		attendant.Assign(secondParkingLot)
		smartAttendant.Assign(secondParkingLot)
		smartAttendant.Assign(thirdParkingLot)

		smartAttendant.Park(firstCar)
		attendant.Park(secondCar)

		assert.False(t, firstParkingLot.IsFull())
		assert.False(t, secondParkingLot.IsFull())

		smartAttendant.Park(thirdCar)
		assert.False(t, secondParkingLot.IsFull())

		attendant.Park(fourthCar)
		assert.True(t, firstParkingLot.IsFull())

		smartAttendant.Park(fifthCar)
		assert.False(t, thirdParkingLot.IsFull())
	})
}
