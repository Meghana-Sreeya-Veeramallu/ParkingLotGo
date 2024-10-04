package Car

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Tests for IsSameColor() method
func TestIsSameColor(t *testing.T) {
	car := NewCar("TS-1234", RED)

	t.Run("IfColorIsSame", func(t *testing.T) {
		assert.True(t, car.IsSameColor(RED))
	})

	t.Run("IfColorIsDifferent", func(t *testing.T) {
		assert.False(t, car.IsSameColor(BLUE))
	})
}

// Tests for IsSameCar() method
func TestIsSameCar(t *testing.T) {
	car := NewCar("TS-1234", RED)

	t.Run("IfCarsAreSame", func(t *testing.T) {
		assert.True(t, car.IsSameCar(car))
	})

	t.Run("IfCarsHaveSameAttributes", func(t *testing.T) {
		otherCar := NewCar("TS-1234", RED)

		assert.True(t, car.IsSameCar(otherCar))
	})

	t.Run("IfCarsAreDifferent", func(t *testing.T) {
		otherCar := NewCar("TS-1235", BLUE)

		assert.False(t, car.IsSameCar(otherCar))
	})
}

// Tests for HasSameRegistrationNumber() method
func TestHasSameRegistrationNumber(t *testing.T) {
	car := NewCar("TS-1234", RED)

	t.Run("IfRegistrationNumbersAreSame", func(t *testing.T) {
		assert.True(t, car.HasSameRegistrationNumber("TS-1234"))
	})

	t.Run("IfRegistrationNumbersAreDifferent", func(t *testing.T) {
		assert.False(t, car.HasSameRegistrationNumber("TS-1235"))
	})
}
