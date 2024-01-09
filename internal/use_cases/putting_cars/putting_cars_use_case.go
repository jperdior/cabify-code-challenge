package putting_cars

import (
	"cabify-code-challenge/internal/carpool"
)

type PuttingCarsUseCase struct {
}

// NewPuttingCarsUseCase returns the default Service interface implementation
func NewPuttingCarsUseCase() PuttingCarsUseCase {
	return PuttingCarsUseCase{}
}

// PutCars sets the cars in the carPool
func (s PuttingCarsUseCase) PutCars(carPool *carpool.CarPool, cars []carpool.Car) error {

	carPool.SetCars(cars)
	return nil
}
