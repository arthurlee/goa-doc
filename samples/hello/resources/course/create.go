package course

import (
	"github.com/arthurlee/goa-doc/samples/hello/models"
	"github.com/arthurlee/goa/database"
	"github.com/arthurlee/goa/route"
	"log"
	"strings"
)

func Create(res *route.GoaResponse) {
	name := strings.Join(res.Form["name"], "")
	summary := strings.Join(res.Form["summary"], "")

	if len(name) == 0 {
		res.SendError("1", "parameter name could not be empty")
		return
	}

	if len(summary) == 0 {
		res.SendError("1", "parameter summary could not be empty")
		return
	}

	course := models.Course{Name: name, Summary: summary}
	rowsInserted, lastInsertId, err := database.Create(course.GetInserter())
	if err != nil {
		res.SendError("1", err.Error())
	} else {
		log.Printf("rowsInserted = %d, lastInsertId = %d", rowsInserted, lastInsertId)
		res.SendOK()
	}
}
