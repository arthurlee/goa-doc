package resources

import (
	"github.com/arthurlee/goa/route"
)

func Register() {
	route.Register("/", HandleHello)
}
