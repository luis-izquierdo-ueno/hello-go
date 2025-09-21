package creating

import (
	"context"
	"errors"
	core "hello-go/internal"
	"hello-go/internal/platform/storage/storagemocks"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_CourseService_CreateCourse_RepositoryError(test *testing.T) {
	courseID, courseName, courseDuration := uuid.New().String(), "Go Course", "40h"

	course, err := core.NewCourse(courseID, courseName, courseDuration)
	require.NoError(test, err)

	courseRepositoryMock := new(storagemocks.CourseRepository)
	courseRepositoryMock.On("Save", mock.Anything, course).Return(errors.New("repository error"))

	courseService := NewCourseService(courseRepositoryMock)

	err = courseService.CreateCourse(context.Background(), courseID, courseName, courseDuration)

	courseRepositoryMock.AssertExpectations(test)
	assert.Error(test, err)
}

func Test_CourseService_CreateCourse_Success(test *testing.T) {
	courseID, courseName, courseDuration := uuid.New().String(), "Go Course", "40h"

	course, err := core.NewCourse(courseID, courseName, courseDuration)
	require.NoError(test, err)

	courseRepositoryMock := new(storagemocks.CourseRepository)
	courseRepositoryMock.On("Save", mock.Anything, course).Return(nil)

	courseService := NewCourseService(courseRepositoryMock)

	err = courseService.CreateCourse(context.Background(), courseID, courseName, courseDuration)

	courseRepositoryMock.AssertExpectations(test)
	assert.NoError(test, err)
}
