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

	carPool := carpool.NewCarPool()
	eventBusMock := new(eventmocks.Bus)
	dropOffUseCase := NewDropOffUseCase(eventBusMock)
	ctx := context.WithValue(context.Background(), "carPool", carPool)

	t.Run("given a carPool with a group on a journey it should deregister the group", func(t *testing.T) {

		car1, err := carpool.NewCar(1, 4)
		require.NoError(t, err)
		car2, err := carpool.NewCar(2, 4)
		require.NoError(t, err)
		car3, err := carpool.NewCar(3, 6)
		require.NoError(t, err)
		carPool.SetCars([]carpool.Car{car1, car2, car3})

		group, err := carpool.NewGroup(1, 3)
		require.NoError(t, err)
		err = carPool.AddGroup(group)
		require.NoError(t, err)

		_, err = carpool.NewJourney(group, car1)
		require.NoError(t, err)

		eventBusMock.On("Publish", mock.Anything, mock.AnythingOfType("[]event.Event")).Return(nil)

		err = dropOffUseCase.DropOff(ctx, 1)
		require.NoError(t, err)

		assert.Equal(t, 0, len(carPool.GetGroups()))
		assert.Equal(t, 0, len(carPool.GetJourneys()))
		eventBusMock.AssertExpectations(t)
	})

	t.Run("given a carPool with a group waiting it should deregister the group", func(t *testing.T) {

		group, err := carpool.NewGroup(1, 3)
		require.NoError(t, err)
		err = carPool.AddGroup(group)
		require.NoError(t, err)
		carPool.AddWaitingGroup(group)

		eventBusMock.On("Publish", mock.Anything, mock.AnythingOfType("[]event.Event")).Return(nil)

		err = dropOffUseCase.DropOff(ctx, 1)
		require.NoError(t, err)

		assert.Equal(t, 0, len(carPool.GetGroups()))
		assert.Equal(t, 0, len(carPool.GetWaitingGroups()))
		assert.Equal(t, 0, len(carPool.GetWaitingGroupsIndexHash()))
		eventBusMock.AssertExpectations(t)
	})

	t.Run("given a carPool without a group it should return an error", func(t *testing.T) {

		err := dropOffUseCase.DropOff(ctx, 1)

		require.True(t, errors.Is(err, carpool.ErrGroupNotFound))
	})

}
