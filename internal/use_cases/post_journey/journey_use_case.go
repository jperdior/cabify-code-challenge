package post_journey

import (
	"cabify-code-challenge/internal/carpool"
	"context"
	"errors"
)

type CreateJourneyUseCase struct{}

func NewCreateJourneyUseCase() CreateJourneyUseCase {
	return CreateJourneyUseCase{}
}

func (s CreateJourneyUseCase) CreateJourney(context context.Context, groupID int, people int) error {
	carPool, ok := context.Value("carPool").(*carpool.CarPool)
	if !ok {
		return errors.New("carPool not found in context")
	}

	group, err := carpool.NewGroup(groupID, people)
	if err != nil {
		return err
	}

	err = carPool.AddGroup(group)
	if err != nil {
		return err
	}

	err = carPool.Journey(group)
	if err != nil {
		return err
	}
	return nil
}
