package actions

import (
	"log"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	paramlogger "github.com/gobuffalo/mw-paramlogger"
	"github.com/markbates/control/mcu"
	"github.com/markbates/oncer"
)

var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App
var Bus *mcu.Bus

func init() {
	oncer.Do("midi.setup", func() {
		var err error
		Bus, err = mcu.New()
		if err != nil {
			log.Fatal(err)
		}
	})
}

func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env:         ENV,
			SessionName: "_control_session",
		})
		// Log request parameters (filters apply).
		app.Use(paramlogger.ParameterLogger)

		app.GET("/", HomeHandler)
		app.GET("/play", Play)
		app.GET("/stop", Stop)

		app.ServeFiles("/", assetsBox) // serve files from the public directory
	}

	return app
}
