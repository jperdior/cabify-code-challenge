package cars

import (
	"cabify-code-challenge/internal/carpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_CarsUseCase_PutCars(t *testing.T) {

	puttingCarsService := NewPuttingCarsUseCase()
	carPool := carpool.NewCarPool()

	t.Run("given a set of cars, it should add them to the car pool", func(t *testing.T) {

		car1, _ := carpool.NewCar(1, 3)
		car2, _ := carpool.NewCar(2, 3)
		car3, _ := carpool.NewCar(3, 5)

		err := puttingCarsService.PutCars(carPool, []carpool.Car{car1, car2, car3})
		require.NoError(t, err)

		assert.Equal(t, 2, len(carPool.GetCarsBySeat()))
		assert.Equal(t, 2, len(carPool.GetCarsBySeat()[3]))
		assert.Equal(t, 1, len(carPool.GetCarsBySeat()[5]))
	})

	t.Run("given a carpool with cars and journeys, it should overwrite the cars and remove the journeys", func(t *testing.T) {

		car1, _ := carpool.NewCar(1, 3)
		car2, _ := carpool.NewCar(2, 3)
		car3, _ := carpool.NewCar(3, 5)
		carPool.SetCars([]carpool.Car{car1, car2, car3})

		journey, _ := carpool.NewJourney(1, 4)
		carPool.RegisterJourney(journey)

		car4, _ := carpool.NewCar(5, 2)
		car5, _ := carpool.NewCar(6, 3)
		car6, _ := carpool.NewCar(7, 5)

		err := puttingCarsService.PutCars(carPool, []carpool.Car{car4, car5, car6})
		require.NoError(t, err)

		assert.Equal(t, 3, len(carPool.GetCarsBySeat()))
		assert.Equal(t, 1, len(carPool.GetCarsBySeat()[2]))
		assert.Equal(t, 1, len(carPool.GetCarsBySeat()[3]))
		assert.Equal(t, 1, len(carPool.GetCarsBySeat()[5]))

		assert.Equal(t, 0, len(carPool.GetJourneys()))

	})

}
