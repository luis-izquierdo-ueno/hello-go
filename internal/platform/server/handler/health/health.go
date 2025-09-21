package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CheckHandler is a handler for the health check endpoint
func CheckHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "Healthy!")
	}
}
