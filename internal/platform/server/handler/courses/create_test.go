package courses

import (
	"bytes"
	"encoding/json"
	"hello-go/internal/creating"
	"hello-go/internal/platform/storage/storagemocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_Create(test *testing.T) {
	
	courseRepository:= new(storagemocks.CourseRepository)
	creatingCourseService := creating.NewCourseService(courseRepository)
	courseRepository.On("Save", mock.Anything, mock.AnythingOfType("core.Course")).Return(nil)

	gin.SetMode(gin.TestMode)
	engine := gin.New()
	engine.POST("/courses", CreateHandler(*creatingCourseService))
	
	test.Run("given and invalid request, should return a 400 status code", func(test *testing.T) {
		createCourseRequest := createRequest{
			ID: "123", 
			Name: "Go Course", 
		}

		body, err := json.Marshal(createCourseRequest)
		require.NoError(test, err)

		request, err := http.NewRequest(http.MethodPost, "/courses", bytes.NewBuffer(body))
		require.NoError(test, err)

		recorder := httptest.NewRecorder()
		engine.ServeHTTP(recorder, request)
		

		response := recorder.Result()
		defer response.Body.Close()

		assert.Equal(test, http.StatusBadRequest, response.StatusCode)
	})

	test.Run("given a valid request, should return a 201 status code", func(test *testing.T) {
		createCourseRequest := createRequest{
			ID: uuid.New().String(), 
			Name: "Go Course", 
			Duration: "40h",
		}

		body, err := json.Marshal(createCourseRequest)
		require.NoError(test, err)

		request, err := http.NewRequest(http.MethodPost, "/courses", bytes.NewBuffer(body))
		require.NoError(test, err)

		recorder := httptest.NewRecorder()
		engine.ServeHTTP(recorder, request)
		

		response := recorder.Result()
		defer response.Body.Close()

		assert.Equal(test, http.StatusCreated, response.StatusCode)
	})

}