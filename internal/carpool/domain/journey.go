package domain

type Journey struct {
	group Group
	car   Car
}

// NewJourney creates a new journey
func NewJourney(group Group, car Car) (Journey, error) {
	return Journey{
		group: group,
		car:   car,
	}, nil
}

// Group returns the journey group
func (j Journey) Group() Group {
	return j.group
}

// Car returns the journey car
func (j Journey) Car() Car {
	return j.car
}
