package resources

import (
	"github.com/arthurlee/goa-doc/samples/hello/resources/course"
	"github.com/arthurlee/goa/route"
	"github.com/arthurlee/goa/server"
)

func Register() {
	route.Get("/", HandleHello)

	// test script:  http http://127.0.0.1:5400/heartbeat
	route.Get("/heartbeat", func(res *server.GoaResponse) {
		res.SendOK()
	})

	// test script:  http http://127.0.0.1:5400/welcome
	route.Get("/welcome", server.GetSendOK())

	// test script:  http http://127.0.0.1:5400/error
	route.Get("/error", server.GetSendError("-1", "You are wrong!"))

	route.Get("/course/list", course.List)
	route.Post("/course/create", course.Create)
	route.Post("/course/update", course.Update)
	route.Post("/course/delete", course.Delete)
}
