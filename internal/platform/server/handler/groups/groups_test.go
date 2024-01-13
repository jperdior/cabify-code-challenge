package groups

import (
	"bytes"
	"cabify-code-challenge/kit/command/commandmocks"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostJourneyHandler(t *testing.T) {
	commandBus := new(commandmocks.Bus)
	commandBus.On("Dispatch", mock.Anything, mock.Anything).Return(nil)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/journey", PostJourneyHandler(commandBus))

	t.Run("given a valid request it returns 200", func(t *testing.T) {
		journeyRequest := postJourneyRequest{
			ID:     1,
			People: 3,
		}
		body, err := json.Marshal(journeyRequest)
		require.NoError(t, err)

		request, err := http.NewRequest(http.MethodPost, "/journey", bytes.NewBuffer(body))
		request.Header.Set("Content-Type", "application/json")
		require.NoError(t, err)

		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, request)

		response := recorder.Result()
		defer response.Body.Close()

		assert.Equal(t, http.StatusOK, response.StatusCode)
		mock.AssertExpectationsForObjects(t, commandBus)
	})

	t.Run("given an invalid body request it returns 400", func(t *testing.T) {
		journeyRequest := postJourneyRequest{
			ID: 1,
		}
		body, err := json.Marshal(journeyRequest)
		require.NoError(t, err)

		request, err := http.NewRequest(http.MethodPost, "/journey", bytes.NewBuffer(body))
		request.Header.Set("Content-Type", "application/json")
		require.NoError(t, err)

		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, request)

		response := recorder.Result()
		defer response.Body.Close()

		assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	})

	t.Run("given an invalid content type it returns 400", func(t *testing.T) {
		journeyRequest := postJourneyRequest{
			ID: 1,
		}
		body, err := json.Marshal(journeyRequest)
		require.NoError(t, err)

		request, err := http.NewRequest(http.MethodPost, "/journey", bytes.NewBuffer(body))
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		require.NoError(t, err)

		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, request)

		response := recorder.Result()
		defer response.Body.Close()

		assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	})

	t.Run("given an empty body request it returns 400", func(t *testing.T) {
		var journeyRequest postJourneyRequest

		body, err := json.Marshal(journeyRequest)
		require.NoError(t, err)

		request, err := http.NewRequest(http.MethodPost, "/journey", bytes.NewBuffer(body))
		request.Header.Set("Content-Type", "application/json")
		require.NoError(t, err)

		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, request)

		response := recorder.Result()
		defer response.Body.Close()

		assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	})
}
