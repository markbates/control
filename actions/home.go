package actions

import (
	"errors"
	"fmt"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
	"github.com/markbates/control/mcu"
	"github.com/markbates/control/mcu/codes"
	"github.com/markbates/control/mcu/transport"
	"github.com/markbates/safe"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	if Device.IsZero() {
		c.Flash().Add("danger", "Could not find an IAC device named MCU")
	}
	return c.Render(200, r.HTML("index.html"))
}

func Trigger(c buffalo.Context) error {
	key := c.Param("events")
	events, ok := mcu.Events.Load(key)

	onErr := func(err error) error {
		c.Set("err", err)
		return c.Render(200, r.JavaScript("trigger.js"))
	}

	if !ok {
		err := fmt.Errorf("can't find %s", key)
		return onErr(err)
	}
	return onErr(Device.Write(events))
}

func ws(c buffalo.Context) error {
	es, err := render.NewEventSource(c.Response())
	if err != nil {
		return err
	}
	for {
		err := safe.RunE(func() error {
			if Device.In == nil {
				return errors.New("nil input")
			}
			events, err := Device.In.Read(1)
			if err != nil {
				return err
			}
			if len(events) == 0 {
				return nil
			}
			e := events[0]
			fmt.Println("### e.Status ->", e.Status)
			fmt.Println("### e ->", e)
			switch e.Status {
			case int64(240):

			case transport.CODE:
				switch e.Data1 {
				case codes.PLAY:
					if e.Data2 == codes.ON {
						es.Write("transport", "transport.Play\n")
					}
				case codes.STOP:
					if e.Data2 == codes.ON {
						es.Write("transport", "transport.Stop\n")
					}
				}
			}
			es.Flush()
			return nil
		})
		if err != nil {
			return err
		}
	}
	return nil
}
