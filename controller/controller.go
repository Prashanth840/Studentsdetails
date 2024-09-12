package controller

import (
	"database/sql"
	"net/http"
	"strconv"
	"studentsdetails/models"
	"studentsdetails/repository"

	"github.com/gin-gonic/gin"
)

func Addstudentdetails(c *gin.Context) {
	if c.Request.Method == "POST" {
		var input models.Students
		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		msr, err := repository.Addstudentdetails(input)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": msr})
	} else {
		data, err := repository.Getallstudents()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}

func GetStudentsdetails(c *gin.Context) {
	id := c.Param("id")
	val, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if c.Request.Method == "GET" {

		data, err := repository.GetStudentdetails(val)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusAccepted, gin.H{"Message": "Student not Found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": data})
	} else if c.Request.Method == "PUT" {
		var input models.Students
		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		msg, err := repository.Updatestudents(val, input)
		if err != nil {
			if err.Error() == "student not found" {
				c.JSON(http.StatusAccepted, gin.H{"message": "Student not Found"})
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": msg})
	} else {
		msg, err := repository.Deletestudent(val)
		if err != nil {
			if err.Error() == "student not found" {
				c.JSON(http.StatusAccepted, gin.H{"message": "Student not Found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": msg})
	}
}
