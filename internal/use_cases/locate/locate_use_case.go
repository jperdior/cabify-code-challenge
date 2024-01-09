package locate

import "cabify-code-challenge/internal/carpool"

type LocateUseCase struct{}

func NewLocateUseCase() LocateUseCase {
	return LocateUseCase{}
}

func (u LocateUseCase) Locate(carPool *carpool.CarPool, groupId int) (carpool.Car, error) {

	groupIdValueObject, err := carpool.NewGroupID(groupId)
	if err != nil {
		return carpool.Car{}, err
	}

	car, err := carPool.Locate(groupIdValueObject)
	if err != nil {
		return carpool.Car{}, err
	}

	return car, nil
}
