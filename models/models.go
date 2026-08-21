package models

type Students struct {
	Id            int       `gorm:"column:id;primaryKey;autoIncrement" json:"id,omitempty"`
	Name          string    `gorm:"column:name" json:"student_name" binding:"required"`
	Phone         string    `gorm:"column:phone" json:"phone,omitempty" binding:"required"`
	Email         string    `gorm:"column:email" json:"email,omitempty" binding:"required,email"`
	Coursedetails []Courses `gorm:"foreignKey:Student_id" json:"course_details,omitempty"`
}

type Courses struct {
	Id         int    `gorm:"column:id;primaryKey;autoIncrement" json:"id,omitempty"`
	Name       string `gorm:"column:name" json:"course_name" binding:"required"`
	Student_id int    `gorm:"column:student_id" json:"student_id,omitempty"`
}
