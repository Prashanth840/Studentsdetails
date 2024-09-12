package repository

import (
	"errors"
	"studentsdetails/data"
	"studentsdetails/models"
)

func Addstudentdetails(input models.Students) (string, error) {
	tx, err := data.Db.Begin()
	if err != nil {
		tx.Rollback()
		return "", err
	}
	str := `insert into student(name,phone,email) values(?,?,?)`
	last, err := data.Db.Exec(str, input.Student_name, input.Phone, input.Email)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	stu, _ := last.LastInsertId()
	str1 := `insert into courses(name,student_id) values(?,?)`
	for _, val := range input.Coursedetails {
		if _, err1 := data.Db.Exec(str1, val.Course_name, stu); err1 != nil {
			tx.Rollback()
			return "", err1
		}
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return "", err
	}
	return "success", nil
}

func GetStudentdetails(input int) (models.Students, error) {
	var res models.Students
	str := `select id,name,phone,email from student where id=?`
	if err := data.Db.QueryRow(str, input).Scan(&res.Student_id, &res.Student_name, &res.Phone, &res.Email); err != nil  {
		return res, err
	}
	var res1 []models.Courses
	str1 := `select name from courses where student_id=?`
	rows, err := data.Db.Query(str1, input)
	if err != nil {
		return res, err
	}
	defer rows.Close()
	for rows.Next() {
		var val models.Courses
		if err := rows.Scan(&val.Course_name); err != nil {
			return res, err
		}
		res1 = append(res1, val)
	}
	res.Coursedetails = res1
	return res, nil
}

func Getallstudents() ([]models.Students, error) {
	var res []models.Students
	str := `select id,name,phone,email from student `
	str1 := `select name from courses where student_id=? `
	rows, err := data.Db.Query(str)
	if err != nil {
		return res, err
	}
	defer rows.Close()
	for rows.Next() {
		var content models.Students
		if err := rows.Scan(&content.Student_id, &content.Student_name, &content.Phone, &content.Email); err != nil {
			return res, err
		}
		rows1, err := data.Db.Query(str1, content.Student_id)
		if err != nil {
			return res, err
		}
		defer rows1.Close()
		var vcont []models.Courses
		for rows1.Next() {
			var con models.Courses
			if err := rows1.Scan(&con.Course_name); err != nil {
				return res, err
			}
			vcont = append(vcont, con)
		}

		content.Coursedetails = vcont
		res = append(res, content)
	}
	return res, nil
}

func Updatestudents(id int, input models.Students) (string, error) {
	str := `update student
	set name=? , email=?
	where id=?`
	res, err := data.Db.Exec(str, input.Student_name, input.Email, id)
	if err != nil {
		return "", err
	}
	val, _ := res.RowsAffected()
	if val == 0 {
		return "", errors.New("student not found")
	}
	return "Success", nil
}

func Deletestudent(id int) (string, error) {
	str1 := `delete from student where id=?`
	str2 := `delete from courses where student_id=?`
	tx, err := data.Db.Begin()
	if err != nil {
		tx.Rollback()
		return "", err
	}
	if _, err = tx.Exec(str2, id); err != nil {
		tx.Rollback()
		return "", err
	}
	res, err := tx.Exec(str1, id)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	val, _ := res.RowsAffected()
	if val == 0 {
		return "", errors.New("student not found")
	}
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return "", err
	}
	return "success", nil
}
