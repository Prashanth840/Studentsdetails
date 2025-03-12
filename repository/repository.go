package repository

import (
	"errors"
	"studentsdetails/data"
	"studentsdetails/models"

	"gorm.io/gorm"
)

func Addstudentdetails(input models.Students) (string, error) {
	if err := data.Db.Create(&input).Error; err != nil {
		return "", errors.New("failed to add student details")
	}
	return "Student details added successfully", nil
}

func GetStudentdetails(input int) (models.Students, error) {
	var res models.Students
	var students models.Students
	result := data.Db.Preload("Coursedetails").Find(&students, input)
	if result.Error != nil {
		return res, result.Error
	}
	return students, nil
}

func Getallstudents() ([]models.Students, error) {
	var students []models.Students
	result := data.Db.Preload("Coursedetails").Find(&students)
	if result.Error != nil {
		return nil, result.Error
	}
	return students, nil
}

func Updatestudents(id int, input models.Students) *gorm.DB {
	result := data.Db.Model(&models.Students{}).Where("id = ?", id).Updates(input)
	return result
}

func Deletestudent(id int) *gorm.DB {
	data.Db.Where("student_id = ?", id).Delete(&models.Courses{})
	result := data.Db.Delete(&models.Students{}, id)
	return result
}
