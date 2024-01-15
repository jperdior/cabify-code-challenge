package bootstrap

import (
	"cabify-code-challenge/internal/carpool"
	"cabify-code-challenge/internal/platform/bus/inmemory"
	"cabify-code-challenge/internal/platform/server"
	"cabify-code-challenge/internal/use_cases/dropoff"
	"cabify-code-challenge/internal/use_cases/locate"
	"cabify-code-challenge/internal/use_cases/postjourney"
	"cabify-code-challenge/internal/use_cases/put_cars"
	"cabify-code-challenge/internal/use_cases/retry_journey"
	"context"
	"github.com/kelseyhightower/envconfig"
	"time"
)

func Run() error {

	var cfg config
	err := envconfig.Process("pooling", &cfg)
	if err != nil {
		return err
	}

	var (
		commandBus = inmemory.NewCommandBus()
		queryBus   = inmemory.NewQueryBus()
		eventBus   = inmemory.NewEventBus()
	)

	puttingCarsUseCase := put_cars.NewPuttingCarsUseCase(eventBus)
	puttingCarsCommandHandler := put_cars.NewPutCarsCommandHandler(puttingCarsUseCase)
	commandBus.Register(put_cars.PutCarsCommandType, puttingCarsCommandHandler)

	journeyUseCase := postjourney.NewCreateJourneyUseCase()
	journeyCommandHandler := postjourney.NewCreatingJourneyCommandHandler(journeyUseCase)
	commandBus.Register(postjourney.CreatingJourneyCommandType, journeyCommandHandler)

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

	ctx, srv := server.New(context.Background(), cfg.Host, cfg.Port, cfg.ShutdownTimeout, commandBus, queryBus, eventBus, carPool)
	return srv.Run(ctx)
}

type config struct {
	// Server configuration
	Host            string        `default:""`
	Port            uint          `default:"9091"`
	ShutdownTimeout time.Duration `default:"10s"`
}
