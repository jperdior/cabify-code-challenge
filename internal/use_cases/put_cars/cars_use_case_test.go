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

	t.Run("given a set of cars, it should add them to the car pool", func(t *testing.T) {

		carPool := carpool.NewTestCarPoolWithoutCars()
		ctx := context.WithValue(context.Background(), "carPool", carPool)
		car1, err := carpool.NewCar(1, 4, 0)
		require.NoError(t, err)
		car2, err := carpool.NewCar(2, 4, 0)
		require.NoError(t, err)
		car3, err := carpool.NewCar(3, 5, 0)
		require.NoError(t, err)

		eventBusMock.On("Publish", mock.Anything, mock.AnythingOfType("[]event.Event")).Return(nil)

		err = puttingCarsService.PutCars(ctx, []carpool.Car{car1, car2, car3})

		require.NoError(t, err)
		//There should be 3 cars
		assert.Equal(t, 3, len(carPool.GetCars()))
		//There should be 2 elements in the map of available seats
		assert.Equal(t, 2, len(carPool.GetCarsByAvailableSeats()))
		//There should be 2 cars with 4 available seats
		fourAvailableSeats, err := carpool.NewAvailableSeats(4)
		require.NoError(t, err)
		assert.Equal(t, 2, len(carPool.GetCarsByAvailableSeats()[fourAvailableSeats]))
		//There should be 1 car with 5 available seats
		fiveAvailableSeats, err := carpool.NewAvailableSeats(5)
		require.NoError(t, err)
		assert.Equal(t, 1, len(carPool.GetCarsByAvailableSeats()[fiveAvailableSeats]))
		eventBusMock.AssertExpectations(t)
	})

	t.Run("given a carpool with cars and journeys, it should overwrite the put_cars and remove the journeys", func(t *testing.T) {

		car1, err := carpool.NewCar(1, 4, 0)
		require.NoError(t, err)
		car2, err := carpool.NewCar(2, 4, 0)
		require.NoError(t, err)
		car3, err := carpool.NewCar(3, 5, 0)
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
			car1.AvailableSeats(): {car1.ID(): car1, car2.ID(): car2},
			car3.AvailableSeats(): {car3.ID(): car3},
		}
		groups := map[carpool.GroupID]carpool.Group{
			group.ID(): group,
		}
		journeys := map[carpool.GroupID]carpool.Journey{
			group.ID(): journey,
		}

		carPool := carpool.NewTestCarPoolWithCarsAndJourneys(cars, carsByAvailableSeats, groups, journeys)
		ctx := context.WithValue(context.Background(), "carPool", carPool)

		car4, _ := carpool.NewCar(5, 4, 0)
		car5, _ := carpool.NewCar(6, 5, 0)

		eventBusMock.On("Publish", mock.Anything, mock.AnythingOfType("[]event.Event")).Return(nil)

		err = puttingCarsService.PutCars(ctx, []carpool.Car{car4, car5})
		require.NoError(t, err)

		//There should be 2 cars
		assert.Equal(t, 2, len(carPool.GetCars()))
		//There should be 2 elements in the map of available seats
		assert.Equal(t, 2, len(carPool.GetCarsByAvailableSeats()))
		//There should be 1 car with 4 available seats
		fourAvailableSeats, err := carpool.NewAvailableSeats(4)
		assert.Equal(t, 1, len(carPool.GetCarsByAvailableSeats()[fourAvailableSeats]))
		require.NoError(t, err)
		//There should be 1 car with 5 available seats
		fiveAvailableSeats, err := carpool.NewAvailableSeats(5)
		assert.Equal(t, 1, len(carPool.GetCarsByAvailableSeats()[fiveAvailableSeats]))
		require.NoError(t, err)
		//There should be no journeys
		assert.Equal(t, 0, len(carPool.GetJourneys()))
		eventBusMock.AssertExpectations(t)
	})

}
