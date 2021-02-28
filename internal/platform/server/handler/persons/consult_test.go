package persons

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"xmen-mutant/kit/command/commandmocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_Consult_Service(t *testing.T) {
	commandBus := new(commandmocks.Bus)
	commandBus.On(
		"Dispatch",
		mock.Anything,
		mock.AnythingOfType("consulting.PersonCommand"),
	).Return(nil, nil)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/stats", ConsultHandler(commandBus))

	t.Run("given a valid request it returns 200", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/stats", nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusOK, res.StatusCode)
	})
}

func Example_Consult_ServiceSucess() {
	commandBus := new(commandmocks.Bus)
	commandBus.On(
		"Dispatch",
		mock.Anything,
		mock.AnythingOfType("consulting.PersonCommand"),
	).Return(nil, nil)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/stats", ConsultHandler(commandBus))
	req, _ := http.NewRequest(http.MethodGet, "/stats", nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)
}

func Benchmark_Consult_ServiceSucess(b *testing.B) {
	commandBus := new(commandmocks.Bus)
	commandBus.On(
		"Dispatch",
		mock.Anything,
		mock.AnythingOfType("consulting.PersonCommand"),
	).Return(nil, nil)

	for i := 0; i < b.N; i++ {
		gin.SetMode(gin.TestMode)
		r := gin.New()
		r.GET("/stats", ConsultHandler(commandBus))
		req, _ := http.NewRequest(http.MethodGet, "/stats", nil)
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)
	}
}
