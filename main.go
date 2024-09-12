package main

import (
	"studentsdetails/data"
	"studentsdetails/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	data.DbConnect()
	r := gin.Default()
	routes.Routes(r)
	r.Run(":9000")

}
