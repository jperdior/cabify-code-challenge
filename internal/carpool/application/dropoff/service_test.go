package dropoff

import (
	"cabify-code-challenge/internal/carpool/domain"
	"cabify-code-challenge/kit/event/eventmocks"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDropOffUseCase_DropOff(t *testing.T) {

	eventBusMock := new(eventmocks.Bus)

	t.Run("given a carPool with a group on a journey it should remove the group from the journey", func(t *testing.T) {

		car1, err := domain.NewCar(1, 4, 1)
		require.NoError(t, err)
		car2, err := domain.NewCar(2, 4, 0)
		require.NoError(t, err)
		car3, err := domain.NewCar(3, 6, 0)
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
			car1.AvailableSeats(): {car1.ID(): car1}, car2.AvailableSeats(): {car2.ID(): car2}, car3.AvailableSeats(): {car3.ID(): car3},
		}
		groups := map[domain.GroupID]domain.Group{group.ID(): group}
		journeys := map[domain.GroupID]domain.Journey{group.ID(): journey}

		carPool := domain.NewTestCarPoolWithCarsAndJourneys(
			cars,
			carsByAvailableSeats,
			groups,
			journeys,
		)

		service := NewDropOffService(eventBusMock, carPool)

		eventBusMock.On("Publish", mock.AnythingOfType("[]event.Event")).Return(nil)

		err = service.Execute(1)
		require.NoError(t, err)

		// there should be no groups
		assert.Equal(t, 0, len(carPool.GetGroups()))
		// there should be no journeys
		assert.Equal(t, 0, len(carPool.GetJourneys()))
		// cars with 4 seats available should have 2 cars
		fourAvailableSeats, err := domain.NewAvailableSeats(4)
		assert.Equal(t, 2, len(carPool.GetCarsByAvailableSeats()[fourAvailableSeats]))
		eventBusMock.AssertExpectations(t)
	})

	t.Run("given a carPool with a group waiting it should deregister the group", func(t *testing.T) {

		group, err := domain.NewGroup(1, 3)
		require.NoError(t, err)
		groups := map[domain.GroupID]domain.Group{group.ID(): group}
		waitingGroups := []domain.Group{group}
		waitingGroupsIndexHash := map[domain.GroupID]int{group.ID(): 0}

		carPool := domain.NewTestCarPoolWithWaitingGroups(
			groups,
			waitingGroups,
			waitingGroupsIndexHash,
		)

		service := NewDropOffService(eventBusMock, carPool)

		eventBusMock.On("Publish", mock.AnythingOfType("[]event.Event")).Return(nil)

		err = service.Execute(1)
		require.NoError(t, err)

		assert.Equal(t, 0, len(carPool.GetGroups()))
		assert.Equal(t, 0, len(carPool.GetWaitingGroups()))
		assert.Equal(t, 0, len(carPool.GetWaitingGroupsIndexHash()))
		eventBusMock.AssertExpectations(t)
	})

	t.Run("given a carPool without a group it should return an error", func(t *testing.T) {

		carPool := domain.NewCarPool()
		service := NewDropOffService(eventBusMock, carPool)
		err := service.Execute(1)

		require.True(t, errors.Is(err, domain.ErrGroupNotFound))
	})

}
