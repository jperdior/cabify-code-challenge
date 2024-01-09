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
	carPool := carpool.NewCarPool()

	t.Run("given a carpool with a car with a group, it should return the car", func(t *testing.T) {
		car1, err := carpool.NewCar(1, 1)
		car2, err := carpool.NewCar(2, 3)
		car3, err := carpool.NewCar(3, 5)
		require.NoError(t, err)
		carPool.SetCars([]carpool.Car{car1, car2, car3})

		group, err := carpool.NewGroup(1, 2)
		require.NoError(t, err)
		carPool.AddGroup(group)

		journey, err := carpool.NewJourney(1, 2)
		require.NoError(t, err)
		carPool.RegisterJourney(journey)

		locationCar, err := locateUseCase.Locate(carPool, 1)
		require.NoError(t, err)

		assert.Equal(t, car2, locationCar)
	})

	t.Run("given a carpool with a group waiting, it should return an empty car", func(t *testing.T) {
		group, err := carpool.NewGroup(1, 2)
		require.NoError(t, err)
		carPool.AddGroup(group)
		carPool.AddWaitingGroup(group)

		locationCar, err := locateUseCase.Locate(carPool, 1)
		require.NoError(t, err)

		assert.Equal(t, carpool.Car{}, locationCar)
	})

	t.Run("given a carpool without the group we want to locate it should return an error", func(t *testing.T) {
		_, err := locateUseCase.Locate(carPool, 1)
		assert.True(t, errors.Is(err, carpool.ErrGroupNotFound))
	})
}
