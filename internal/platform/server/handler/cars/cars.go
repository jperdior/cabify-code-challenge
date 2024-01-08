package cars

import (
	"cabify-code-challenge/internal/use_cases/putting_cars"
	"cabify-code-challenge/kit/command"
	"github.com/gin-gonic/gin"
	"net/http"
)

type putCarsRequest struct {
	ID    int `json:"id" binding:"required"`
	Seats int `json:"seats" binding:"required"`
}

func PutCarsHandler(commandBus command.Bus) gin.HandlerFunc {
	return func(context *gin.Context) {
		var request []putCarsRequest
		if err := context.BindJSON(&request); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var cars []putting_cars.CarDTO
		for _, car := range request {
			newCar := putting_cars.NewCarDTO(car.ID, car.Seats)
			cars = append(cars, newCar)
		}
		commandBus.Dispatch(context, putting_cars.NewPutCarsCommand(cars))

		context.Status(http.StatusOK)
	}
}
