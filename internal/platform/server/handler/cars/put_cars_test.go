package cars

import (
	"bytes"
	"cabify-code-challenge/kit/command/commandmocks"
	"encoding/json"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"

	//"github.com/CodelyTV/go-hexagonal_http_api-course/02-03-controller-test/internal/platform/storage/storagemocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPutCarsHandler(t *testing.T) {

	commandBus := new(commandmocks.Bus)
	commandBus.On(
		"Dispatch",
		mock.Anything,
		mock.Anything,
	).Return(nil)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.PUT("/cars", PutCarsHandler(
		commandBus,
	))

	t.Run("given an valid request it returns 200", func(t *testing.T) {
		putCarsRequests := []putCarsRequest{
			{ID: 2, Seats: 3},
		}

		body, err := json.Marshal(putCarsRequests)
		require.NoError(t, err)

		request, err := http.NewRequest(http.MethodPut, "/cars", bytes.NewBuffer(body))
		require.NoError(t, err)

		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, request)

		response := recorder.Result()
		defer response.Body.Close()

		assert.Equal(t, http.StatusOK, response.StatusCode)
	})

	t.Run("given an invalid request it returns 400", func(t *testing.T) {
		putCarsRequests := []putCarsRequest{
			{ID: 2},
		}

		body, err := json.Marshal(putCarsRequests)
		require.NoError(t, err)

		request, err := http.NewRequest(http.MethodPut, "/cars", bytes.NewBuffer(body))
		require.NoError(t, err)

		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, request)

		response := recorder.Result()
		defer response.Body.Close()

		assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	})
}
