package presentation

import (
	"cabify-code-challenge/kit/command/commandmocks"
	"cabify-code-challenge/kit/query/querymocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

func TestPostDropOffHandler(t *testing.T) {
	commandBus := new(commandmocks.Bus)
	commandBus.On("Dispatch", mock.Anything, mock.Anything).Return(nil)
	queryBus := new(querymocks.Bus)
	queryBus.On("Ask", mock.Anything, mock.Anything).Return(nil, nil)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/dropoff", PostDropOffHandler(commandBus, queryBus))

	t.Run("given a valid request it returns 200", func(t *testing.T) {
		dropOffRequest := postDropOffRequest{
			ID: 1,
		}
		formData := url.Values{}
		formData.Set("ID", strconv.Itoa(dropOffRequest.ID))

		request, err := http.NewRequest(http.MethodPost, "/dropoff", strings.NewReader(formData.Encode()))
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		require.NoError(t, err)

		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, request)

		response := recorder.Result()
		defer response.Body.Close()

		assert.Equal(t, http.StatusOK, response.StatusCode)
		mock.AssertExpectationsForObjects(t, commandBus)
	})

	t.Run("given an invalid body request it returns 400", func(t *testing.T) {
		dropOffRequest := postDropOffRequest{}
		formData := url.Values{}
		formData.Set("ID", strconv.Itoa(dropOffRequest.ID))

		request, err := http.NewRequest(http.MethodPost, "/dropoff", strings.NewReader(formData.Encode()))
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		require.NoError(t, err)

		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, request)

		response := recorder.Result()
		defer response.Body.Close()

		assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	})

	t.Run("given an invalid content type it returns 400", func(t *testing.T) {
		dropOffRequest := postDropOffRequest{
			ID: 1,
		}
		formData := url.Values{}
		formData.Set("ID", strconv.Itoa(dropOffRequest.ID))

		request, err := http.NewRequest(http.MethodPost, "/dropoff", strings.NewReader(formData.Encode()))
		request.Header.Set("Content-Type", "application/json")
		require.NoError(t, err)

		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, request)

		response := recorder.Result()
		defer response.Body.Close()

		assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	})

	t.Run("given an empty body request it returns 400", func(t *testing.T) {
		var dropOffRequest postDropOffRequest

		formData := url.Values{}
		formData.Set("ID", strconv.Itoa(dropOffRequest.ID))

		request, err := http.NewRequest(http.MethodPost, "/dropoff", strings.NewReader(formData.Encode()))
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		require.NoError(t, err)

		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, request)

		response := recorder.Result()
		defer response.Body.Close()

		assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	})
}
