package repository

import (
	"errors"
	"studentsdetails/data"
	"studentsdetails/models"
)

func CreateStudentDetails(input models.Students) (string, error) {
	if err := data.Db.Create(&input).Error; err != nil {
		return "", errors.New("failed to add student details")
	}
	return "Student details added successfully", nil
}

func CreateCourseDetails(input models.Courses) (string, error) {
	if err := data.Db.Create(&input).Error; err != nil {
		return "", errors.New("failed to add course details")
	}
	return "Course details added successfully", nil
}

func UpdateCourse(id int, input models.Courses) (int64, error) {
	result := data.Db.Model(&models.Courses{}).Where("id = ?", id).Updates(input)
	return result.RowsAffected, result.Error
}

func DeleteCourse(id int) (int64, error) {
	result := data.Db.Delete(&models.Courses{}, id)
	return result.RowsAffected, result.Error
}

func GetStudentDetails(input int) (models.Students, error) {
	var res models.Students
	var students models.Students
	result := data.Db.Preload("Coursedetails").Find(&students, input)
	if result.Error != nil {
		return res, result.Error
	}
	return students, nil
}

func GetAllStudents() ([]models.Students, error) {
	var students []models.Students
	result := data.Db.Preload("Coursedetails").Find(&students)
	if result.Error != nil {
		return nil, result.Error
	}
	return students, nil
}

func UpdateStudent(id int, input models.Students) (int64, error) {
	result := data.Db.Model(&models.Students{}).Where("id = ?", id).Updates(input)
	return result.RowsAffected, result.Error
}

func DeleteStudent(id int) (int64, error) {
	data.Db.Where("student_id = ?", id).Delete(&models.Courses{})
	result := data.Db.Delete(&models.Students{}, id)
	return result.RowsAffected, result.Error
}
