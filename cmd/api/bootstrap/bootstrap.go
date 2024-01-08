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

	puttingCarsUseCase := putting_cars.NewPuttingCarsUseCase()

	var (
		commandBus = inmemory.NewCommandBus()
	)

	srv := server.New(host, port, puttingCarsUseCase)
	return srv.Run()
}
