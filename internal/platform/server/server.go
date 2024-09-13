package server

import (
	"cabify-code-challenge/internal/carpool/domain"
	"cabify-code-challenge/internal/carpool/presentation"
	"cabify-code-challenge/internal/platform/server/middleware/logging"
	"cabify-code-challenge/kit/command"
	"cabify-code-challenge/kit/event"
	"cabify-code-challenge/kit/query"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	shutdownTimeout time.Duration

	//deps
	commandBus command.Bus
	queryBus   query.Bus
	eventBus   event.Bus
	carPool    *domain.CarPool
}

func CarPoolMiddleware(carPool *domain.CarPool) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("carPool", carPool)
		c.Next()
	}
}

func New(ctx context.Context, host string, port uint, shutdownTimeout time.Duration, commandBus command.Bus, queryBus query.Bus, eventBus event.Bus) (context.Context, Server) {
	srv := Server{
		httpAddr: fmt.Sprintf("%s:%d", host, port),
		engine:   gin.Default(),

		shutdownTimeout: shutdownTimeout,

		//deps
		commandBus: commandBus,
		queryBus:   queryBus,
		eventBus:   eventBus,
	}

	srv.registerRoutes()
	srv.engine.HandleMethodNotAllowed = true
	return serverContext(ctx), srv
}

func (s *Server) Run(ctx context.Context) error {
	log.Println("Server running on", s.httpAddr)

	srv := &http.Server{
		Addr:    s.httpAddr,
		Handler: s.engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server shut down", err)
		}
	}()

	<-ctx.Done()
	ctxShutDown, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return srv.Shutdown(ctxShutDown)
}

func (s *Server) registerRoutes() {

	s.engine.Use(CarPoolMiddleware(s.carPool), logging.Middleware())

	s.engine.GET("/status", presentation.StatusHandler())
	s.engine.PUT("/cars", presentation.PutCarsHandler(s.commandBus))
	s.engine.POST("/journey", presentation.PostJourneyHandler(s.commandBus))
	s.engine.POST("/dropoff", presentation.PostDropOffHandler(s.commandBus, s.queryBus))
	s.engine.POST("/locate", presentation.PostLocateHandler(s.queryBus))
}

func serverContext(ctx context.Context) context.Context {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		<-c
		cancel()
	}()

	return ctx
}
