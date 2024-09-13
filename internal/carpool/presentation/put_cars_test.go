package presentation

import (
	"bytes"
	"cabify-code-challenge/kit/command/commandmocks"

	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
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

	t.Run("given a valid request it returns 200", func(t *testing.T) {
		putCarsRequests := []putCarsRequest{
			{ID: 2, Seats: 4},
		}

		body, err := json.Marshal(putCarsRequests)
		require.NoError(t, err)

		request, err := http.NewRequest(http.MethodPut, "/cars", bytes.NewBuffer(body))
		request.Header.Set("Content-Type", "application/json")
		require.NoError(t, err)

		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, request)

		response := recorder.Result()
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			require.NoError(t, err)
		}(response.Body)

		assert.Equal(t, http.StatusOK, response.StatusCode)
		mock.AssertExpectationsForObjects(t, commandBus)
	})

	t.Run("given an invalid body request it returns 400", func(t *testing.T) {
		putCarsRequests := []putCarsRequest{
			{ID: 2},
		}

		body, err := json.Marshal(putCarsRequests)
		require.NoError(t, err)

		request, err := http.NewRequest(http.MethodPut, "/cars", bytes.NewBuffer(body))
		request.Header.Set("Content-Type", "application/json")
		require.NoError(t, err)

		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, request)

		response := recorder.Result()
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			require.NoError(t, err)
		}(response.Body)

		assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	})

	t.Run("given an empty body request it returns 400", func(t *testing.T) {
		var putCarsRequests []putCarsRequest

		body, err := json.Marshal(putCarsRequests)
		require.NoError(t, err)

		request, err := http.NewRequest(http.MethodPut, "/cars", bytes.NewBuffer(body))
		request.Header.Set("Content-Type", "application/json")
		require.NoError(t, err)

		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, request)

		response := recorder.Result()
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			require.NoError(t, err)
		}(response.Body)

		assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	})

	t.Run("given an invalid request content type it returns 400", func(t *testing.T) {
		putCarsRequests := []putCarsRequest{
			{ID: 2, Seats: 4},
		}

		body, err := json.Marshal(putCarsRequests)
		require.NoError(t, err)

		request, err := http.NewRequest(http.MethodPut, "/cars", bytes.NewBuffer(body))
		request.Header.Set("Content-Type", "application/form-data")
		require.NoError(t, err)

		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, request)

		response := recorder.Result()
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			require.NoError(t, err)
		}(response.Body)

		assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	})

}
