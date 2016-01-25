package main

import (
	"github.com/go-martini/martini"
	"gobed/controllers"
	"gobed/middlewares"
)

func Routes(m *martini.Martini) {
	r := martini.NewRouter()

	r.Get(`/`, controllers.Index)
	r.Get(`/hello`, middlewares.ApiVerify(), controllers.Hello)

	// Add the router action
	m.Action(r.Handle)
}
