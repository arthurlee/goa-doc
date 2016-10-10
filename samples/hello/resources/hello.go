package resources

import (
	"github.com/arthurlee/goa/route"
)

func HandleHello(res *route.GoaResponse) {
	person := struct {
		Name string `json:"name", db:"pk"`
		Age  int    `json:"age"`
	}{"Tom", 12}

	res.SendJson(person)
}
