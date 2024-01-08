package carpool

type Group struct {
	id     int
	people int
}

// NewGroup creates a new group
func NewGroup(id int, people int) Group {
	return Group{
		id:     id,
		people: people,
	}
}

// ID returns the group id
func (g Group) ID() int {
	return g.id
}

// People returns the group people
func (g Group) People() int {
	return g.people
}
