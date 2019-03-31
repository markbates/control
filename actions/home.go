package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/markbates/control/mcu/transport"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("index.html"))
}

func Play(c buffalo.Context) error {
	if err := Bus.Write(transport.Play); err != nil {
		return err
	}
	return c.Render(200, r.JavaScript("play.js"))
}

func Stop(c buffalo.Context) error {
	if err := Bus.Write(transport.Stop); err != nil {
		return err
	}
	return c.Render(200, r.JavaScript("stop.js"))
}
