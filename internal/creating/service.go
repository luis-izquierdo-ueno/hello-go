package creating

import (
	"context"
	core "hello-go/internal"
)

type CourseService struct {
	courseRepository core.CourseRepository
}

func NewCourseService(courseRepository core.CourseRepository) *CourseService {
	return &CourseService{
		courseRepository: courseRepository,
	}
}

func (s *CourseService) CreateCourse(ctx context.Context, id, name, duration string) error {
	course, err := core.NewCourse(id, name, duration)
	if err != nil {
		return err
	}

	return s.courseRepository.Save(ctx, course)
}
