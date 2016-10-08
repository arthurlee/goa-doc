package course

import (
	"encoding/json"
	"github.com/arthurlee/goa/database"
	"log"
	"net/http"
)

type GoaBaseRes struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (res *GoaBaseRes) setError(code string, message string) {
	res.Code = code
	res.Message = message
}

func (res *GoaBaseRes) setOK() {
	res.Code = "0"
	res.Message = "OK"
}

type Course struct {
	Id      int    `json:"courseId"`
	Name    string `json:"courseName"`
	Summary string `json:"courseSummary"`
}

type CourseListRes struct {
	GoaBaseRes
	Courses []Course `json:"course_list"`
}

func List(w http.ResponseWriter, r *http.Request) {
	resCourseList := CourseListRes{}
	resCourseList.setOK()

	rows, err := database.Db.Query("SELECT id, course_name, course_summary from course")
	if err == nil {
		columns, err := rows.Columns()
		if err == nil {
			log.Println(columns)
		} else {
			log.Fatal(err)
		}

		resCourseList.Courses = make([]Course, 0, 10)

		for rows.Next() {
			course := Course{}
			err = rows.Scan(&course.Id, &course.Name, &course.Summary)
			if err != nil {
				log.Fatal(err)
				break
			}

			resCourseList.Courses = append(resCourseList.Courses, course)
		}
	} else {
		resCourseList.setError("1", err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resCourseList)
}
