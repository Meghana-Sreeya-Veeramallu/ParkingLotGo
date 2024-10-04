package Car

type Car struct {
	registrationNumber string
	color              CarColor
}

func NewCar(registrationNumber string, color CarColor) *Car {
	return &Car{
		registrationNumber: registrationNumber,
		color:              color,
	}
}

func (car *Car) IsSameColor(color CarColor) bool {
	return car.color == color
}

func (car *Car) IsSameCar(otherCar *Car) bool {
	return car.HasSameRegistrationNumber(otherCar.registrationNumber)
}

func (car *Car) HasSameRegistrationNumber(registrationNumber string) bool {
	return car.registrationNumber == registrationNumber
}
