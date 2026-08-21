package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"studentsdetails/request"
	"studentsdetails/service"

	"github.com/gin-gonic/gin"
)

func parseID(c *gin.Context) (int, bool) {
	val, err := strconv.Atoi(c.Param("id"))
	if err != nil || val <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be a positive integer"})
		return 0, false
	}
	return val, true
}

func Createstudentdetails(c *gin.Context) {
	var input request.CreateStudentRequest
	if err := json.NewDecoder(c.Request.Body).Decode(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	msr, err := service.CreateStudent(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": msr})
}

func Createcoursedetails(c *gin.Context) {
	var input request.CreateCourseRequest
	if err := json.NewDecoder(c.Request.Body).Decode(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	msr, err := service.CreateCourse(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": msr})
}

func Updatecoursedetails(c *gin.Context) {
	val, ok := parseID(c)
	if !ok {
		return
	}
	var input request.UpdateCourseRequest
	if err := json.NewDecoder(c.Request.Body).Decode(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rows, err := service.UpdateCourse(val, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if rows == 0 {
		c.JSON(http.StatusAccepted, gin.H{"message": "Course not Found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Course updated successfully"})
}

func Deletecoursedetails(c *gin.Context) {
	val, ok := parseID(c)
	if !ok {
		return
	}
	rows, err := service.DeleteCourse(val)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if rows == 0 {
		c.JSON(http.StatusAccepted, gin.H{"message": "Course not Found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted Successfully"})
}

func Getallstudents(c *gin.Context) {
	data, err := service.GetAllStudents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}

func GetStudentsdetails(c *gin.Context) {
	val, ok := parseID(c)
	if !ok {
		return
	}

	data, err := service.GetStudent(val)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if data.Id == 0 {
		c.JSON(http.StatusAccepted, gin.H{"Message": "Student not Found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}

func Updatestudentdetails(c *gin.Context) {
	val, ok := parseID(c)
	if !ok {
		return
	}
	var input request.UpdateStudentRequest
	if err := json.NewDecoder(c.Request.Body).Decode(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rows, err := service.UpdateStudent(val, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if rows == 0 {
		c.JSON(http.StatusAccepted, gin.H{"message": "Student not Found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Student updated successfully"})
}

func Deletestudentdetails(c *gin.Context) {
	val, ok := parseID(c)
	if !ok {
		return
	}
	rows, err := service.DeleteStudent(val)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if rows == 0 {
		c.JSON(http.StatusAccepted, gin.H{"message": "Student not Found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted Successfully"})
}
