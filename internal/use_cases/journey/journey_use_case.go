package journey

import (
	"cabify-code-challenge/internal/carpool"
)

type CreateJourneyUseCase struct{}

func NewCreateJourneyUseCase() CreateJourneyUseCase {
	return CreateJourneyUseCase{}
}

func (s CreateJourneyUseCase) CreateJourney(carPool *carpool.CarPool, groupID int, people int) error {

	group, err := carpool.NewGroup(groupID, people)
	if err != nil {
		return err
	}
	//register the group in the carpool
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
