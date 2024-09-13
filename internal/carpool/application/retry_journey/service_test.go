package retry_journey

import (
	"cabify-code-challenge/internal/carpool/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRetryJourneysUseCase_RetryJourneys(t *testing.T) {

	t.Run("given a journey dropped event, a car is now free and a journey for waiting groups should be retried", func(t *testing.T) {
		group, err := domain.NewGroup(1, 3)
		require.NoError(t, err)
		groups := map[domain.GroupID]domain.Group{group.ID(): group}
		waitingGroups := []domain.Group{group}
		waitingGroupsIndexHash := map[domain.GroupID]int{group.ID(): 0}

		car1, err := domain.NewCar(1, 4, 3)
		require.NoError(t, err)

		cars := map[domain.CarID]domain.Car{
			car1.ID(): car1,
		}
		carsByAvailableSeats := map[domain.AvailableSeats]map[domain.CarID]domain.Car{
			car1.AvailableSeats(): {car1.ID(): car1},
		}

		carPool := domain.NewTestCarPoolWithCarsAndWaitingGroups(cars, carsByAvailableSeats, groups, waitingGroups, waitingGroupsIndexHash)

		service := NewRetryJourneyService(carPool)

		//todo eventBusMock.On("Publish", mock.Anything, mock.AnythingOfType("[]event.Event")).Return(nil)

		journeyDroppedEvent := domain.NewJourneyDroppedEvent(car1.ID().Value(), car1.Seats().Value(), car1.AvailableSeats().Value())

		err = service.Execute(journeyDroppedEvent.AvailableSeats())
		require.NoError(t, err)

		// There should be 1 journey
		assert.Equal(t, 1, len(carPool.GetJourneys()))
		// There should be 0 waiting groups
		assert.Equal(t, 0, len(carPool.GetWaitingGroups()))
	})

	t.Run("given a put car event and there are waiting groups, a journey for waiting groups should be retried", func(t *testing.T) {
		group, err := domain.NewGroup(1, 3)
		require.NoError(t, err)
		groups := map[domain.GroupID]domain.Group{group.ID(): group}
		waitingGroups := []domain.Group{group}
		waitingGroupsIndexHash := map[domain.GroupID]int{group.ID(): 0}

		car1, err := domain.NewCar(1, 4, 3)
		require.NoError(t, err)

		cars := map[domain.CarID]domain.Car{
			car1.ID(): car1,
		}
		carsByAvailableSeats := map[domain.AvailableSeats]map[domain.CarID]domain.Car{
			car1.AvailableSeats(): {car1.ID(): car1},
		}

		carPool := domain.NewTestCarPoolWithCarsAndWaitingGroups(cars, carsByAvailableSeats, groups, waitingGroups, waitingGroupsIndexHash)

		carPutEvent := domain.NewCarPutEvent(car1.ID().Value(), car1.Seats().Value(), car1.AvailableSeats().Value())

		//todo eventBusMock.On("Publish", mock.Anything, mock.AnythingOfType("[]event.Event")).Return(nil)

		service := NewRetryJourneyService(carPool)

		err = service.Execute(carPutEvent.AvailableSeats())
		require.NoError(t, err)

		// There should be 1 journey
		assert.Equal(t, 1, len(carPool.GetJourneys()))
		// There should be 0 waiting groups
		assert.Equal(t, 0, len(carPool.GetWaitingGroups()))
	})
}
