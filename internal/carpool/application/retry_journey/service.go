package retry_journey

import (
	"cabify-code-challenge/internal/carpool/domain"
)

type RetryJourneyService struct {
	carPool *domain.CarPool
}

func NewRetryJourneyService(carPool *domain.CarPool) RetryJourneyService {
	return RetryJourneyService{carPool: carPool}
}

func (s *RetryJourneyService) Execute(availableSeats int) error {
	for _, group := range s.carPool.GetWaitingGroups() {
		if availableSeats >= group.People().Value() {
			err := s.carPool.Journey(group)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
