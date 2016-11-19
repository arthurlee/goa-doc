package course

import (
	"github.com/arthurlee/goa-doc/samples/hello/models"
	"github.com/arthurlee/goa/database"
	"github.com/arthurlee/goa/logger"
	"github.com/arthurlee/goa/server"
	"strings"
)

// http -f post http://127.0.0.1:5400/course/update name=mathmatic summary="math higher and higher"

func Update(res *server.GoaResponse) {
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
	rowsUpdated, err := database.Update(course.GetUpdater())
	if err != nil {
		res.SendError("1", err.Error())
	} else {
		logger.Debug("rowsUpdated = %d", rowsUpdated)
		res.SendOK()
	}
}
