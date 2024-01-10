package carpool

import "cabify-code-challenge/kit/event"

const JourneyDroppedEventType event.Type = "journey_dropped"

type JourneyDroppedEvent struct {
	event.BaseEvent
	carId int
}

func NewJourneyDroppedEvent(carId int) JourneyDroppedEvent {
	return JourneyDroppedEvent{
		BaseEvent: event.NewBaseEvent(carId),
		carId:     carId,
	}
}

func (e JourneyDroppedEvent) Type() event.Type {
	return JourneyDroppedEventType
}

func (e JourneyDroppedEvent) CarId() int {
	return e.carId
}
