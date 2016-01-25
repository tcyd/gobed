package controllers

import (
	"github.com/martini-contrib/render"
	"github.com/tcyd/goconf"
)

func Index(r render.Render, conf *goconf.ConfigFile) {
	appName, err := conf.GetString(`default`, `appname`)
	if err != nil {
		appName = ""
	}
	r.HTML(200, "index", appName)
}
