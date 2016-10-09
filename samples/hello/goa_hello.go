package main

import (
	"github.com/arthurlee/goa"
	"github.com/arthurlee/goa-doc/samples/hello/resources"
)

func main() {
	resources.Register()
	goa.Serve()
}
