package groups

import (
	"cabify-code-challenge/internal/carpool"
	"cabify-code-challenge/internal/use_cases/dropoff"
	"cabify-code-challenge/internal/use_cases/journey"
	"cabify-code-challenge/internal/use_cases/locate"
	"cabify-code-challenge/kit/command"
	"cabify-code-challenge/kit/query"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type postJourneyRequest struct {
	ID     int `json:"id" binding:"required"`
	People int `json:"people" binding:"required"`
}

func PostJourneyHandler(commandBus command.Bus) gin.HandlerFunc {
	return func(context *gin.Context) {
		contentType := context.Request.Header.Get("Content-Type")
		if contentType != "application/json" {
			context.JSON(http.StatusBadRequest, gin.H{"error": "invalid content type"})
			return
		}
		var request postJourneyRequest
		if err := context.ShouldBindJSON(&request); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := commandBus.Dispatch(context, journey.NewCreatingJourneyCommand(request.ID, request.People))
		if err != nil {
			switch {
			case errors.Is(err, carpool.ErrInvalidGroupID), errors.Is(err, carpool.ErrInvalidPeople):
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

type postDropOffRequest struct {
	ID int `form:"ID" binding:"required"`
}

func PostDropOffHandler(commandBus command.Bus) gin.HandlerFunc {
	return func(context *gin.Context) {
		contentType := context.Request.Header.Get("Content-Type")
		if contentType != "application/x-www-form-urlencoded" {
			context.JSON(http.StatusBadRequest, gin.H{"error": "invalid content type"})
			return
		}
		var request postDropOffRequest
		if err := context.ShouldBind(&request); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := commandBus.Dispatch(context, dropoff.NewDropOffCommand(request.ID))
		if err != nil {
			switch {
			case errors.Is(err, carpool.ErrInvalidGroupID):
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			case errors.Is(err, carpool.ErrGroupNotFound):
				context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			default:
				context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}

		context.Status(http.StatusOK)
	}
}

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
