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
	if err := Device.Write(transport.Play); err != nil {
		return err
	}
	return c.Render(200, r.JavaScript("play.js"))
}

func Stop(c buffalo.Context) error {
	if err := Device.Write(transport.Stop); err != nil {
		return err
	}
	return c.Render(200, r.JavaScript("stop.js"))
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
