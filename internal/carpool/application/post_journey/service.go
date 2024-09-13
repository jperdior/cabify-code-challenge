package post_journey

import (
	"cabify-code-challenge/internal/carpool/domain"
)

type JourneyService struct {
	carPool *domain.CarPool
}

func NewJourneyService(carPool *domain.CarPool) JourneyService {
	return JourneyService{
		carPool: carPool,
	}
}

func (s *JourneyService) Execute(groupID int, people int) error {

	group, err := domain.NewGroup(groupID, people)
	if err != nil {
		return err
	}

	err = s.carPool.AddGroup(group)
	if err != nil {
		return err
	}

	err = s.carPool.Journey(group)
	if err != nil {
		return err
	}
	return nil
}
