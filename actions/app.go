package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	paramlogger "github.com/gobuffalo/mw-paramlogger"
	"github.com/markbates/control/mcu"
	"github.com/markbates/control/mcu/transport"
	"github.com/markbates/oncer"
)

var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App
var Device mcu.Device

func init() {
	oncer.Do("midi.setup", func() {
		Device = mcu.MCU
		// go func() {
		// 	err = Bus.Start()
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// }()
	})
}

func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env:         ENV,
			SessionName: "_control_session",
		})
		app.Use(func(next buffalo.Handler) buffalo.Handler {
			return func(c buffalo.Context) error {
				c.Set("device", Device)
				c.Set("transportPlay", transport.Play)
				return next(c)
			}
		})
		// Log request parameters (filters apply).
		app.Use(paramlogger.ParameterLogger)

		app.GET("/", HomeHandler)
		app.ANY("/ws", ws)
		app.POST("/trigger", Trigger)

		app.ServeFiles("/", assetsBox) // serve files from the public directory
	}

	return app
}
