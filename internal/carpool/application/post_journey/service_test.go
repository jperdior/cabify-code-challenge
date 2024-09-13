package post_journey

import (
	"cabify-code-challenge/internal/carpool/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_JourneyUseCase(t *testing.T) {

	t.Run("given a carpool without cars available it should have a group waiting ", func(t *testing.T) {
		carPool := domain.NewTestCarPoolWithoutCars()
		service := NewJourneyService(carPool)
		err := service.Execute(2, 3)

		require.NoError(t, err)

		// There should be 1 group
		assert.Equal(t, 1, len(carPool.GetGroups()))
		// There should be 1 group waiting
		assert.Equal(t, 1, len(carPool.GetWaitingGroups()))
		groupId, _ := domain.NewGroupID(2)
		// The group should be in the index 0 of the waiting groups queue
		assert.Equal(t, 0, carPool.GetWaitingGroupsIndexHash()[groupId])

	})

	t.Run("given a carpool with cars available it should have a post_journey and in the only car with enough space", func(t *testing.T) {
		car1, err := domain.NewCar(1, 4, 0)
		require.NoError(t, err)
		car2, err := domain.NewCar(2, 4, 0)
		require.NoError(t, err)
		car3, err := domain.NewCar(3, 5, 0)
		require.NoError(t, err)

		cars := map[domain.CarID]domain.Car{
			car1.ID(): car1,
			car2.ID(): car2,
			car3.ID(): car3,
		}
		carsByAvailableSeats := map[domain.AvailableSeats]map[domain.CarID]domain.Car{
			car1.AvailableSeats(): {car1.ID(): car1, car2.ID(): car2},
			car3.AvailableSeats(): {car3.ID(): car3},
		}

		carPool := domain.NewTestCarPoolWithCars(cars, carsByAvailableSeats)

		group, err := domain.NewGroup(2, 5)
		require.NoError(t, err)
		service := NewJourneyService(carPool)
		err = service.Execute(group.ID().Value(), group.People().Value())
		require.NoError(t, err)

		require.NoError(t, err)
		// There should be 1 group
		assert.Equal(t, 1, len(carPool.GetGroups()))
		// There should be 0 groups waiting
		assert.Equal(t, 0, len(carPool.GetWaitingGroups()))
		// There should be 1 post_journey
		assert.Equal(t, 1, len(carPool.GetJourneys()))
		// The post_journey should be in the only car with enough space which is car with id 3
		groupId, _ := domain.NewGroupID(2)
		var car = carPool.GetJourneys()[groupId].Car()
		assert.Equal(t, 3, car.ID().Value())
	})

	t.Run("given a carpool with a car journeying but with enough space, there should be 2 journeys in the same car", func(t *testing.T) {

		car, err := domain.NewCar(1, 4, 0)
		require.NoError(t, err)
		group1, err := domain.NewGroup(1, 2)
		require.NoError(t, err)
		group2, err := domain.NewGroup(2, 2)
		require.NoError(t, err)
		journey, err := domain.NewJourney(group1, car)
		require.NoError(t, err)

		cars := map[domain.CarID]domain.Car{
			car.ID(): car,
		}
		carsByAvailableSeats := map[domain.AvailableSeats]map[domain.CarID]domain.Car{
			car.AvailableSeats(): {car.ID(): car},
		}
		groups := map[domain.GroupID]domain.Group{
			group1.ID(): group1,
		}
		journeys := map[domain.GroupID]domain.Journey{
			group1.ID(): journey,
		}

		carPool := domain.NewTestCarPoolWithCarsAndJourneys(cars, carsByAvailableSeats, groups, journeys)
		service := NewJourneyService(carPool)
		err = service.Execute(group2.ID().Value(), group2.People().Value())
		require.NoError(t, err)

		// There should be 2 groups
		assert.Equal(t, 2, len(carPool.GetGroups()))
		// There should be 0 groups waiting
		assert.Equal(t, 0, len(carPool.GetWaitingGroups()))
		// There should be 2 journeys
		assert.Equal(t, 2, len(carPool.GetJourneys()))
		// The 2 groups should be in the same car
		var group1Car = carPool.GetJourneys()[group1.ID()].Car()
		assert.Equal(t, car.ID(), group1Car.ID())
		var group2Car = carPool.GetJourneys()[group2.ID()].Car()
		assert.Equal(t, car.ID(), group2Car.ID())
	})

	t.Run("given a carpool with 2 cars and 1 group fitting a whole car, a new post_journey should be in the other car", func(t *testing.T) {

		car1, err := domain.NewCar(1, 4, 0)
		require.NoError(t, err)
		car2, err := domain.NewCar(2, 5, 0)
		require.NoError(t, err)
		group1, err := domain.NewGroup(1, 5)
		require.NoError(t, err)
		group2, err := domain.NewGroup(2, 2)
		require.NoError(t, err)
		journey, err := domain.NewJourney(group1, car2)
		require.NoError(t, err)

		cars := map[domain.CarID]domain.Car{
			car1.ID(): car1,
			car2.ID(): car2,
		}
		carsByAvailableSeats := map[domain.AvailableSeats]map[domain.CarID]domain.Car{
			car1.AvailableSeats(): {car1.ID(): car1},
			car2.AvailableSeats(): {car2.ID(): car2},
		}
		groups := map[domain.GroupID]domain.Group{
			group1.ID(): group1,
		}
		journeys := map[domain.GroupID]domain.Journey{
			group1.ID(): journey,
		}

		carPool := domain.NewTestCarPoolWithCarsAndJourneys(cars, carsByAvailableSeats, groups, journeys)

		service := NewJourneyService(carPool)

		err = service.Execute(group2.ID().Value(), group2.People().Value())
		require.NoError(t, err)

		// There should be 2 groups
		assert.Equal(t, 2, len(carPool.GetGroups()))
		// There should be 0 groups waiting
		assert.Equal(t, 0, len(carPool.GetWaitingGroups()))
		// There should be 2 journeys
		assert.Equal(t, 2, len(carPool.GetJourneys()))
		// Group 1 should be in car 2
		var group1Car = carPool.GetJourneys()[group1.ID()].Car()
		assert.Equal(t, car2.ID(), group1Car.ID())
		// Group 2 should be in car 1
		var group2Car = carPool.GetJourneys()[group2.ID()].Car()
		assert.Equal(t, car1.ID(), group2Car.ID())
	})
}
