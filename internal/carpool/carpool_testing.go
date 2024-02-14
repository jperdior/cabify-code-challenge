package carpool

func NewTestCarPoolWithCarsAndJourneys(
	cars map[CarID]Car,
	carsByAvailableSeats map[AvailableSeats]map[CarID]Car,
	groups map[GroupID]Group,
	journeys map[GroupID]Journey,
) *CarPool {
	return &CarPool{
		cars:                   cars,
		carsByAvailableSeats:   carsByAvailableSeats,
		waitingGroupsIndexHash: make(map[GroupID]int),
		groups:                 groups,
		journeys:               journeys,
	}
}

func NewTestCarPoolWithoutCars() *CarPool {
	return &CarPool{
		cars:                   make(map[CarID]Car),
		carsByAvailableSeats:   make(map[AvailableSeats]map[CarID]Car),
		groups:                 make(map[GroupID]Group),
		journeys:               make(map[GroupID]Journey),
		waitingGroupsIndexHash: make(map[GroupID]int),
	}
}

func NewTestCarPoolWithCars(
	cars map[CarID]Car,
	carsByAvailableSeats map[AvailableSeats]map[CarID]Car,
) *CarPool {
	return &CarPool{
		cars:                   cars,
		carsByAvailableSeats:   carsByAvailableSeats,
		groups:                 make(map[GroupID]Group),
		journeys:               make(map[GroupID]Journey),
		waitingGroupsIndexHash: make(map[GroupID]int),
	}
}

func NewTestCarPoolWithWaitingGroups(
	groups map[GroupID]Group,
	waitingGroups []Group,
	waitingGroupsIndexHash map[GroupID]int,
) *CarPool {
	return &CarPool{
		cars:                   make(map[CarID]Car),
		carsByAvailableSeats:   make(map[AvailableSeats]map[CarID]Car),
		groups:                 groups,
		journeys:               make(map[GroupID]Journey),
		waitingGroups:          waitingGroups,
		waitingGroupsIndexHash: waitingGroupsIndexHash,
	}
}

func NewTestCarPoolWithCarsAndWaitingGroups(
	cars map[CarID]Car,
	carsByAvailableSeats map[AvailableSeats]map[CarID]Car,
	groups map[GroupID]Group,
	waitingGroups []Group,
	waitingGroupsIndexHash map[GroupID]int,
) *CarPool {
	return &CarPool{
		cars:                   cars,
		carsByAvailableSeats:   carsByAvailableSeats,
		waitingGroups:          waitingGroups,
		waitingGroupsIndexHash: waitingGroupsIndexHash,
		groups:                 groups,
		journeys:               make(map[GroupID]Journey),
	}
}
