package put_cars

import (
	"cabify-code-challenge/internal/carpool"
	"cabify-code-challenge/internal/use_cases/retry_journey"
	"cabify-code-challenge/kit/event"
	"context"
	"errors"
)

type RetryJourneysOnCarPut struct {
	retryUseCase retry_journey.RetryJourneysUseCase
}

func NewRetryJourneysOnCarPut(retryUseCase retry_journey.RetryJourneysUseCase) RetryJourneysOnCarPut {
	return RetryJourneysOnCarPut{
		retryUseCase: retryUseCase,
	}
}

func (e RetryJourneysOnCarPut) Handle(context context.Context, evt event.Event) error {
	_, ok := evt.(carpool.CarPutEvent)
	if !ok {
		return errors.New("unexpected event")
	}
	return e.retryUseCase.RetryJourneys(context, evt)
}
