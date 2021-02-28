package persons

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"xmen-mutant/kit/command/commandmocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_Create_ServiceError(t *testing.T) {
	commandBus := new(commandmocks.Bus)
	commandBus.On(
		"Dispatch",
		mock.Anything,
		mock.AnythingOfType("creating.PersonCommand"),
	).Return(nil, nil)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/mutant", CreateHandler(commandBus))

	t.Run("given an invalid request it returns 400", func(t *testing.T) {
		createPersonReq := createRequest{
			Dna: []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"},
		}

		b, err := json.Marshal(createPersonReq)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/mutant", bytes.NewBuffer(b))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("given a valid request it returns 201", func(t *testing.T) {
		createPersonReq := createRequest{
			Dna: []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"},
		}

		b, err := json.Marshal(createPersonReq)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/mutant", bytes.NewBuffer(b))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})
}
