package carpool

import (
	"fmt"
)

type Journey struct {
	groupID GroupID
	carID   CarID
}

// NewJourney creates a new journey
func NewJourney(groupID int, carID int) (Journey, error) {
	groupIDValueObject, err := NewGroupID(groupID)
	if err != nil {
		return Journey{}, fmt.Errorf("%w: %d", ErrInvalidGroupID, groupID)
	}
	carIDValueObject, err := NewCarID(carID)
	if err != nil {
		return Journey{}, fmt.Errorf("%w: %d", ErrInvalidCarID, carID)
	}
	return Journey{
		groupID: groupIDValueObject,
		carID:   carIDValueObject,
	}, nil
}

// GroupID returns the journey group id
func (j Journey) GroupID() GroupID {
	return j.groupID
}

// CarID returns the journey car id
func (j Journey) CarID() CarID {
	return j.carID
}
