package server

import (
	"cabify-code-challenge/internal/carpool"
	"cabify-code-challenge/internal/platform/server/handler/cars"
	"cabify-code-challenge/internal/platform/server/handler/groups"
	"cabify-code-challenge/internal/platform/server/handler/status"
	"cabify-code-challenge/kit/command"
	"cabify-code-challenge/kit/query"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	//deps
	commandBus command.Bus
	queryBus   query.Bus
	carPool    *carpool.CarPool
}

func CarPoolMiddleware(carPool *carpool.CarPool) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("carPool", carPool)
		c.Next()
	}
}

func New(host string, port uint, commandBus command.Bus, queryBus query.Bus, carPool *carpool.CarPool) Server {
	srv := Server{
		httpAddr: fmt.Sprintf("%s:%d", host, port),
		engine:   gin.Default(),

		//deps
		commandBus: commandBus,
		queryBus:   queryBus,
		carPool:    carPool,
	}

	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {

	s.engine.Use(CarPoolMiddleware(s.carPool))

	s.engine.GET("/status", status.StatusHandler())
	s.engine.PUT("/cars", cars.PutCarsHandler(s.commandBus))
	s.engine.POST("/journey", groups.PostJourneyHandler(s.commandBus))
	s.engine.POST("/dropoff", groups.PostDropOffHandler(s.commandBus))
	s.engine.POST("/locate", groups.PostLocateHandler(s.queryBus))
}
