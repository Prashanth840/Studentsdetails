package routes

import (
	"studentsdetails/controller"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.POST("/api/students", controller.Createstudentdetails)
	r.GET("/api/students/:id", controller.GetStudentsdetails)
	r.GET("/api/students", controller.Getallstudents)
	r.PUT("/api/students/:id", controller.Updatestudentdetails)
	r.DELETE("/api/students/:id", controller.Deletestudentdetails)
	r.POST("/api/courses", controller.Createcoursedetails)
	r.PUT("/api/courses/:id", controller.Updatecoursedetails)
	r.DELETE("/api/courses/:id", controller.Deletecoursedetails)
}
