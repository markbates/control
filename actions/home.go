package actions

import (
	"fmt"

	"github.com/gobuffalo/buffalo"
	"github.com/markbates/control/mcu"
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

// func Switch(c buffalo.Context) error {
// 	// in := c.Param("InDeviceID")
// 	out := c.Param("OutDeviceID")
//
// 	// iid, err := strconv.Atoi(in)
// 	// if err != nil {
// 	// 	iid = -1
// 	// }
// 	oid, err := strconv.Atoi(out)
// 	if err != nil {
// 		oid = -1
// 	}
// 	Device, err = mcu.Find(portmidi.DeviceID(oid))
// 	if err != nil {
// 		return err
// 	}
// 	return c.Redirect(301, "/")
// }
