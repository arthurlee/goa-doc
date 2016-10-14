package course

import (
	"github.com/arthurlee/goa-doc/samples/hello/models"
	"github.com/arthurlee/goa/database"
	"github.com/arthurlee/goa/server"
	//"log"
)

type CourseListRes struct {
	server.GoaBaseRes
	Courses []models.Course `json:"courses"`
}

func List(res *server.GoaResponse) {
	courseList := models.NewCourseList()
	err := database.GetList(courseList)
	if err != nil {
		res.SendError("1", err.Error())
	} else {
		resCourseList := CourseListRes{server.GoaBaseRes{"0", "ok"}, courseList.Courses}
		res.SendJson(resCourseList)
	}
}
