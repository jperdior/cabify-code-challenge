package retry_journey

import (
	"cabify-code-challenge/internal/carpool/domain"
	"cabify-code-challenge/kit/event"
	"errors"
)

type RetryJourneysOnCarPut struct {
	service RetryJourneyService
}

func NewRetryJourneysOnCarPut(service RetryJourneyService) RetryJourneysOnCarPut {
	return RetryJourneysOnCarPut{service: service}
}

func (e RetryJourneysOnCarPut) Handle(evt event.Event) error {
	catPutEvent, ok := evt.(domain.CarPutEvent)
	if !ok {
		return errors.New("unexpected event")
	}
	return e.service.Execute(catPutEvent.AvailableSeats())
}

type RetryJourneysOnJourneyDropped struct {
	service RetryJourneyService
}

func NewRetryJourneysOnJourneyDropped(service RetryJourneyService) RetryJourneysOnJourneyDropped {
	return RetryJourneysOnJourneyDropped{service: service}
}

func (e RetryJourneysOnJourneyDropped) Handle(evt event.Event) error {
	journeyDroppedEvent, ok := evt.(domain.JourneyDroppedEvent)
	if !ok {
		return errors.New("unexpected event")
	}
	return e.service.Execute(journeyDroppedEvent.AvailableSeats())
}
