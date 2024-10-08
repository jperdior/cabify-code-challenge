package domain

import (
	"errors"
	"fmt"
)

type GroupID struct {
	value int
}

func (g GroupID) Value() int {
	return g.value
}

var ErrInvalidGroupID = errors.New("invalid group id")

// NewGroupID creates a new group id
func NewGroupID(value int) (GroupID, error) {
	if value < 0 {
		return GroupID{}, fmt.Errorf("%w: %d", ErrInvalidGroupID, value)
	}
	return GroupID{value: value}, nil
}

const MaxPeople = 6
const MinPeople = 1

type People struct {
	value int
}

func (p People) Value() int {
	return p.value
}

var ErrInvalidPeople = errors.New("invalid people")

// NewPeople creates a new people
func NewPeople(value int) (People, error) {
	if value < MinPeople || value > MaxPeople {
		return People{}, fmt.Errorf("%w: %d", ErrInvalidPeople, value)
	}
	return People{value: value}, nil
}

type Group struct {
	id     GroupID
	people People
}

// NewGroup creates a new group
func NewGroup(id int, people int) (Group, error) {
	idValueObject, err := NewGroupID(id)
	if err != nil {
		return Group{}, err
	}
	peopleValueObject, err := NewPeople(people)
	if err != nil {
		return Group{}, err
	}
	return Group{
		id:     idValueObject,
		people: peopleValueObject,
	}, nil
}

// ID returns the group id
func (g Group) ID() GroupID {
	return g.id
}

// People returns the group people
func (g Group) People() People {
	return g.people
}
