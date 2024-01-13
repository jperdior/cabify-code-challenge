package carpool

type Journey struct {
	group Group
	car   Car
}

// NewJourney creates a new post_journey
func NewJourney(group Group, car Car) (Journey, error) {
	return Journey{
		group: group,
		car:   car,
	}, nil
}

// Group returns the post_journey group
func (j Journey) Group() Group {
	return j.group
}

// Car returns the post_journey car
func (j Journey) Car() Car {
	return j.car
}
