package carpool

import (
	"errors"
	"fmt"
)

type CarID struct {
	value int
}

var ErrInvalidCarID = errors.New("invalid car id")

func NewCarID(value int) (CarID, error) {
	if value < 0 {
		return CarID{}, fmt.Errorf("%w: %d", ErrInvalidCarID, value)
	}
	return CarID{value: value}, nil
}

type Seats struct {
	value int
}

var ErrInvalidSeats = errors.New("invalid seats")

func NewSeats(value int) (Seats, error) {
	if value < 0 {
		return Seats{}, fmt.Errorf("%w: %d", ErrInvalidSeats, value)
	}
	return Seats{value: value}, nil
}

type Car struct {
	id    CarID
	seats Seats
}

// NewCar creates a new car
func NewCar(id int, seats int) (Car, error) {

	idValueObject, err := NewCarID(id)
	if err != nil {
		return Car{}, err
	}
	seatsValueObject, err := NewSeats(seats)
	if err != nil {
		return Car{}, err
	}

	return Car{
		id:    idValueObject,
		seats: seatsValueObject,
	}, nil
}

// ID returns the car id
func (c Car) ID() CarID {
	return c.id
}

// Seats returns the car seats
func (c Car) Seats() Seats {
	return c.seats
}
