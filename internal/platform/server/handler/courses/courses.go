package courses

import (
	"errors"
	core "hello-go/internal"
	"hello-go/internal/creating"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createRequest struct {
	ID       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}

// CreateHandler is a handler for the create course endpoint
func CreateHandler(creatingCourseService creating.CourseService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request createRequest
		if err := c.BindJSON(&request); err != nil { // Se modifica la referencia de request para que se pueda modificar el struct
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		err := creatingCourseService.CreateCourse(c, request.ID, request.Name, request.Duration)
		if err != nil {
			switch {
			case errors.Is(err, core.ErrInvalidCourseID):
				c.JSON(http.StatusBadRequest, err.Error())
				return
			default:
				c.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		}

		c.Status(http.StatusCreated)
	}
}
