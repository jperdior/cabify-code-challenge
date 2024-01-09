package dropoff

import "cabify-code-challenge/internal/carpool"

type DropOffUseCase struct{}

func NewDropOffUseCase() DropOffUseCase {
	return DropOffUseCase{}
}

func (u DropOffUseCase) DropOff(carPool *carpool.CarPool, groupId int) error {

	groupIdValueObject, err := carpool.NewGroupID(groupId)
	if err != nil {
		return err
	}

	err = carPool.DropOff(groupIdValueObject)
	if err != nil {
		return err
	}

	return nil
}
