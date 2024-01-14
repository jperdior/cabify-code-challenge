package dropoff

import (
	"cabify-code-challenge/internal/carpool"
	"cabify-code-challenge/kit/event/eventmocks"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDropOffUseCase_DropOff(t *testing.T) {

	eventBusMock := new(eventmocks.Bus)
	dropOffUseCase := NewDropOffUseCase(eventBusMock)

	t.Run("given a carPool with a group on a journey it should remove the group from the journey", func(t *testing.T) {

		car1, err := carpool.NewCar(1, 4, 1)
		require.NoError(t, err)
		car2, err := carpool.NewCar(2, 4, 0)
		require.NoError(t, err)
		car3, err := carpool.NewCar(3, 6, 0)
		require.NoError(t, err)
		group, err := carpool.NewGroup(1, 3)
		require.NoError(t, err)
		journey, err := carpool.NewJourney(group, car1)
		require.NoError(t, err)
		cars := map[carpool.CarID]carpool.Car{
			car1.ID(): car1,
			car2.ID(): car2,
			car3.ID(): car3,
		}
		carsByAvailableSeats := map[carpool.AvailableSeats]map[carpool.CarID]carpool.Car{
			car1.AvailableSeats(): {car1.ID(): car1}, car2.AvailableSeats(): {car2.ID(): car2}, car3.AvailableSeats(): {car3.ID(): car3},
		}
		groups := map[carpool.GroupID]carpool.Group{group.ID(): group}
		journeys := map[carpool.GroupID]carpool.Journey{group.ID(): journey}

		carPool := carpool.NewTestCarPoolWithCarsAndJourneys(
			cars,
			carsByAvailableSeats,
			groups,
			journeys,
		)
		ctx := context.WithValue(context.Background(), "carPool", carPool)

		eventBusMock.On("Publish", mock.Anything, mock.AnythingOfType("[]event.Event")).Return(nil)

		err = dropOffUseCase.DropOff(ctx, 1)
		require.NoError(t, err)

		// there should be no groups
		assert.Equal(t, 0, len(carPool.GetGroups()))
		// there should be no journeys
		assert.Equal(t, 0, len(carPool.GetJourneys()))
		// there should be only 2 types of available seats
		assert.Equal(t, 2, len(carPool.GetCarsByAvailableSeats()))
		// cars with 4 seats available should have 2 cars
		fourAvailableSeats, err := carpool.NewAvailableSeats(4)
		assert.Equal(t, 2, len(carPool.GetCarsByAvailableSeats()[fourAvailableSeats]))
		eventBusMock.AssertExpectations(t)
	})

	t.Run("given a carPool with a group waiting it should deregister the group", func(t *testing.T) {

		group, err := carpool.NewGroup(1, 3)
		require.NoError(t, err)
		groups := map[carpool.GroupID]carpool.Group{group.ID(): group}
		waitingGroups := []carpool.Group{group}
		waitingGroupsIndexHash := map[carpool.GroupID]int{group.ID(): 0}

		carPool := carpool.NewTestCarPoolWithWaitingGroups(
			groups,
			waitingGroups,
			waitingGroupsIndexHash,
		)
		ctx := context.WithValue(context.Background(), "carPool", carPool)

		eventBusMock.On("Publish", mock.Anything, mock.AnythingOfType("[]event.Event")).Return(nil)

		err = dropOffUseCase.DropOff(ctx, 1)
		require.NoError(t, err)

		assert.Equal(t, 0, len(carPool.GetGroups()))
		assert.Equal(t, 0, len(carPool.GetWaitingGroups()))
		assert.Equal(t, 0, len(carPool.GetWaitingGroupsIndexHash()))
		eventBusMock.AssertExpectations(t)
	})

	t.Run("given a carPool without a group it should return an error", func(t *testing.T) {

		carPool := carpool.NewCarPool()
		ctx := context.WithValue(context.Background(), "carPool", carPool)
		err := dropOffUseCase.DropOff(ctx, 1)

		require.True(t, errors.Is(err, carpool.ErrGroupNotFound))
	})

}
