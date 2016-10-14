package resources

import (
	"github.com/arthurlee/goa/server"
)

func HandleHello(res *server.GoaResponse) {
	person := struct {
		Name string `json:"name", db:"pk"`
		Age  int    `json:"age"`
	}{"Tom", 12}

	res.SendJson(person)
}
