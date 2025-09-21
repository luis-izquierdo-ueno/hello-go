package core

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

var ErrInvalidCourseID = errors.New("invalid course id")

type CourseID struct {
	value string
}

func NewCourseID(value string) (CourseID, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return CourseID{}, fmt.Errorf("%w: %v", ErrInvalidCourseID, value)
	}

	return CourseID{
		value: v.String(),
	}, nil
}

type CourseRepository interface {
	Save(ctx context.Context, course Course) error
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=CourseRepository

type Course struct {
	id CourseID 
	name string 
	duration string 
}

func NewCourse(id string, name string, duration string) (Course, error) {
	idVO, err := NewCourseID(id)
	if err != nil{
		return Course{}, err
	}

	return Course{
		id: idVO,
		name: name,
		duration: duration,
	}, nil
}

func (c Course) ID() string {
	return c.id.value
}

func (c Course) Name() string {
	return c.name
}

func (c Course) Duration() string {
	return c.duration
}


