package course

import (
	"github.com/arthurlee/goa-doc/samples/hello/models"
	"github.com/arthurlee/goa/database"
	"github.com/arthurlee/goa/logger"
	"github.com/arthurlee/goa/server"
	"strings"
)

// http -f post http://127.0.0.1:5400/course/delete name=mathmatic

func Delete(res *server.GoaResponse) {
	name := strings.Join(res.Form["name"], "")

	if len(name) == 0 {
		res.SendError("1", "parameter name could not be empty")
		return
	}

	course := models.Course{Name: name}
	rowsDeleted, err := database.Delete(course.GetDeleter())
	if err != nil {
		res.SendError("1", err.Error())
	} else {
		logger.Debug("rowsDeleted = %d", rowsDeleted)
		res.SendOK()
	}
}
