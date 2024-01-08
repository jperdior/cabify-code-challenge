package creating_journey

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_PuttingCarsService_PutCars(t *testing.T) {

	creatingJourneyService := NewCreateJourneyUseCase()

	err := creatingJourneyService.CreateJourney(2, 2)

	assert.NoError(t, err)
}
