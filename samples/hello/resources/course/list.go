package course

import (
	"github.com/arthurlee/goa-doc/samples/hello/models"
	"github.com/arthurlee/goa/database"
	"github.com/arthurlee/goa/route"
	//"log"
)

type CourseListRes struct {
	route.GoaBaseRes
	Courses []models.Course `json:"courses"`
}

func List(res *route.GoaResponse) {
	courseList := models.NewCourseList()
	err := database.GetList(courseList)
	if err != nil {
		res.SendError("1", err.Error())
	} else {
		resCourseList := CourseListRes{route.GoaBaseRes{"0", "ok"}, courseList.Courses}
		res.SendJson(resCourseList)
	}
}
