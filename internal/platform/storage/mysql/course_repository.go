package mysql

import (
	"context"
	"database/sql"
	"fmt"
	core "hello-go/internal"
	"time"

	"github.com/huandu/go-sqlbuilder"
)

// CourseRepository is a repository for courses MySQL implementation of core.CourseRepository
type CourseRepository struct {
	db *sql.DB
	dbTimeout time.Duration
}

func NewCourseRepository(db *sql.DB, dbTimeout time.Duration) *CourseRepository {
	return &CourseRepository{
		db: db,
		dbTimeout: dbTimeout,
	}
}

func (repo *CourseRepository) Save(ctx context.Context, course core.Course) error {
	courseSQLStruct := sqlbuilder.NewStruct(new(sqlCourse))
	query, args := courseSQLStruct.InsertInto(sqlCourseTable, sqlCourse{
		ID: course.ID().String(),
		Name: course.Name(),
		Duration: course.Duration(),
	}).Build()

	ctxTimeout,cancel := context.WithTimeout(ctx, repo.dbTimeout)
	defer cancel()

	_, err := repo.db.ExecContext(ctxTimeout, query, args...)
	if err != nil {
		return fmt.Errorf("failed to save course: %v", err)
	}

	return nil
}