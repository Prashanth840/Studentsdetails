package models

type Input struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Students struct {
	Student_id    int       `json:"id,omitempty"`
	Student_name  string    `json:"student_name"`
	Phone         string    `json:"phone,omitempty"`
	Email         string    `json:"email,omitempty"`
	Coursedetails []Courses `json:"course_details,omitempty"`
}

type Courses struct {
	Course_name string `json:"course_name"`
}
