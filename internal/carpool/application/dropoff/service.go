package dropoff

import (
	"cabify-code-challenge/internal/carpool/domain"
	"cabify-code-challenge/kit/event"
)

type DropOffService struct {
	eventBus event.Bus
	carPool  *domain.CarPool
}

func NewDropOffService(eventBus event.Bus, carPool *domain.CarPool) DropOffService {
	return DropOffService{
		eventBus: eventBus,
		carPool:  carPool,
	}
}

func (s *DropOffService) Execute(groupId int) error {

	groupIdValueObject, err := domain.NewGroupID(groupId)
	if err != nil {
		return err
	}

	err = s.carPool.DropOff(groupIdValueObject)
	if err != nil {
		return err
	}

	err = s.eventBus.Publish(s.carPool.PullEvents())
	if err != nil {
		return err
	}
	return nil
}
