package models

import (
	"database/sql"
)

type Course struct {
	Id      int    `json:"courseId"`
	Name    string `json:"courseName"`
	Summary string `json:"courseSummary"`
}

type CourseList struct {
	Courses []Course
}

func NewCourseList() *CourseList {
	return &CourseList{make([]Course, 0, 10)}
}

func (this *CourseList) Sql() string {
	return "SELECT id, course_name, course_summary from course"
}

func (this *CourseList) SetItem(rows *sql.Rows) error {
	course := Course{}
	err := rows.Scan(&course.Id, &course.Name, &course.Summary)
	if err == nil {
		this.Courses = append(this.Courses, course)
	}
	return err
}
