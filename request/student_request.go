package request

type CreateStudentRequest struct {
	Name  string `json:"student_name" binding:"required,max=255"`
	Phone string `json:"phone" binding:"required,min=7,max=20"`
	Email string `json:"email" binding:"required,email"`
}

type UpdateStudentRequest struct {
	Name  string `json:"student_name" binding:"required,max=255"`
	Phone string `json:"phone" binding:"required,min=7,max=20"`
	Email string `json:"email" binding:"required,email"`
}
