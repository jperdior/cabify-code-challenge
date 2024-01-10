package cars

import (
	"cabify-code-challenge/internal/carpool"
	"cabify-code-challenge/internal/use_cases/put_cars"
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
		if err := context.BindJSON(&request); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var cars []put_cars.CarDTO
		for _, car := range request {
			newCar := put_cars.NewCarDTO(car.ID, car.Seats)
			cars = append(cars, newCar)
		}
		err := commandBus.Dispatch(context, put_cars.NewPutCarsCommand(cars))
		if err != nil {
			switch {
			case errors.Is(err, carpool.ErrInvalidCarID), errors.Is(err, carpool.ErrInvalidSeats):
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
