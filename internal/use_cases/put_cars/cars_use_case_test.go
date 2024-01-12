package put_cars

import (
	"cabify-code-challenge/internal/carpool"
	"cabify-code-challenge/kit/event/eventmocks"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_CarsUseCase_PutCars(t *testing.T) {

	eventBusMock := new(eventmocks.Bus)
	puttingCarsService := NewPuttingCarsUseCase(eventBusMock)
	carPool := carpool.NewCarPool()
	ctx := context.WithValue(context.Background(), "carPool", carPool)

	t.Run("given a set of put_cars, it should add them to the car pool", func(t *testing.T) {

		car1, err := carpool.NewCar(1, 4)
		require.NoError(t, err)
		car2, err := carpool.NewCar(2, 4)
		require.NoError(t, err)
		car3, err := carpool.NewCar(3, 5)
		require.NoError(t, err)

		eventBusMock.On("Publish", mock.Anything, mock.AnythingOfType("[]event.Event")).Return(nil)

		err = puttingCarsService.PutCars(ctx, []carpool.Car{car1, car2, car3})
		require.NoError(t, err)

		assert.Equal(t, 2, len(carPool.GetCarsByAvailableSeats()))
		fourAvailableSeats, err := carpool.NewAvailableSeats(4)
		require.NoError(t, err)
		assert.Equal(t, 2, len(carPool.GetCarsByAvailableSeats()[fourAvailableSeats]))
		fiveAvailableSeats, err := carpool.NewAvailableSeats(5)
		require.NoError(t, err)
		assert.Equal(t, 1, len(carPool.GetCarsByAvailableSeats()[fiveAvailableSeats]))
		eventBusMock.AssertExpectations(t)
	})

	t.Run("given a carpool with put_cars and journeys, it should overwrite the put_cars and remove the journeys", func(t *testing.T) {

		car1, err := carpool.NewCar(1, 4)
		require.NoError(t, err)
		car2, err := carpool.NewCar(2, 4)
		require.NoError(t, err)
		car3, err := carpool.NewCar(3, 5)
		require.NoError(t, err)
		carPool.SetCars([]carpool.Car{car1, car2, car3})
		group, err := carpool.NewGroup(1, 3)
		require.NoError(t, err)
		_, err = carpool.NewJourney(group, car2)
		require.NoError(t, err)

		car4, _ := carpool.NewCar(5, 4)
		car5, _ := carpool.NewCar(6, 5)
		car6, _ := carpool.NewCar(7, 6)

		eventBusMock.On("Publish", mock.Anything, mock.AnythingOfType("[]event.Event")).Return(nil)

		err = puttingCarsService.PutCars(ctx, []carpool.Car{car4, car5, car6})
		require.NoError(t, err)

		assert.Equal(t, 3, len(carPool.GetCarsByAvailableSeats()))
		fourAvailableSeats, err := carpool.NewAvailableSeats(4)
		assert.Equal(t, 1, len(carPool.GetCarsByAvailableSeats()[fourAvailableSeats]))
		require.NoError(t, err)
		fiveAvailableSeats, err := carpool.NewAvailableSeats(5)
		assert.Equal(t, 1, len(carPool.GetCarsByAvailableSeats()[fiveAvailableSeats]))
		require.NoError(t, err)
		sixAvailableSeats, err := carpool.NewAvailableSeats(6)
		assert.Equal(t, 1, len(carPool.GetCarsByAvailableSeats()[sixAvailableSeats]))
		require.NoError(t, err)

		assert.Equal(t, 0, len(carPool.GetJourneys()))
		eventBusMock.AssertExpectations(t)
	})

}
