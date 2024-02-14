package carpool

import "cabify-code-challenge/kit/event"

const CarPutEventType event.Type = "car_put"

type CarPutEvent struct {
	event.BaseEvent
	carId          int
	seats          int
	availableSeats int
}

func NewCarPutEvent(carId int, seats int, availableSeats int) CarPutEvent {
	return CarPutEvent{
		BaseEvent:      event.NewBaseEvent(carId),
		carId:          carId,
		seats:          seats,
		availableSeats: availableSeats,
	}
}

func (e CarPutEvent) Type() event.Type {
	return CarPutEventType
}

func (e CarPutEvent) CarId() int {
	return e.carId
}

func (e CarPutEvent) Seats() int {
	return e.seats
}

func (e CarPutEvent) AvailableSeats() int {
	return e.availableSeats
}

const JourneyDroppedEventType event.Type = "journey_dropped"

type JourneyDroppedEvent struct {
	event.BaseEvent
	carId          int
	seats          int
	availableSeats int
}

func NewJourneyDroppedEvent(carId int, seats int, availableSeats int) JourneyDroppedEvent {
	return JourneyDroppedEvent{
		BaseEvent:      event.NewBaseEvent(carId),
		carId:          carId,
		seats:          seats,
		availableSeats: availableSeats,
	}
}

func (e JourneyDroppedEvent) Type() event.Type {
	return JourneyDroppedEventType
}

func (e JourneyDroppedEvent) CarId() int {
	return e.carId
}

func (e JourneyDroppedEvent) Seats() int {
	return e.seats
}

func (e JourneyDroppedEvent) AvailableSeats() int {
	return e.availableSeats
}
