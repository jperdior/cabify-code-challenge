package carpool

import (
	"errors"
	"sync"
)

type CarPool struct {
	cars                   map[CarID]Car
	carsBySeat             map[Seats][]Car
	waitingGroups          []Group
	waitingGroupsIndexHash map[GroupID]int
	groups                 map[GroupID]Group
	journeys               map[GroupID]Journey
	mu                     sync.Mutex
}

// NewCarPoolingService creates a new carpool
func NewCarPoolingService() *CarPool {
	return &CarPool{
		cars:       make(map[CarID]Car),
		carsBySeat: make(map[Seats][]Car),
		groups:     make(map[GroupID]Group),
		journeys:   make(map[GroupID]Journey),
	}
}

// SetCars sets the carsBySeat and resets journeys
func (carpool *CarPool) SetCars(cars []Car) {
	carpool.mu.Lock()
	defer carpool.mu.Unlock()

	carpool.carsBySeat = make(map[Seats][]Car)
	carpool.journeys = make(map[GroupID]Journey)

	for _, car := range cars {
		carpool.carsBySeat[car.Seats()] = append(carpool.carsBySeat[car.Seats()], car)
	}
}

// GetCarBySeats returns a car with the given number of seats
func (carpool *CarPool) GetCarBySeats(seats Seats) (Car, bool) {
	carpool.mu.Lock()
	defer carpool.mu.Unlock()

	cars, exists := carpool.carsBySeat[seats]
	if !exists || len(cars) == 0 {
		return Car{}, false
	}
	car := cars[0]

	carpool.carsBySeat[seats] = cars[1:]

	if len(carpool.carsBySeat[seats]) == 0 {
		delete(carpool.carsBySeat, seats)
	}

	return car, true
}

// AddGroup adds a new group
func (carpool *CarPool) AddGroup(group Group) {
	carpool.mu.Lock()
	defer carpool.mu.Unlock()

	carpool.groups[group.ID()] = group
}

// GetGroup returns a group with the given id
func (carpool *CarPool) GetGroup(groupID GroupID) (Group, bool) {
	carpool.mu.Lock()
	defer carpool.mu.Unlock()

	group, exists := carpool.groups[groupID]
	return group, exists
}

// AddWaitingGroup adds a new group to the queue of waiting groups
func (carpool *CarPool) AddWaitingGroup(group Group) {
	carpool.mu.Lock()
	defer carpool.mu.Unlock()

	carpool.waitingGroups = append(carpool.waitingGroups, group)
	carpool.waitingGroupsIndexHash[group.ID()] = len(carpool.waitingGroups) - 1
}

// RegisterJourney adds a new journey
func (carpool *CarPool) RegisterJourney(groupID GroupID, journey Journey) {
	carpool.mu.Lock()
	defer carpool.mu.Unlock()

	carpool.journeys[groupID] = journey
}

var ErrGroupNotFound = errors.New("group not found")

// DropOff drops off a group either from the queue or from a journey
func (carpool *CarPool) DropOff(groupID GroupID) error {
	carpool.mu.Lock()
	defer carpool.mu.Unlock()
	_, exists := carpool.groups[groupID]
	if !exists {
		return ErrGroupNotFound
	}
	delete(carpool.groups, groupID)
	_, exists = carpool.journeys[groupID]
	if !exists {
		return ErrGroupNotFound
	}
	delete(carpool.journeys, groupID)
	waitingGroupIndex, exists := carpool.waitingGroupsIndexHash[groupID]
	if !exists {
		return ErrGroupNotFound
	}
	carpool.waitingGroups = append(carpool.waitingGroups[:waitingGroupIndex], carpool.waitingGroups[waitingGroupIndex+1:]...)
	delete(carpool.waitingGroupsIndexHash, groupID)

	return nil
}

// Locate returns the car where a group is located
func (carpool *CarPool) Locate(groupID GroupID) (Car, error) {
	carpool.mu.Lock()
	defer carpool.mu.Unlock()

	_, exists := carpool.groups[groupID]
	if !exists {
		return Car{}, ErrGroupNotFound
	}

	journey, exists := carpool.journeys[groupID]
	if !exists {
		return Car{}, nil
	}
	return carpool.cars[journey.Car()], nil
}
