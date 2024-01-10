package carpool

import (
	"errors"
	"sync"
)

const MaxSeats = 6

type CarPool struct {
	//contains all the cars in the carpool
	cars map[int]Car
	//contains all the cars in the carpool grouped by number of seats
	carsBySeat map[int][]Car
	//queue of groups waiting for a car
	waitingGroups []Group
	//hash map to find the index of a group in the queue
	waitingGroupsIndexHash map[GroupID]int
	//contains all the groups in the carpool
	groups map[GroupID]Group
	//contains all the journeys in the carpool indexed by group id
	journeys map[GroupID]Journey
	mu       sync.Mutex
}

// NewCarPool creates a new carpool
func NewCarPool() *CarPool {
	return &CarPool{
		cars:                   make(map[int]Car),
		carsBySeat:             make(map[int][]Car),
		groups:                 make(map[GroupID]Group),
		journeys:               make(map[GroupID]Journey),
		waitingGroupsIndexHash: make(map[GroupID]int),
	}
}

// GetGroups
func (carpool *CarPool) GetGroups() map[GroupID]Group {
	return carpool.groups
}

// GetWaitingGroups returns the waiting groups
func (carpool *CarPool) GetWaitingGroups() []Group {
	return carpool.waitingGroups
}

// GetWaitingGroupsIndexHash returns the waiting groups index hash
func (carpool *CarPool) GetWaitingGroupsIndexHash() map[GroupID]int {
	return carpool.waitingGroupsIndexHash
}

// GetJourneys returns the journeys
func (carpool *CarPool) GetJourneys() map[GroupID]Journey {
	return carpool.journeys
}

// GetCarsBySeat returns the put_cars
func (carpool *CarPool) GetCarsBySeat() map[int][]Car {
	return carpool.carsBySeat
}

// SetCars sets the carsBySeat and resets journeys
func (carpool *CarPool) SetCars(cars []Car) {
	carpool.mu.Lock()
	defer carpool.mu.Unlock()

	carpool.carsBySeat = make(map[int][]Car)
	carpool.journeys = make(map[GroupID]Journey)

	for _, car := range cars {
		carpool.carsBySeat[car.Seats().Value()] = append(carpool.carsBySeat[car.Seats().Value()], car)
		carpool.cars[car.ID().Value()] = car
	}
}

// GetCarBySeats returns a car with the given number of seats
func (carpool *CarPool) GetCarBySeats(seats int) (Car, bool) {
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
func (carpool *CarPool) RegisterJourney(journey Journey) {
	carpool.mu.Lock()
	defer carpool.mu.Unlock()

	carpool.journeys[journey.groupID] = journey
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
	if exists {
		delete(carpool.journeys, groupID)
	}

	waitingGroupIndex, exists := carpool.waitingGroupsIndexHash[groupID]

	if exists {
		carpool.waitingGroups = append(carpool.waitingGroups[:waitingGroupIndex], carpool.waitingGroups[waitingGroupIndex+1:]...)
		delete(carpool.waitingGroupsIndexHash, groupID)
	}
	return nil
}

// Locate returns the car where a group is located
func (carpool *CarPool) Locate(groupID GroupID) (Car, error) {

	_, exists := carpool.groups[groupID]
	if !exists {
		return Car{}, ErrGroupNotFound
	}

	journey, exists := carpool.journeys[groupID]
	if !exists {
		return Car{}, nil
	}
	return carpool.cars[journey.CarID().Value()], nil
}
