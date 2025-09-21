package courses

import (
	core "hello-go/internal"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createRequest struct {
	ID string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}

// CreateHandler is a handler for the create course endpoint
func CreateHandler(courseRepository core.CourseRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request createRequest
		if err := c.BindJSON(&request); err != nil { // Se modifica la referencia de request para que se pueda modificar el struct
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		course, err := core.NewCourse(request.ID, request.Name, request.Duration)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		if err := courseRepository.Save(c,course); err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.Status(http.StatusCreated)
	}
}
