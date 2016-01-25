package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	_ "github.com/martini-contrib/sessions"
	"github.com/tcyd/goconf"
)

var m *martini.Martini

func init() {
	m = martini.New()
	// Setup config
	conf, err := goconf.ReadConfigFile(`app.conf`)
	if err != nil {
		panic(err)
	}
	m.Map(conf)

	// Setup enviroment
	env, err := conf.GetString(`default`, `env`)
	if err != nil {
		env = `dev`
	}
	switch env {
	case `prod`:
		martini.Env = martini.Prod
	case `test`:
		martini.Env = martini.Test
	default:
		martini.Env = martini.Dev
	}

	// Setup middlewares
	m.Use(martini.Recovery())
	m.Use(render.Renderer())
	m.Use(martini.Logger())

	staticPath, err := conf.GetString(`default`, `staticpath`)
	if err != nil {
		panic(err)
	}
	m.Use(martini.Static(staticPath))

	// sessionRedisHost, _ := conf.GetString(`sessionredis`, `host`)
	// sessionRedisPort, _ := conf.GetString(`sessionredis`, `port`)
	// sessionRedisPassword, _ := conf.GetString(`sessionredis`, `password`)
	// sessionRedisKey, _ := conf.GetString(`sessionredis`, `key`)
	// //store := sessions.NewCookieStore([]byte("secret123"))
	// store, err := sessions.NewRediStore(10, `tcp`, sessionRedisHost+`:`+sessionRedisPort, sessionRedisPassword, []byte(sessionRedisKey))
	// if err != nil {
	// 	panic(err)
	// }
	// m.Use(sessions.Sessions("session", store))

	// Setup routes
	Routes(m)
}

func main() {
	m.RunOnAddr(`` + `:` + `7429`)
}
