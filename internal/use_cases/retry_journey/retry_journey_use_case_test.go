package retry_journey

import (
	"cabify-code-challenge/internal/carpool"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRetryJourneysUseCase_RetryJourneys(t *testing.T) {

	retryJourneyUseCase := NewRetryJourneyUseCase()

	t.Run("given a journey dropped event, a car is now free and a journey for waiting groups should be retried", func(t *testing.T) {
		group, err := carpool.NewGroup(1, 3)
		require.NoError(t, err)
		groups := map[carpool.GroupID]carpool.Group{group.ID(): group}
		waitingGroups := []carpool.Group{group}
		waitingGroupsIndexHash := map[carpool.GroupID]int{group.ID(): 0}

		car1, err := carpool.NewCar(1, 4, 3)
		require.NoError(t, err)

		cars := map[carpool.CarID]carpool.Car{
			car1.ID(): car1,
		}
		carsByAvailableSeats := map[carpool.AvailableSeats]map[carpool.CarID]carpool.Car{
			car1.AvailableSeats(): {car1.ID(): car1},
		}

		carPool := carpool.NewTestCarPoolWithCarsAndWaitingGroups(cars, carsByAvailableSeats, groups, waitingGroups, waitingGroupsIndexHash)
		ctx := context.WithValue(context.Background(), "carPool", carPool)

		journeyDroppedEvent := carpool.NewJourneyDroppedEvent(car1.ID().Value(), car1.Seats().Value(), car1.AvailableSeats().Value())

		err = retryJourneyUseCase.RetryJourneys(ctx, journeyDroppedEvent)
		require.NoError(t, err)

		// There should be 1 journey
		assert.Equal(t, 1, len(carPool.GetJourneys()))
		// There should be 0 waiting groups
		assert.Equal(t, 0, len(carPool.GetWaitingGroups()))
	})

	t.Run("given a put car event and there are waiting groups, a journey for waiting groups should be retried", func(t *testing.T) {
		group, err := carpool.NewGroup(1, 3)
		require.NoError(t, err)
		groups := map[carpool.GroupID]carpool.Group{group.ID(): group}
		waitingGroups := []carpool.Group{group}
		waitingGroupsIndexHash := map[carpool.GroupID]int{group.ID(): 0}

		car1, err := carpool.NewCar(1, 4, 3)
		require.NoError(t, err)

		cars := map[carpool.CarID]carpool.Car{
			car1.ID(): car1,
		}
		carsByAvailableSeats := map[carpool.AvailableSeats]map[carpool.CarID]carpool.Car{
			car1.AvailableSeats(): {car1.ID(): car1},
		}

		carPool := carpool.NewTestCarPoolWithCarsAndWaitingGroups(cars, carsByAvailableSeats, groups, waitingGroups, waitingGroupsIndexHash)
		ctx := context.WithValue(context.Background(), "carPool", carPool)

		carPutEvent := carpool.NewCarPutEvent(car1.ID().Value(), car1.Seats().Value(), car1.AvailableSeats().Value())

		err = retryJourneyUseCase.RetryJourneys(ctx, carPutEvent)
		require.NoError(t, err)

		// There should be 1 journey
		assert.Equal(t, 1, len(carPool.GetJourneys()))
		// There should be 0 waiting groups
		assert.Equal(t, 0, len(carPool.GetWaitingGroups()))
	})
}
