package request

type CreateCourseRequest struct {
	Name      string `json:"course_name" binding:"required,max=255"`
	StudentID int    `json:"student_id" binding:"required,min=1"`
}

type UpdateCourseRequest struct {
	Name      string `json:"course_name" binding:"required,max=255"`
	StudentID int    `json:"student_id" binding:"required,min=1"`
}
