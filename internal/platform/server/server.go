package server

import (
	"cabify-code-challenge/internal/carpool"
	"cabify-code-challenge/internal/platform/server/handler/cars"
	"cabify-code-challenge/internal/platform/server/handler/dropoff"
	"cabify-code-challenge/internal/platform/server/handler/locate"
	"cabify-code-challenge/internal/platform/server/handler/post_journey"
	"cabify-code-challenge/internal/platform/server/handler/status"
	"cabify-code-challenge/internal/platform/server/middleware/logging"
	"cabify-code-challenge/kit/command"
	"cabify-code-challenge/kit/event"
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
	eventBus   event.Bus
	carPool    *carpool.CarPool
}

func CarPoolMiddleware(carPool *carpool.CarPool) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("carPool", carPool)
		c.Next()
	}
}

func New(host string, port uint, commandBus command.Bus, queryBus query.Bus, eventBus event.Bus, carPool *carpool.CarPool) Server {
	srv := Server{
		httpAddr: fmt.Sprintf("%s:%d", host, port),
		engine:   gin.Default(),

		//deps
		commandBus: commandBus,
		queryBus:   queryBus,
		eventBus:   eventBus,
		carPool:    carPool,
	}

	srv.registerRoutes()
	srv.engine.HandleMethodNotAllowed = true
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {

	s.engine.Use(CarPoolMiddleware(s.carPool), logging.Middleware())

	s.engine.GET("/status", status.StatusHandler())
	s.engine.PUT("/cars", cars.PutCarsHandler(s.commandBus))
	s.engine.POST("/post_journey", post_journey.PostJourneyHandler(s.commandBus))
	s.engine.POST("/dropoff", dropoff.PostDropOffHandler(s.commandBus))
	s.engine.POST("/locate", locate.PostLocateHandler(s.queryBus))
}
