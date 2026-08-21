package main

import (
	"os"
	"studentsdetails/data"
	"studentsdetails/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	data.DbConnect()
	r := gin.Default()
	routes.Routes(r)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "9001"
	}
	r.Run(":" + port)
}
