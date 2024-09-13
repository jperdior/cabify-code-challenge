package presentation

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func StatusHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.String(http.StatusOK, "OK")
	}
}
