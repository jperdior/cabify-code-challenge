package creating_journey

type CreateJourneyUseCase struct{}

func NewCreateJourneyUseCase() CreateJourneyUseCase {
	return CreateJourneyUseCase{}
}

func (s CreateJourneyUseCase) CreateJourney(groupID int, people int) error {
	return nil
}
