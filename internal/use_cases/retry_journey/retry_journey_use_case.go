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
		for _, group := range carPool.GetWaitingGroups() {
			if evt.AvailableSeats() > group.People().Value() {
				err := carPool.Journey(group)
				if err != nil {
					return err
				}
			}
		}
	case carpool.CarPutEvent:
		for _, group := range carPool.GetWaitingGroups() {
			if evt.AvailableSeats() > group.People().Value() {
				err := carPool.Journey(group)
				if err != nil {
					return err
				}
			}
		}
	default:
		return errors.New("unexpected event")
	}

	return nil
}
