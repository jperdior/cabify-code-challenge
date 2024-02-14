package status

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatusHandler(t *testing.T) {

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/status", StatusHandler())

	t.Run("given a request it returns 200", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodGet, "/status", nil)
		require.NoError(t, err)

		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, request)

		response := recorder.Result()
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			require.New(t).NoError(err)
		}(response.Body)

		require.Equal(t, http.StatusOK, response.StatusCode)
	})
}
