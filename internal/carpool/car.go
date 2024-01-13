package carpool

import (
	"errors"
	"fmt"
)

type CarID struct {
	value int
}

func (c CarID) Value() int {
	return c.value
}

var ErrInvalidCarID = errors.New("invalid car id")

func NewCarID(value int) (CarID, error) {
	if value < 0 {
		return CarID{}, fmt.Errorf("%w: %d", ErrInvalidCarID, value)
	}
	return CarID{value: value}, nil
}

const MinSeats = 4
const MaxSeats = 6

type Seats struct {
	value int
}

func (s Seats) Value() int {
	return s.value
}

var ErrInvalidSeats = errors.New("invalid seats")

func NewSeats(value int) (Seats, error) {
	if value < MinSeats || value > MaxSeats {
		return Seats{}, fmt.Errorf("%w: %d", ErrInvalidSeats, value)
	}
	return Seats{value: value}, nil
}

type AvailableSeats struct {
	value int
}

func (s AvailableSeats) Value() int {
	return s.value
}

var ErrInvalidAvailableSeats = errors.New("invalid available seats")

func NewAvailableSeats(value int) (AvailableSeats, error) {
	if value < 0 {
		return AvailableSeats{}, fmt.Errorf("%w: %d", ErrInvalidAvailableSeats, value)
	}
	return AvailableSeats{value: value}, nil
}

type Car struct {
	id             CarID
	seats          Seats
	availableSeats AvailableSeats
}

// NewCar creates a new car
func NewCar(id int, seats int, availableSeats int) (Car, error) {

	idValueObject, err := NewCarID(id)
	if err != nil {
		return Car{}, err
	}
	seatsValueObject, err := NewSeats(seats)
	if err != nil {
		return Car{}, err
	}
	if availableSeats < 0 {
		availableSeats = seats
	}
	availableSeatsValueObject, err := NewAvailableSeats(availableSeats)
	if err != nil {
		return Car{}, err
	}

	return Car{
		id:             idValueObject,
		seats:          seatsValueObject,
		availableSeats: availableSeatsValueObject,
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

func (c Car) AvailableSeats() AvailableSeats {
	return c.availableSeats
}

var ErrNotEnoughAvailableSeats = errors.New("not enough available seats")

func (c *Car) SitPeople(people int) error {
	if people > c.availableSeats.value {
		return ErrNotEnoughAvailableSeats
	}
	c.availableSeats.value -= people
	return nil
}

var ErrInvalidDropPeople = errors.New("invalid drop people")

func (c *Car) DropPeople(people int) error {
	usedSeats := c.seats.value - c.availableSeats.value
	if people < 0 || people > usedSeats {
		return ErrInvalidDropPeople
	}
	c.availableSeats.value += people
	return nil
}
