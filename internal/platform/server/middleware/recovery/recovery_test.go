package recovery

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRecoveryMiddleware(test *testing.T) {

	gin.SetMode(gin.TestMode)
	engine := gin.New()
	// engine := gin.Default() // es igual a gin.New() pero con el middleware de logger y recovery por defecto
	engine.Use(Middleware())
	engine.GET("/test-middleware", func(c *gin.Context) {
		panic("test panic")
	})

	httpRecorder := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/test-middleware", nil)
	require.NoError(test, err)

	assert.NotPanics(test, func() {
		engine.ServeHTTP(httpRecorder, request)
	})



}