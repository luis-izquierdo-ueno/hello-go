package recovery

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Middleware() gin.HandlerFunc { // Middleware para recuperar de panic y no dejar que se detenga el servidor
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("[Middleware] recovered from panic: %v in date %s", r, time.Now().Format(time.RFC3339))	
				
				c.Abort()
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()

		c.Next()
	}
}
