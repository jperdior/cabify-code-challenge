package put_cars

import (
	"cabify-code-challenge/internal/carpool/domain"
	"cabify-code-challenge/kit/event/eventmocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_CarsUseCase_PutCars(t *testing.T) {

	eventBusMock := new(eventmocks.Bus)

	t.Run("given a set of cars, it should add them to the car pool", func(t *testing.T) {

		carPool := domain.NewTestCarPoolWithoutCars()
		car1, err := domain.NewCar(1, 4, 0)
		require.NoError(t, err)
		car2, err := domain.NewCar(2, 4, 0)
		require.NoError(t, err)
		car3, err := domain.NewCar(3, 5, 0)
		require.NoError(t, err)

		service := NewPutCarsService(eventBusMock, carPool)

		eventBusMock.On("Publish", mock.AnythingOfType("[]event.Event")).Return(nil)

		err = service.Execute([]domain.Car{car1, car2, car3})

		require.NoError(t, err)
		//There should be 3 cars
		assert.Equal(t, 3, len(carPool.GetCars()))
		//There should be 2 elements in the map of available seats
		assert.Equal(t, 2, len(carPool.GetCarsByAvailableSeats()))
		//There should be 2 cars with 4 available seats
		fourAvailableSeats, err := domain.NewAvailableSeats(4)
		require.NoError(t, err)
		assert.Equal(t, 2, len(carPool.GetCarsByAvailableSeats()[fourAvailableSeats]))
		//There should be 1 car with 5 available seats
		fiveAvailableSeats, err := domain.NewAvailableSeats(5)
		require.NoError(t, err)
		assert.Equal(t, 1, len(carPool.GetCarsByAvailableSeats()[fiveAvailableSeats]))
		eventBusMock.AssertExpectations(t)
	})

	t.Run("given a carpool with cars and journeys, it should overwrite the put_cars and remove the journeys", func(t *testing.T) {

		car1, err := domain.NewCar(1, 4, 0)
		require.NoError(t, err)
		car2, err := domain.NewCar(2, 4, 0)
		require.NoError(t, err)
		car3, err := domain.NewCar(3, 5, 0)
		require.NoError(t, err)
		group, err := domain.NewGroup(1, 3)
		require.NoError(t, err)
		journey, err := domain.NewJourney(group, car1)
		require.NoError(t, err)

		cars := map[domain.CarID]domain.Car{
			car1.ID(): car1,
			car2.ID(): car2,
			car3.ID(): car3,
		}
		carsByAvailableSeats := map[domain.AvailableSeats]map[domain.CarID]domain.Car{
			car1.AvailableSeats(): {car1.ID(): car1, car2.ID(): car2},
			car3.AvailableSeats(): {car3.ID(): car3},
		}
		groups := map[domain.GroupID]domain.Group{
			group.ID(): group,
		}
		journeys := map[domain.GroupID]domain.Journey{
			group.ID(): journey,
		}

		carPool := domain.NewTestCarPoolWithCarsAndJourneys(cars, carsByAvailableSeats, groups, journeys)

		car4, _ := domain.NewCar(5, 4, 0)
		car5, _ := domain.NewCar(6, 5, 0)

		service := NewPutCarsService(eventBusMock, carPool)

		eventBusMock.On("Publish", mock.AnythingOfType("[]event.Event")).Return(nil)

		err = service.Execute([]domain.Car{car4, car5})
		require.NoError(t, err)

		//There should be 2 cars
		assert.Equal(t, 2, len(carPool.GetCars()))
		//There should be 2 elements in the map of available seats
		assert.Equal(t, 2, len(carPool.GetCarsByAvailableSeats()))
		//There should be 1 car with 4 available seats
		fourAvailableSeats, err := domain.NewAvailableSeats(4)
		assert.Equal(t, 1, len(carPool.GetCarsByAvailableSeats()[fourAvailableSeats]))
		require.NoError(t, err)
		//There should be 1 car with 5 available seats
		fiveAvailableSeats, err := domain.NewAvailableSeats(5)
		assert.Equal(t, 1, len(carPool.GetCarsByAvailableSeats()[fiveAvailableSeats]))
		require.NoError(t, err)
		//There should be no journeys
		assert.Equal(t, 0, len(carPool.GetJourneys()))
		eventBusMock.AssertExpectations(t)
	})

}
