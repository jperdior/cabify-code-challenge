package carpool

import "sync"

type CarPool struct {
	cars     []Car
	groups   []Group
	journeys []Journey
	mu       sync.Mutex
}

// NewCarPoolingService creates a new car pool service
func NewCarPoolingService() CarPool {
	return CarPool{}
}

// SetCars sets the cars
func (s *CarPool) SetCars(cars []Car) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.journeys = nil
	s.cars = cars
}

// AddCar adds a new car
func (s *CarPool) AddCar(car Car) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.cars = append(s.cars, car)
}

// AddGroup adds a new group
func (s *CarPool) AddGroup(group Group) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.groups = append(s.groups, group)
}

// AddJourney adds a new journey
func (s *CarPool) AddJourney(journey Journey) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.journeys = append(s.journeys, journey)
}
