package resources

import (
	"github.com/arthurlee/goa-doc/samples/hello/resources/course"
	"github.com/arthurlee/goa/route"
)

func Register() {
	route.Get("/", HandleHello)
	route.Get("/course/list", course.List)
	route.Post("/course/create", course.Create)
}
