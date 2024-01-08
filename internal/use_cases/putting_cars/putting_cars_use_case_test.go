package putting_cars

import (
	"cabify-code-challenge/internal/carpool"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_PuttingCarsService_PutCars(t *testing.T) {

	puttingCarsService := NewPuttingCarsUseCase()

	err := puttingCarsService.PutCars(context.Background(), []carpool.Car{})

	assert.NoError(t, err)
}
