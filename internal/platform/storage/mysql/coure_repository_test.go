package mysql

import (
	"context"
	"errors"
	core "hello-go/internal"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_CourseRepository_Save_RepositoryError(test *testing.T) {
	courseID, courseName, courseDuration := uuid.New().String(), "Go Course", "40h"
	course, err := core.NewCourse(courseID, courseName, courseDuration)
	require.NoError(test, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(test, err)

	sqlMock.ExpectExec("INSERT INTO courses (id, name, duration) VALUES (?, ?, ?)").
		WithArgs(courseID, courseName, courseDuration).
		WillReturnError(errors.New("repository error"))

	repo := NewCourseRepository(db, 10*time.Second)

	repoErr := repo.Save(context.Background(), course)

	assert.NoError(test, sqlMock.ExpectationsWereMet())
	assert.Error(test, repoErr)
}

func Test_CourseRepository_Save_Success(test *testing.T) {
	courseID, courseName, courseDuration := uuid.New().String(), "Go Course", "40h"
	course, err := core.NewCourse(courseID, courseName, courseDuration)
	require.NoError(test, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(test, err)

	sqlMock.ExpectExec("INSERT INTO courses (id, name, duration) VALUES (?, ?, ?)").
		WithArgs(courseID, courseName, courseDuration).
		WillReturnResult(sqlmock.NewResult(0,1))

	repo := NewCourseRepository(db, 10*time.Second)

	repoErr := repo.Save(context.Background(), course)

	assert.NoError(test, sqlMock.ExpectationsWereMet())
	assert.NoError(test, repoErr)
}