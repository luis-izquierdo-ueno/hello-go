package mysql

import (
	"context"
	"database/sql"
	"fmt"
	core "hello-go/internal"

	"github.com/huandu/go-sqlbuilder"
)




type CourseRepository struct {
	db *sql.DB
}

func NewCourseRepository(db *sql.DB) *CourseRepository {
	return &CourseRepository{
		db: db,
	}
}

func (r *CourseRepository) Save(ctx context.Context, course core.Course) error {
	courseSQLStruct := sqlbuilder.NewStruct(new(sqlCourse))
	query, args := courseSQLStruct.InsertInto(sqlCourseTable, sqlCourse{
		ID: course.ID(),
		Name: course.Name(),
		Duration: course.Duration(),
	}).Build()

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to save course: %w", err)
	}

	return nil
}