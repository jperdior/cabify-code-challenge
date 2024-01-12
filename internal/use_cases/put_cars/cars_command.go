package put_cars

import (
	"cabify-code-challenge/internal/carpool"
	"cabify-code-challenge/kit/command"
	"context"
	"errors"
)

const PutCarsCommandType command.Type = "put_cars"

type CarDTO struct {
	id    int
	seats int
}

func NewCarDTO(id int, seats int) CarDTO {
	return CarDTO{id: id, seats: seats}
}

func getID(car CarDTO) int {
	return car.id
}

func getSeats(car CarDTO) int {
	return car.seats
}

type PutCarsCommand struct {
	cars []CarDTO
}

func NewPutCarsCommand(cars []CarDTO) PutCarsCommand {
	return PutCarsCommand{
		cars: cars,
	}
}

func (c PutCarsCommand) Type() command.Type {
	return PutCarsCommandType
}

type PutCarsCommandHandler struct {
	useCase PuttingCarsUseCase
}

func NewPutCarsCommandHandler(useCase PuttingCarsUseCase) PutCarsCommandHandler {
	return PutCarsCommandHandler{useCase: useCase}
}

// Handle implements the command.Handler interface
func (h PutCarsCommandHandler) Handle(context context.Context, command command.Command) error {
	putCarsCommand, ok := command.(PutCarsCommand)
	if !ok {
		return errors.New("unexpected command")
	}
	var cars []carpool.Car
	for _, car := range putCarsCommand.cars {
		newCar, err := carpool.NewCar(getID(car), getSeats(car))
		if err != nil {
			return err
		}
		cars = append(cars, newCar)
	}
	return h.useCase.PutCars(
		context,
		cars,
	)
}
