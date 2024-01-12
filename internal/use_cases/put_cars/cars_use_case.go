package put_cars

import (
	"cabify-code-challenge/internal/carpool"
	"cabify-code-challenge/kit/event"
	"context"
	"errors"
)

type PuttingCarsUseCase struct {
	eventBus event.Bus
}

// NewPuttingCarsUseCase returns the default Service interface implementation
func NewPuttingCarsUseCase(eventBus event.Bus) PuttingCarsUseCase {
	return PuttingCarsUseCase{
		eventBus: eventBus,
	}
}

// PutCars sets the put_cars in the carPool
func (s PuttingCarsUseCase) PutCars(context context.Context, cars []carpool.Car) error {
	carPool, ok := context.Value("carPool").(*carpool.CarPool)
	if !ok {
		return errors.New("carPool not found in context")
	}

	carPool.SetCars(cars)
	err := s.eventBus.Publish(context, carPool.PullEvents())
	if err != nil {
		return err
	}
	return nil
}
