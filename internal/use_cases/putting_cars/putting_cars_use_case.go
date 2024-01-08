package putting_cars

import (
	"cabify-code-challenge/internal/carpool"
	"context"
)

type PuttingCarsUseCase struct {
}

// NewPuttingCarsUseCase returns the default Service interface implementation
func NewPuttingCarsUseCase() PuttingCarsUseCase {
	return PuttingCarsUseCase{}
}

// PutCars implements the putting_cars CarService interface
func (s PuttingCarsUseCase) PutCars(ctx context.Context, cars []carpool.Car) error {
	return nil
}
