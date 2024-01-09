package creating_journey

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
	carPool.AddGroup(group)
	//look for a car available for the group
	for seats := group.People().Value(); seats <= carpool.MaxSeats; seats++ {
		car, exists := carPool.GetCarBySeats(seats)
		if exists {
			journey, _ := carpool.NewJourney(group.ID().Value(), car.ID().Value())
			carPool.RegisterJourney(journey)
			return nil
		}
	}
	//keep the group in queue
	carPool.AddWaitingGroup(group)

	return nil
}
