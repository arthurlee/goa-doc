package resources

import (
	"github.com/arthurlee/goa-doc/samples/hello/resources/course"
	"github.com/arthurlee/goa/route"
)

func Register() {
	route.Register("/", HandleHello)
	route.Register("/course/list", course.List)
}
