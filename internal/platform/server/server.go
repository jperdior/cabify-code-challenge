package server

import (
	"cabify-code-challenge/internal/platform/server/handler/cars"
	"cabify-code-challenge/internal/platform/server/handler/groups"
	"cabify-code-challenge/internal/platform/server/handler/status"
	"cabify-code-challenge/internal/use_cases/putting_cars"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	//deps
	puttingCarsUseCase putting_cars.PuttingCarsUseCase
}

func New(host string, port uint, puttingCarsUseCase putting_cars.PuttingCarsUseCase) Server {
	srv := Server{
		httpAddr: fmt.Sprintf("%s:%d", host, port),
		engine:   gin.New(),

		puttingCarsUseCase: puttingCarsUseCase,
	}

	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/status", status.StatusHandler())
	s.engine.PUT("/cars", cars.PutCarsHandler(s.puttingCarsUseCase))
	s.engine.POST("/journey", groups.PostJourneyHandler())
	s.engine.POST("/dropoff", groups.PostDropOffHandler())
	s.engine.POST("/locate", groups.PostLocateHandler())
}
