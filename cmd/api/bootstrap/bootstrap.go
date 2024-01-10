package bootstrap

import (
	"cabify-code-challenge/internal/carpool"
	"cabify-code-challenge/internal/platform/bus/inmemory"
	"cabify-code-challenge/internal/platform/server"
	"cabify-code-challenge/internal/use_cases/dropoff"
	"cabify-code-challenge/internal/use_cases/journey"
	"cabify-code-challenge/internal/use_cases/locate"
	"cabify-code-challenge/internal/use_cases/put_cars"
)

const (
	host = ""
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
	journeyUseCase := journey.NewCreateJourneyUseCase()
	journeyCommandHandler := journey.NewCreatingJourneyCommandHandler(journeyUseCase)
	commandBus.Register(journey.CreatingJourneyCommandType, journeyCommandHandler)
	dropOffUseCase := dropoff.NewDropOffUseCase()
	dropOffCommandHandler := dropoff.NewDropOffCommandHandler(dropOffUseCase)
	commandBus.Register(dropoff.DropOffCommandType, dropOffCommandHandler)

	locateUseCase := locate.NewLocateUseCase()
	locateQueryHandler := locate.NewLocateQueryHandler(locateUseCase)
	queryBus.Register(locate.LocateQueryType, locateQueryHandler)

	carPool := carpool.NewCarPool()

	srv := server.New(host, port, commandBus, queryBus, carPool)
	return srv.Run()
}
