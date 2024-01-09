package bootstrap

import (
	"cabify-code-challenge/internal/carpool"
	"cabify-code-challenge/internal/platform/bus/inmemory"
	"cabify-code-challenge/internal/platform/server"
	"cabify-code-challenge/internal/use_cases/journey"
	"cabify-code-challenge/internal/use_cases/put_cars"
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

	puttingCarsUseCase := put_cars.NewPuttingCarsUseCase()
	puttingCarsCommandHandler := put_cars.NewPutCarsCommandHandler(puttingCarsUseCase)
	commandBus.Register(put_cars.PutCarsCommandType, puttingCarsCommandHandler)
	createJourneyUseCase := journey.NewCreateJourneyUseCase()
	createJourneyCommandHandler := journey.NewCreatingJourneyCommandHandler(createJourneyUseCase)
	commandBus.Register(journey.CreatingJourneyCommandType, createJourneyCommandHandler)

	carPool := carpool.NewCarPool()

	srv := server.New(host, port, commandBus, queryBus, carPool)
	return srv.Run()
}
