package service

import (
	"errors"
	"studentsdetails/models"
	"studentsdetails/repository"
	"studentsdetails/request"
)

func CreateStudent(input request.CreateStudentRequest) (string, error) {
	student := models.Students{
		Name:  input.Name,
		Phone: input.Phone,
		Email: input.Email,
	}
	return repository.CreateStudentDetails(student)
}

func GetStudent(id int) (models.Students, error) {
	if id <= 0 {
		return models.Students{}, errors.New("id must be a positive integer")
	}
	return repository.GetStudentDetails(id)
}

func GetAllStudents() ([]models.Students, error) {
	return repository.GetAllStudents()
}

func UpdateStudent(id int, input request.UpdateStudentRequest) (int64, error) {
	if id <= 0 {
		return 0, errors.New("id must be a positive integer")
	}
	student := models.Students{
		Name:  input.Name,
		Phone: input.Phone,
		Email: input.Email,
	}
	return repository.UpdateStudent(id, student)
}

func DeleteStudent(id int) (int64, error) {
	if id <= 0 {
		return 0, errors.New("id must be a positive integer")
	}
	return repository.DeleteStudent(id)
}
