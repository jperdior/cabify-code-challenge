package server

import (
	"cabify-code-challenge/internal/platform/server/handler/cars"
	"cabify-code-challenge/internal/platform/server/handler/groups"
	"cabify-code-challenge/internal/platform/server/handler/status"
	"cabify-code-challenge/kit/command"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	//deps
	commandBus command.Bus
}

func New(host string, port uint, commandBus command.Bus) Server {
	srv := Server{
		httpAddr: fmt.Sprintf("%s:%d", host, port),
		engine:   gin.New(),

		//deps
		commandBus: commandBus,
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
	s.engine.PUT("/cars", cars.PutCarsHandler(s.commandBus))
	s.engine.POST("/journey", groups.PostJourneyHandler())
	s.engine.POST("/dropoff", groups.PostDropOffHandler())
	s.engine.POST("/locate", groups.PostLocateHandler())
}
