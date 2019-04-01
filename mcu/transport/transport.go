package transport

import "github.com/markbates/portmidi"

const Status = 144

var Play = []portmidi.Event{
	{
		Status: Status,
		Data1:  94,
		Data2:  127,
	},
	{
		Status: Status,
		Data1:  94,
		Data2:  0,
	},
}

var Stop = []portmidi.Event{
	{
		Status: Status,
		Data1:  93,
		Data2:  127,
	},
	{
		Status: Status,
		Data1:  93,
		Data2:  0,
	},
}
