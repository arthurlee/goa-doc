package models

import (
	"database/sql"
	"github.com/arthurlee/goa/database"
)

type Course struct {
	Id      int    `json:"courseId"`
	Name    string `json:"courseName"`
	Summary string `json:"courseSummary"`
}

func (me *Course) GetInserter() database.DbInserter {
	inserter := database.DbInserter{}

	inserter.SetTableName("course")
	inserter.
		AddStringField("course_name", me.Name).
		AddStringField("course_summary", me.Summary)
	inserter.Done()

	return inserter
}

type CourseList struct {
	Courses []Course
}

func NewCourseList() *CourseList {
	return &CourseList{make([]Course, 0, 10)}
}

func (this CourseList) GetSql() string {
	return "SELECT id, course_name, course_summary from course"
}

func (this CourseList) GetArgs() []interface{} {
	return nil
}

func (this *CourseList) SetItem(rows *sql.Rows) error {
	course := Course{}
	err := rows.Scan(&course.Id, &course.Name, &course.Summary)
	if err == nil {
		this.Courses = append(this.Courses, course)
	}
	return err
}
