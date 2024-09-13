package put_cars

import (
	"cabify-code-challenge/internal/carpool/domain"
	"cabify-code-challenge/kit/event"
)

type PutCarsService struct {
	eventBus event.Bus
	carPool  *domain.CarPool
}

// NewPutCarsService returns the default Service interface implementation
func NewPutCarsService(eventBus event.Bus, carPool *domain.CarPool) PutCarsService {
	return PutCarsService{
		eventBus: eventBus,
		carPool:  carPool,
	}
}

// Execute sets the put_cars in the carPool
func (s *PutCarsService) Execute(cars []domain.Car) error {
	s.carPool.SetCars(cars)
	err := s.eventBus.Publish(s.carPool.PullEvents())
	if err != nil {
		return err
	}
	return nil
}
