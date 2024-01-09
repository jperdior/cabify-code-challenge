package putting_cars

import (
	"cabify-code-challenge/internal/carpool"
	"context"
	"fmt"
)

type PuttingCarsUseCase struct {
}

// NewPuttingCarsUseCase returns the default Service interface implementation
func NewPuttingCarsUseCase() PuttingCarsUseCase {
	return PuttingCarsUseCase{}
}

// PutCars implements the putting_cars CarService interface
func (s PuttingCarsUseCase) PutCars(ctx context.Context, cars []carpool.Car) error {
	for _, car := range cars {
		fmt.Printf("Car with ID %d and %d seats\n", car.ID(), car.Seats())
	}
	return nil
}
