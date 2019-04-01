package mcu

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/markbates/portmidi"
	"github.com/pkg/errors"
)

type Device struct {
	*portmidi.DeviceInfo
	In  *portmidi.Stream
	Out *portmidi.Stream
}

func (d Device) String() string {
	b, _ := json.Marshal(d)
	return string(b)
}

func (d Device) Write(events []portmidi.Event) error {
	return d.Out.Write(events)
}

var MCU = func() Device {
	d := Device{}
	var info *portmidi.DeviceInfo
	for i := 0; i < portmidi.CountDevices(); i++ {
		id := portmidi.DeviceID(i)
		info = portmidi.Info(id)
		if info == nil {
			continue
		}
		if !strings.Contains(info.Name, "MCU") {
			continue
		}
		d.DeviceInfo = info
		if info.IsInputAvailable {
			in, err := portmidi.NewInputStream(id, 128)
			if err != nil {
				log.Fatal(errors.WithMessage(err, "setting up input: "+info.Name))
			}
			d.In = in
		}
		if info.IsOutputAvailable {
			out, err := portmidi.NewOutputStream(id, 128, 128)
			if err != nil {
				log.Fatal(errors.WithMessage(err, "setting up output: "+info.Name))
			}
			d.Out = out
		}
	}

	return d
}()

// var Devices = func() map[portmidi.DeviceID]Device {
// 	m := map[portmidi.DeviceID]Device{}
//
// 	m[-1] = func() Device {
// 		d := Device{ID: -1}
// 		info := portmidi.Info(portmidi.DefaultOutputDeviceID())
// 		if info == nil {
// 			log.Fatal("couldn't find default device")
// 		}
// 		d.DeviceInfo = info
// 		var err error
// 		d.In, err = portmidi.NewInputStream(portmidi.DefaultInputDeviceID(), 512)
// 		if err != nil {
// 			log.Fatal(errors.WithMessage(err, "setting up default input"))
// 		}
// 		d.Out, err = portmidi.NewOutputStream(portmidi.DefaultOutputDeviceID(), 512, 512)
// 		if err != nil {
// 			fmt.Println("### err ->", err)
// 			log.Fatal(errors.WithMessage(err, "setting up default output"))
// 		}
// 		return d
// 	}()
//
// 	for i := 0; i < portmidi.CountDevices(); i++ {
// 		safe.Run(func() {
// 			di := portmidi.DeviceID(i)
// 			d := portmidi.Info(di)
// 			if d == nil {
// 				return
// 			}
// 			var in *portmidi.Stream
// 			var out *portmidi.Stream
// 			var err error
//
// 			if d.IsInputAvailable {
// 				in, err = portmidi.NewInputStream(di, 512)
// 				if err != nil {
// 					// log.Fatal(errors.WithMessage(err, "input: "+d.Name))
// 				}
// 			}
// 			if d.IsOutputAvailable {
// 				out, err = portmidi.NewOutputStream(di, 512, 512)
// 				if err != nil {
// 					// log.Fatal(errors.WithMessage(err, "output: "+d.Name))
// 				}
// 			}
//
// 			out.SetChannelMask(1)
// 			m[di] = Device{
// 				DeviceInfo: d,
// 				ID:         i,
// 				In:         in,
// 				Out:        out,
// 			}
// 		})
// 	}
// 	return m
// }()
//
// // var Devices = map[int]Device{
// // }
// // var Devices = func() []*portmidi.DeviceInfo {
// // 	var devices []*portmidi.DeviceInfo
// // 	err := safe.Run(func() {
// // 		fmt.Printf("### portmidi.CountDevices() -> %+v\n", portmidi.CountDevices())
// // 		for i := 0; i < portmidi.CountDevices(); i++ {
// // 			d := portmidi.Info(portmidi.DeviceID(i))
// // 			fmt.Printf("### d -> %+v\n", d)
// // 			devices = append(devices, d)
// // 		}
// // 	})
// // 	if err != nil {
// // 		log.Fatal(err)
// // 	}
// // 	return devices
// // }()
//
// type Bus struct {
// 	In        *portmidi.Stream
// 	Out       *portmidi.Stream
// 	InDevice  Device
// 	OutDevice Device
// }
//
// func Find(id portmidi.DeviceID) (Device, error) {
// 	if id == -1 {
// 	}
// 	d, ok := Devices[id]
// 	if !ok {
// 		return d, fmt.Errorf("can not find device with id %d", id)
// 	}
// 	return d, nil
// }
//
//
// // func (b *Bus) Close() error {
// // 	if b.In != nil {
// // 		b.In.Close()
// // 	}
// // 	if b.Out != nil {
// // 		b.Out.Close()
// // 	}
// // 	return nil
// // }
// //
// // func (b *Bus) Start() error {
// // 	if b.Out == nil {
// // 		return errors.New("there is no output stream")
// // 	}
// // 	for e := range b.Out.Listen() {
// // 		fmt.Printf("### out e -> %+v\n", e)
// // 	}
// // 	return nil
// // }
// //
// // func (b *Bus) Write(events []portmidi.Event) error {
// // 	return b.Out.Write(events)
// // }
