package retry_journey

import (
	"cabify-code-challenge/internal/carpool"
	"cabify-code-challenge/kit/event"
	"context"
	"errors"
)

type RetryJourneysUseCase struct{}

func NewRetryJourneyUseCase() RetryJourneysUseCase {
	return RetryJourneysUseCase{}
}

func (s RetryJourneysUseCase) RetryJourneys(context context.Context, evt event.Event) error {
	carPool := context.Value("carPool").(*carpool.CarPool)

	switch evt := evt.(type) {
	case carpool.JourneyDroppedEvent:
		return handleEvent(carPool, evt.AvailableSeats())
	case carpool.CarPutEvent:
		return handleEvent(carPool, evt.AvailableSeats())
	default:
		return errors.New("unexpected event")
	}
}

func handleEvent(carPool *carpool.CarPool, availableSeats int) error {
	for _, group := range carPool.GetWaitingGroups() {
		if availableSeats > group.People().Value() {
			err := carPool.Journey(group)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
