package routes

import (
	"studentsdetails/controller"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine)  {
	r.POST("/students", controller.Addstudentdetails)
	r.GET("/students/:id", controller.GetStudentsdetails)
	r.GET("/students", controller.Addstudentdetails)
	r.PUT("/students/:id", controller.GetStudentsdetails)
	r.DELETE("//students/:id", controller.GetStudentsdetails)
}
