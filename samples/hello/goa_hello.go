package main

import (
	"./resources"
	"github.com/arthurlee/goa"
)

func main() {
	resources.Register()
	goa.Serve()
}
