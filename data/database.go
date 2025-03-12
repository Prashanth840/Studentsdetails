package data

import (
	"fmt"
	"studentsdetails/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func DbConnect() {
	dsn := "root:@Pachi840@tcp(127.0.0.1:3306)/prashanth?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())

	}
	db.AutoMigrate(&models.Students{}, &models.Courses{})
	fmt.Println("Database Connected Successfully")
	Db = db
}
