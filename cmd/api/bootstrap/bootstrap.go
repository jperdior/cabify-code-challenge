package bootstrap

import (
	"cabify-code-challenge/internal/platform/bus/inmemory"
	"cabify-code-challenge/internal/platform/server"
	"cabify-code-challenge/internal/use_cases/putting_cars"
)

const (
	host = "localhost"
	port = 8080
)

func Run() error {

	var (
		commandBus = inmemory.NewCommandBus()
	)

	puttingCarsUseCase := putting_cars.NewPuttingCarsUseCase()
	puttingCarsCommandHandler := putting_cars.NewPutCarsCommandHandler(puttingCarsUseCase)
	commandBus.Register(putting_cars.PutCarsCommandType, puttingCarsCommandHandler)

	srv := server.New(host, port, commandBus)
	return srv.Run()
}
