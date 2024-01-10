package carpool

import (
	"errors"
	"sync"
)

type CarPool struct {
	//contains all the cars in the carpool
	cars map[int]Car
	//contains all the cars in the carpool grouped by number of seats
	carsByAvailableSeats map[AvailableSeats]map[CarID]Car
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
		carsByAvailableSeats:   make(map[AvailableSeats]map[CarID]Car),
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

// GetCarsByAvailableSeats returns the put_cars
func (carpool *CarPool) GetCarsByAvailableSeats() map[AvailableSeats]map[CarID]Car {
	return carpool.carsByAvailableSeats
}

// SetCars sets the carsByAvailableSeats and resets journeys
func (carpool *CarPool) SetCars(cars []Car) {
	carpool.mu.Lock()
	defer carpool.mu.Unlock()
	carpool.carsByAvailableSeats = make(map[AvailableSeats]map[CarID]Car)
	carpool.journeys = make(map[GroupID]Journey)

	for _, car := range cars {
		carpool.relocateCarInCarsByAvailableSeatsMap(car)
		carpool.cars[car.ID().Value()] = car
	}
}

func (carpool *CarPool) relocateCarInCarsByAvailableSeatsMap(car Car) {
	_, exists := carpool.carsByAvailableSeats[car.AvailableSeats()]
	if !exists {
		carpool.carsByAvailableSeats[car.AvailableSeats()] = make(map[CarID]Car)
	}
	carpool.carsByAvailableSeats[car.AvailableSeats()][car.ID()] = car
}

var ErrGroupAlreadyExists = errors.New("group already exists")

// AddGroup adds a new group
func (carpool *CarPool) AddGroup(group Group) error {
	carpool.mu.Lock()
	defer carpool.mu.Unlock()

	_, exists := carpool.groups[group.ID()]
	if exists {
		return ErrGroupAlreadyExists
	}
	carpool.groups[group.ID()] = group
	return nil
}

func (carpool *CarPool) Journey(group Group) error {

	for seats := group.People().Value(); seats <= MaxSeats; seats++ {
		seatsValueObject, err := NewAvailableSeats(seats)
		if err != nil {
			return err
		}
		car, exists := carpool.getFirstCarByAvailableSeats(seatsValueObject)
		if exists {
			err := carpool.registerJourney(group, car)
			if err != nil {
				return err
			}
			return nil
		}
	}
	carpool.addWaitingGroup(group)
	return nil
}

// GetFirstCarByAvailableSeats returns a car with the given number of seats, modifies it in the hash of cars and updates the carsAvailableSeats hash
func (carpool *CarPool) getFirstCarByAvailableSeats(availableSeats AvailableSeats) (Car, bool) {
	cars, exists := carpool.carsByAvailableSeats[availableSeats]
	if !exists || len(cars) == 0 {
		return Car{}, false
	}
	for _, car := range cars {
		delete(carpool.carsByAvailableSeats[availableSeats], car.ID())
		return car, true
	}
	return Car{}, false
}

// AddWaitingGroup adds a new group to the queue of waiting groups
func (carpool *CarPool) addWaitingGroup(group Group) {
	carpool.mu.Lock()
	defer carpool.mu.Unlock()

	carpool.waitingGroups = append(carpool.waitingGroups, group)
	carpool.waitingGroupsIndexHash[group.ID()] = len(carpool.waitingGroups) - 1
}

// RegisterJourney adds a new journey
func (carpool *CarPool) registerJourney(group Group, car Car) error {
	carpool.mu.Lock()
	defer carpool.mu.Unlock()

	err := car.SitPeople(group.People().Value())
	if err != nil {
		return err
	}
	journey, err := NewJourney(group, car)
	if err != nil {
		return err
	}
	carpool.journeys[group.ID()] = journey
	carpool.relocateCarInCarsByAvailableSeatsMap(car)

	return nil
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

	//if the group is on journey
	journey, exists := carpool.journeys[groupID]
	if exists {
		err := carpool.deregisterJourney(journey)
		if err != nil {
			return err
		}
		delete(carpool.journeys, groupID)
	}

	waitingGroupIndex, exists := carpool.waitingGroupsIndexHash[groupID]
	if exists {
		carpool.waitingGroups = append(carpool.waitingGroups[:waitingGroupIndex], carpool.waitingGroups[waitingGroupIndex+1:]...)
		delete(carpool.waitingGroupsIndexHash, groupID)
	}

	delete(carpool.groups, groupID)
	return nil
}

func (carpool *CarPool) deregisterJourney(journey Journey) error {
	car := journey.Car()
	delete(carpool.carsByAvailableSeats[car.AvailableSeats()], car.ID())
	err := car.DropPeople(journey.Group().People().Value())
	if err != nil {
		return err
	}
	carpool.relocateCarInCarsByAvailableSeatsMap(car)
	return nil
}

func (carpool *CarPool) recalculateCarBySeat(carID CarID, seats int) {
	carpool.mu.Lock()
	defer carpool.mu.Unlock()

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
	return journey.Car(), nil
}
