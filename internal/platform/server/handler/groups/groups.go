package groups

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type postJourneyRequest struct {
	ID     int `json:"id" binding:"required"`
	People int `json:"people" binding:"required"`
}

func PostJourneyHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		var request postJourneyRequest
		if err := context.BindJSON(&request); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//var group = carpooling.NewGroup(request.ID, request.People)
		//
		//fmt.Printf("Group %d with %d people\n", group.ID, group.People)

		context.Status(http.StatusOK)
	}
}

type postDropOffRequest struct {
	ID int `form:"ID" binding:"required"`
}

func PostDropOffHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		var request postDropOffRequest
		if err := context.ShouldBind(&request); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Printf("Drop Group %d\n", request.ID)

		context.Status(http.StatusOK)
	}
}

type postLocateRequest struct {
	ID int `form:"ID" binding:"required"`
}

func PostLocateHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		var request postLocateRequest
		if err := context.ShouldBind(&request); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Printf("Locate Group %d\n", request.ID)

		context.Status(http.StatusOK)
	}
}
