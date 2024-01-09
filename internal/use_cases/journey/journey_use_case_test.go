package journey

import (
	"cabify-code-challenge/internal/carpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_JourneyUseCase(t *testing.T) {

	creatingJourneyUseCase := NewCreateJourneyUseCase()

	carPool := carpool.NewCarPool()

	t.Run("given a carpool without cars available it should have a group waiting ", func(t *testing.T) {

		err := creatingJourneyUseCase.CreateJourney(carPool, 2, 3)

		require.NoError(t, err)

		assert.Equal(t, 1, len(carPool.GetGroups()))
		assert.Equal(t, 1, len(carPool.GetWaitingGroups()))
		groupId, _ := carpool.NewGroupID(2)
		assert.Equal(t, 0, carPool.GetWaitingGroupsIndexHash()[groupId])

	})

	t.Run("given a carpool with cars available it should have a journey and in the only car with enough space", func(t *testing.T) {

		car1, err := carpool.NewCar(1, 3)
		require.NoError(t, err)
		car2, err := carpool.NewCar(2, 3)
		require.NoError(t, err)
		car3, err := carpool.NewCar(3, 6)
		require.NoError(t, err)
		var cars []carpool.Car

		cars = append(cars, car1)
		cars = append(cars, car2)
		cars = append(cars, car3)

		carPool.SetCars(cars)

		err = creatingJourneyUseCase.CreateJourney(carPool, 2, 4)
		if err != nil {
			return
		}

		require.NoError(t, err)

		assert.Equal(t, 1, len(carPool.GetGroups()))
		assert.Equal(t, 1, len(carPool.GetJourneys()))
		groupId, _ := carpool.NewGroupID(2)
		var carId = carPool.GetJourneys()[groupId].CarID()
		assert.Equal(t, 3, carId.Value())
	})
}
