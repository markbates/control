package mcu

import (
	"errors"
	"fmt"

	"github.com/markbates/safe"
	"github.com/rakyll/portmidi"
)

type Bus struct {
	In     *portmidi.Stream
	Out    *portmidi.Stream
	Device *portmidi.DeviceInfo
}

func New() (*Bus, error) {
	var iid portmidi.DeviceID
	var oid portmidi.DeviceID
	var device *portmidi.DeviceInfo

	err := safe.Run(func() {
		iid = portmidi.DefaultInputDeviceID()  // returns the ID of the system default input
		oid = portmidi.DefaultOutputDeviceID() // returns the ID of the system default output
		device = portmidi.Info(oid)
	})
	if err != nil {
		return nil, err
	}

	in, err := portmidi.NewInputStream(iid, 1024)
	if err != nil {
		return nil, err
	}

	out, err := portmidi.NewOutputStream(oid, 1024, 1024)
	if err != nil {
		return nil, err
	}
	out.SetChannelMask(1)

	b := &Bus{
		Device: device,
		In:     in,
		Out:    out,
	}
	return b, nil
}

func (b *Bus) Close() error {
	if b.In != nil {
		b.In.Close()
	}
	if b.Out != nil {
		b.Out.Close()
	}
	return nil
}

func (b *Bus) Start() error {
	if b.Out == nil {
		return errors.New("there is no output stream")
	}
	for e := range b.Out.Listen() {
		fmt.Printf("### out e -> %+v\n", e)
	}
	return nil
}

func (b *Bus) Write(events []portmidi.Event) error {
	return b.Out.Write(events)
}
