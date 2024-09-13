package locate

import (
	"cabify-code-challenge/internal/carpool/domain"
)

type LocateService struct {
	carPool *domain.CarPool
}

func NewLocateService(carPool *domain.CarPool) LocateService {
	return LocateService{carPool: carPool}
}

func (s *LocateService) Execute(groupId int) (domain.Car, error) {

	groupIdValueObject, err := domain.NewGroupID(groupId)
	if err != nil {
		return domain.Car{}, err
	}

	car, err := s.carPool.Locate(groupIdValueObject)
	if err != nil {
		return domain.Car{}, err
	}

	return car, nil
}
