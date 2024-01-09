package bootstrap

import (
	"cabify-code-challenge/internal/carpool"
	"cabify-code-challenge/internal/platform/bus/inmemory"
	"cabify-code-challenge/internal/platform/server"
	"cabify-code-challenge/internal/use_cases/cars"
	"cabify-code-challenge/internal/use_cases/journey"
)

const (
	host = "localhost"
	port = 8080
)

func Run() error {

	var (
		commandBus = inmemory.NewCommandBus()
		queryBus   = inmemory.NewQueryBus()
	)

	puttingCarsUseCase := cars.NewPuttingCarsUseCase()
	puttingCarsCommandHandler := cars.NewPutCarsCommandHandler(puttingCarsUseCase)
	commandBus.Register(cars.PutCarsCommandType, puttingCarsCommandHandler)
	createJourneyUseCase := journey.NewCreateJourneyUseCase()
	createJourneyCommandHandler := journey.NewCreatingJourneyCommandHandler(createJourneyUseCase)
	commandBus.Register(journey.CreatingJourneyCommandType, createJourneyCommandHandler)

	carPool := carpool.NewCarPool()

	srv := server.New(host, port, commandBus, queryBus, carPool)
	return srv.Run()
}
