package course

import (
	"encoding/json"
	"github.com/arthurlee/goa-doc/samples/hello/models"
	"github.com/arthurlee/goa/database"
	//"log"
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

type CourseListRes struct {
	GoaBaseRes
	Courses []models.Course `json:"course_list"`
}

func List(w http.ResponseWriter, r *http.Request) {
	resCourseList := CourseListRes{}

	courseList := models.NewCourseList()
	err := database.GetList(courseList)
	if err != nil {
		resCourseList.setError("1", err.Error())
	} else {
		resCourseList.Courses = courseList.Courses
		resCourseList.setOK()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resCourseList)
}
