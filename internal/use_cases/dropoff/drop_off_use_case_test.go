package dropoff

import (
	"cabify-code-challenge/internal/carpool"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDropOffUseCase_DropOff(t *testing.T) {

	dropOffUseCase := NewDropOffUseCase()
	carPool := carpool.NewCarPool()

	t.Run("given a carPool with a group on a journey it should deregister the group", func(t *testing.T) {

		car1, _ := carpool.NewCar(1, 3)
		car2, _ := carpool.NewCar(2, 3)
		car3, _ := carpool.NewCar(3, 6)
		carPool.SetCars([]carpool.Car{car1, car2, car3})

		group, err := carpool.NewGroup(1, 3)
		require.NoError(t, err)
		carPool.AddGroup(group)

		journey, err := carpool.NewJourney(1, 6)
		require.NoError(t, err)
		carPool.RegisterJourney(journey)

		err = dropOffUseCase.DropOff(carPool, 1)
		require.NoError(t, err)

		assert.Equal(t, 0, len(carPool.GetGroups()))
		assert.Equal(t, 0, len(carPool.GetJourneys()))
	})

	t.Run("given a carPool with a group waiting it should deregister the group", func(t *testing.T) {

		group, err := carpool.NewGroup(1, 3)
		require.NoError(t, err)
		carPool.AddGroup(group)
		carPool.AddWaitingGroup(group)

		err = dropOffUseCase.DropOff(carPool, 1)
		require.NoError(t, err)

		assert.Equal(t, 0, len(carPool.GetGroups()))
		assert.Equal(t, 0, len(carPool.GetWaitingGroups()))
		assert.Equal(t, 0, len(carPool.GetWaitingGroupsIndexHash()))
	})

	t.Run("given a carPool without a group it should return an error", func(t *testing.T) {

		err := dropOffUseCase.DropOff(carPool, 1)

		require.True(t, errors.Is(err, carpool.ErrGroupNotFound))
	})

}
