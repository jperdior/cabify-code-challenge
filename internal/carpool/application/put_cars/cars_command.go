package put_cars

import (
	"cabify-code-challenge/internal/carpool/domain"
	"cabify-code-challenge/kit/command"
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
	service PutCarsService
}

func NewPutCarsCommandHandler(service PutCarsService) PutCarsCommandHandler {
	return PutCarsCommandHandler{service: service}
}

// Handle implements the command.Handler interface
func (h PutCarsCommandHandler) Handle(command command.Command) error {
	putCarsCommand, ok := command.(PutCarsCommand)
	if !ok {
		return errors.New("unexpected command")
	}
	var cars []domain.Car
	for _, car := range putCarsCommand.cars {
		newCar, err := domain.NewCar(getID(car), getSeats(car), 0)
		if err != nil {
			return err
		}
		cars = append(cars, newCar)
	}
	return h.service.Execute(
		cars,
	)
}
