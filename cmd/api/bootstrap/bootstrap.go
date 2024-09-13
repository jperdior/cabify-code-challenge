package bootstrap

import (
	"cabify-code-challenge/internal/carpool/application/dropoff"
	"cabify-code-challenge/internal/carpool/application/locate"
	"cabify-code-challenge/internal/carpool/application/post_journey"
	"cabify-code-challenge/internal/carpool/application/put_cars"
	"cabify-code-challenge/internal/carpool/application/retry_journey"
	"cabify-code-challenge/internal/carpool/domain"
	"cabify-code-challenge/internal/platform/bus/inmemory"
	"cabify-code-challenge/internal/platform/server"
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
		carPool    = domain.NewCarPool()
	)

	putCarsService := put_cars.NewPutCarsService(eventBus, carPool)
	puttingCarsCommandHandler := put_cars.NewPutCarsCommandHandler(putCarsService)
	commandBus.Register(put_cars.PutCarsCommandType, puttingCarsCommandHandler)

	journeyService := post_journey.NewJourneyService(carPool)
	journeyCommandHandler := post_journey.NewCreatingJourneyCommandHandler(journeyService)
	commandBus.Register(post_journey.CreatingJourneyCommandType, journeyCommandHandler)

	dropOffService := dropoff.NewDropOffService(eventBus, carPool)
	dropOffCommandHandler := dropoff.NewDropOffCommandHandler(dropOffService)
	commandBus.Register(dropoff.DropOffCommandType, dropOffCommandHandler)

	locateService := locate.NewLocateService(carPool)
	locateQueryHandler := locate.NewLocateQueryHandler(locateService)
	queryBus.Register(locate.LocateQueryType, locateQueryHandler)

	retryJourneyUseCase := retry_journey.NewRetryJourneyService(carPool)
	eventBus.Subscribe(domain.JourneyDroppedEventType, retry_journey.NewRetryJourneysOnJourneyDropped(retryJourneyUseCase))

	ctx, srv := server.New(context.Background(), cfg.Host, cfg.Port, cfg.ShutdownTimeout, commandBus, queryBus, eventBus)
	return srv.Run(ctx)
}

type config struct {
	// Server configuration
	Host            string        `default:""`
	Port            uint          `default:"9091"`
	ShutdownTimeout time.Duration `default:"10s"`
}
