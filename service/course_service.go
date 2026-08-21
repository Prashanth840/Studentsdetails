package service

import (
	"errors"
	"studentsdetails/models"
	"studentsdetails/repository"
	"studentsdetails/request"
)

func CreateCourse(input request.CreateCourseRequest) (string, error) {
	if input.StudentID <= 0 {
		return "", errors.New("student_id must be a positive integer")
	}
	course := models.Courses{
		Name:       input.Name,
		Student_id: input.StudentID,
	}
	return repository.CreateCourseDetails(course)
}

func UpdateCourse(id int, input request.UpdateCourseRequest) (int64, error) {
	if id <= 0 {
		return 0, errors.New("id must be a positive integer")
	}
	course := models.Courses{
		Name:       input.Name,
		Student_id: input.StudentID,
	}
	return repository.UpdateCourse(id, course)
}

func DeleteCourse(id int) (int64, error) {
	if id <= 0 {
		return 0, errors.New("id must be a positive integer")
	}
	return repository.DeleteCourse(id)
}
