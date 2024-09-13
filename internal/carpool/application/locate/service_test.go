package locate

import (
	"cabify-code-challenge/internal/carpool/domain"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLocateUseCase_Locate(t *testing.T) {

	t.Run("given a carpool with a group in a journey, it should return the car", func(t *testing.T) {
		car1, err := domain.NewCar(1, 4, 0)
		require.NoError(t, err)
		car2, err := domain.NewCar(2, 5, 0)
		require.NoError(t, err)
		car3, err := domain.NewCar(3, 5, 0)
		require.NoError(t, err)
		cars := map[domain.CarID]domain.Car{
			car1.ID(): car1, car2.ID(): car2, car3.ID(): car3,
		}
		carsByAvailableSeats := map[domain.AvailableSeats]map[domain.CarID]domain.Car{
			car1.AvailableSeats(): {car1.ID(): car1}, car2.AvailableSeats(): {car2.ID(): car2, car3.ID(): car3},
		}
		group, err := domain.NewGroup(1, 2)
		require.NoError(t, err)
		groups := map[domain.GroupID]domain.Group{group.ID(): group}
		journey, err := domain.NewJourney(group, car2)
		require.NoError(t, err)
		journeys := map[domain.GroupID]domain.Journey{group.ID(): journey}
		carPool := domain.NewTestCarPoolWithCarsAndJourneys(
			cars,
			carsByAvailableSeats,
			groups,
			journeys,
		)

		service := NewLocateService(carPool)

		locationCar, err := service.Execute(group.ID().Value())
		expectedCar, err := domain.NewCar(car2.ID().Value(), car2.Seats().Value(), car2.AvailableSeats().Value())

		require.NoError(t, err)
		// should be in car2
		assert.Equal(t, expectedCar, locationCar)
	})

	t.Run("given a carpool with a group waiting, it should return an empty car", func(t *testing.T) {
		group, err := domain.NewGroup(1, 2)
		require.NoError(t, err)
		groups := map[domain.GroupID]domain.Group{group.ID(): group}
		waitingGroups := []domain.Group{group}
		waitingGroupsIndexHash := map[domain.GroupID]int{group.ID(): 0}

		carPool := domain.NewTestCarPoolWithWaitingGroups(
			groups,
			waitingGroups,
			waitingGroupsIndexHash,
		)

		service := NewLocateService(carPool)

		locationCar, err := service.Execute(group.ID().Value())
		require.NoError(t, err)

		assert.Equal(t, domain.Car{}, locationCar)
	})

	t.Run("given a carpool without the group we want to locate it should return an error", func(t *testing.T) {
		carPool := domain.NewCarPool()
		service := NewLocateService(carPool)
		_, err := service.Execute(1)
		assert.True(t, errors.Is(err, domain.ErrGroupNotFound))
	})
}
