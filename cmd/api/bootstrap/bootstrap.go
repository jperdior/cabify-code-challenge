package bootstrap

import (
	"cabify-code-challenge/internal/carpool"
	"cabify-code-challenge/internal/platform/bus/inmemory"
	"cabify-code-challenge/internal/platform/server"
	"cabify-code-challenge/internal/use_cases/creating_journey"
	"cabify-code-challenge/internal/use_cases/putting_cars"
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

	puttingCarsUseCase := putting_cars.NewPuttingCarsUseCase()
	puttingCarsCommandHandler := putting_cars.NewPutCarsCommandHandler(puttingCarsUseCase)
	commandBus.Register(putting_cars.PutCarsCommandType, puttingCarsCommandHandler)
	createJourneyUseCase := creating_journey.NewCreateJourneyUseCase()
	createJourneyCommandHandler := creating_journey.NewCreatingJourneyCommandHandler(createJourneyUseCase)
	commandBus.Register(creating_journey.CreatingJourneyCommandType, createJourneyCommandHandler)

	carPool := carpool.NewCarPool()

	srv := server.New(host, port, commandBus, queryBus, carPool)
	return srv.Run()
}
