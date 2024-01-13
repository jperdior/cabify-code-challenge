package locate

import (
	"cabify-code-challenge/internal/carpool"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLocateUseCase_Locate(t *testing.T) {

	locateUseCase := NewLocateUseCase()

	t.Run("given a carpool with a group in a post_journey, it should return the car", func(t *testing.T) {
		car1, err := carpool.NewCar(1, 4, 0)
		require.NoError(t, err)
		car2, err := carpool.NewCar(2, 5, 0)
		require.NoError(t, err)
		car3, err := carpool.NewCar(3, 5, 0)
		require.NoError(t, err)
		cars := map[carpool.CarID]carpool.Car{
			car1.ID(): car1, car2.ID(): car2, car3.ID(): car3,
		}
		carsByAvailableSeats := map[carpool.AvailableSeats]map[carpool.CarID]carpool.Car{
			car1.AvailableSeats(): {car1.ID(): car1}, car2.AvailableSeats(): {car2.ID(): car2, car3.ID(): car3},
		}
		group, err := carpool.NewGroup(1, 2)
		require.NoError(t, err)
		groups := map[carpool.GroupID]carpool.Group{group.ID(): group}
		journey, err := carpool.NewJourney(group, car2)
		require.NoError(t, err)
		journeys := map[carpool.GroupID]carpool.Journey{group.ID(): journey}
		carPool := carpool.NewTestCarPoolWithCarsAndJourneys(
			cars,
			carsByAvailableSeats,
			groups,
			journeys,
		)

		locationCar, err := locateUseCase.Locate(carPool, group.ID().Value())
		expectedCar, err := carpool.NewCar(car2.ID().Value(), car2.Seats().Value(), car2.AvailableSeats().Value())

		require.NoError(t, err)
		// should be in car2
		assert.Equal(t, expectedCar, locationCar)
	})

	t.Run("given a carpool with a group waiting, it should return an empty car", func(t *testing.T) {
		group, err := carpool.NewGroup(1, 2)
		require.NoError(t, err)
		groups := map[carpool.GroupID]carpool.Group{group.ID(): group}
		waitingGroups := []carpool.Group{group}
		waitingGroupsIndexHash := map[carpool.GroupID]int{group.ID(): 0}

		carPool := carpool.NewTestCarPoolWithWaitingGroups(
			groups,
			waitingGroups,
			waitingGroupsIndexHash,
		)

		locationCar, err := locateUseCase.Locate(carPool, group.ID().Value())
		require.NoError(t, err)

		assert.Equal(t, carpool.Car{}, locationCar)
	})

	t.Run("given a carpool without the group we want to locate it should return an error", func(t *testing.T) {
		carPool := carpool.NewCarPool()
		_, err := locateUseCase.Locate(carPool, 1)
		assert.True(t, errors.Is(err, carpool.ErrGroupNotFound))
	})
}
