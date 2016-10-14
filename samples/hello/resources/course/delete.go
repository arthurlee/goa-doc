package course

import (
	"github.com/arthurlee/goa-doc/samples/hello/models"
	"github.com/arthurlee/goa/database"
	"github.com/arthurlee/goa/route"
	"log"
	"strings"
)

func Delete(res *route.GoaResponse) {
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
		log.Printf("rowsDeleted = %d", rowsDeleted)
		res.SendOK()
	}
}
