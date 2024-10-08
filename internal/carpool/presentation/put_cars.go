package presentation

import (
	"cabify-code-challenge/internal/carpool/application/put_cars"
	"cabify-code-challenge/internal/carpool/domain"
	"cabify-code-challenge/kit/command"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type putCarsRequest struct {
	ID    int `json:"id" binding:"required"`
	Seats int `json:"seats" binding:"required,min=4,max=6"`
}

func PutCarsHandler(commandBus command.Bus) gin.HandlerFunc {
	return func(context *gin.Context) {
		var request []putCarsRequest
		contentType := context.Request.Header.Get("Content-Type")
		if contentType != "application/json" {
			context.JSON(http.StatusBadRequest, gin.H{"error": "invalid content type"})
			return
		}
		if err := context.ShouldBindJSON(&request); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var cars []put_cars.CarDTO
		for _, car := range request {
			newCar := put_cars.NewCarDTO(car.ID, car.Seats)
			cars = append(cars, newCar)
		}
		if len(cars) == 0 {
			context.JSON(http.StatusBadRequest, gin.H{"error": "no cars provided"})
			return
		}
		err := commandBus.Dispatch(put_cars.NewPutCarsCommand(cars))
		if err != nil {
			switch {
			case errors.Is(err, domain.ErrInvalidCarID), errors.Is(err, domain.ErrInvalidSeats):
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			default:
				context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}

		context.Status(http.StatusOK)
	}
}
