package dropoff

import (
	"cabify-code-challenge/internal/carpool"
	"cabify-code-challenge/internal/use_cases/retry_journey"
	"cabify-code-challenge/kit/event"
	"context"
	"errors"
)

type RetryJourneysOnJourneyDropped struct {
	retryUseCase retry_journey.RetryJourneysUseCase
}

func NewRetryJourneysOnJourneyDropped(retryUseCase retry_journey.RetryJourneysUseCase) RetryJourneysOnJourneyDropped {
	return RetryJourneysOnJourneyDropped{
		retryUseCase: retryUseCase,
	}
}

func (e RetryJourneysOnJourneyDropped) Handle(context context.Context, evt event.Event) error {
	_, ok := evt.(carpool.JourneyDroppedEvent)
	if !ok {
		return errors.New("unexpected event")
	}
	return e.retryUseCase.RetryJourneys(context, evt)
}
