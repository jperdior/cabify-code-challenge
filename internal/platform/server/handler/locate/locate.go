package locate

import (
	"cabify-code-challenge/internal/carpool"
	"cabify-code-challenge/internal/use_cases/locate"
	"cabify-code-challenge/kit/query"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type postLocateRequest struct {
	ID int `form:"ID" binding:"required"`
}

func PostLocateHandler(queryBus query.Bus) gin.HandlerFunc {
	return func(context *gin.Context) {
		var request postLocateRequest
		if err := context.ShouldBind(&request); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		car, err := queryBus.Ask(context, locate.NewLocateQuery(request.ID))
		if err != nil {
			switch {
			case errors.Is(err, carpool.ErrInvalidGroupID):
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			case errors.Is(err, carpool.ErrGroupNotFound):
				context.String(http.StatusNotFound, "")
				return
			default:
				context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}
		if locResponse, ok := car.(locate.LocationResponse); ok && locResponse.Id == 0 && locResponse.Seats == 0 {
			context.Status(http.StatusNoContent)
			return
		}
		context.JSON(http.StatusOK, car)
		return
	}
}
