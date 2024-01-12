package dropoff

import (
	"cabify-code-challenge/internal/carpool"
	"cabify-code-challenge/kit/event"
	"context"
)

type DropOffUseCase struct {
	eventBus event.Bus
}

func NewDropOffUseCase(eventBus event.Bus) DropOffUseCase {
	return DropOffUseCase{
		eventBus: eventBus,
	}
}

func (u DropOffUseCase) DropOff(context context.Context, groupId int) error {
	carPool := context.Value("carPool").(*carpool.CarPool)

	groupIdValueObject, err := carpool.NewGroupID(groupId)
	if err != nil {
		return err
	}

	err = carPool.DropOff(groupIdValueObject)
	if err != nil {
		return err
	}

	err = u.eventBus.Publish(context, carPool.PullEvents())
	if err != nil {
		return err
	}
	return nil
}
