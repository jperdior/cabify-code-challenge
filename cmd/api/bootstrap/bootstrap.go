package bootstrap

import (
	"cabify-code-challenge/internal/carpool"
	"cabify-code-challenge/internal/platform/bus/inmemory"
	"cabify-code-challenge/internal/platform/server"
	"cabify-code-challenge/internal/use_cases/dropoff"
	"cabify-code-challenge/internal/use_cases/journey"
	"cabify-code-challenge/internal/use_cases/locate"
	"cabify-code-challenge/internal/use_cases/put_cars"
	"cabify-code-challenge/internal/use_cases/retry_journey"
)

const (
	host = ""
	port = 8080
)

func Run() error {

	var (
		commandBus = inmemory.NewCommandBus()
		queryBus   = inmemory.NewQueryBus()
		eventBus   = inmemory.NewEventBus()
	)

	puttingCarsUseCase := put_cars.NewPuttingCarsUseCase(eventBus)
	puttingCarsCommandHandler := put_cars.NewPutCarsCommandHandler(puttingCarsUseCase)
	commandBus.Register(put_cars.PutCarsCommandType, puttingCarsCommandHandler)

	journeyUseCase := journey.NewCreateJourneyUseCase()
	journeyCommandHandler := journey.NewCreatingJourneyCommandHandler(journeyUseCase)
	commandBus.Register(journey.CreatingJourneyCommandType, journeyCommandHandler)

	dropOffUseCase := dropoff.NewDropOffUseCase(eventBus)
	dropOffCommandHandler := dropoff.NewDropOffCommandHandler(dropOffUseCase)
	commandBus.Register(dropoff.DropOffCommandType, dropOffCommandHandler)

	locateUseCase := locate.NewLocateUseCase()
	locateQueryHandler := locate.NewLocateQueryHandler(locateUseCase)
	queryBus.Register(locate.LocateQueryType, locateQueryHandler)

	retryJourneyUseCase := retry_journey.NewRetryJourneyUseCase()
	eventBus.Subscribe(carpool.JourneyDroppedEventType, dropoff.NewRetryJourneysOnJourneyDropped(retryJourneyUseCase))
	eventBus.Subscribe(carpool.CarPutEventType, put_cars.RetryJourneysOnCarPut{})

	carPool := carpool.NewCarPool()

	srv := server.New(host, port, commandBus, queryBus, eventBus, carPool)
	return srv.Run()
}
