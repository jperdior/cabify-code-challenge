package presentation

import (
	"cabify-code-challenge/internal/carpool/application/post_journey"
	"cabify-code-challenge/internal/carpool/domain"
	"cabify-code-challenge/kit/command"
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

		err := commandBus.Dispatch(post_journey.NewCreatingJourneyCommand(request.ID, request.People))
		if err != nil {
			switch {
			case errors.Is(err, domain.ErrInvalidGroupID), errors.Is(err, domain.ErrInvalidPeople):
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
