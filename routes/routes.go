package routes

import (
	"studentsdetails/controller"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.POST("/api/students", controller.Addstudentdetails)
	r.GET("/api/students/:id", controller.GetStudentsdetails)
	r.GET("/api/students", controller.Addstudentdetails)
	r.PUT("/api/students/:id", controller.GetStudentsdetails)
	r.DELETE("/api/students/:id", controller.GetStudentsdetails)
}
