package status

import (
	"cabify-code-challenge/internal/carpool"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func StatusHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		carPool := context.MustGet("carPool").(*carpool.CarPool)
		// marshall the carPool to json
		carsBySeats := carPool.GetCarsBySeat()
		for seats, cars := range carsBySeats {
			fmt.Printf("Cars with %d seats:\n", seats)
			for _, car := range cars {
				fmt.Printf("CarID %d with %d seats\n", car.ID(), car.Seats())
			}
		}
		context.String(http.StatusOK, "OK")
	}
}
