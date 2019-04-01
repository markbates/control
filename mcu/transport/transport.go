package transport

import (
	"github.com/markbates/control/mcu/codes"
	"github.com/markbates/portmidi"
)

const (
	CODE int64 = 144
)

var Play = []portmidi.Event{
	{
		Status: CODE,
		Data1:  codes.PLAY,
		Data2:  codes.ON,
	},
	{
		Status: CODE,
		Data1:  codes.PLAY,
		Data2:  codes.OFF,
	},
}

var Stop = []portmidi.Event{
	{
		Status: CODE,
		Data1:  codes.STOP,
		Data2:  codes.ON,
	},
	{
		Status: CODE,
		Data1:  codes.STOP,
		Data2:  codes.OFF,
	},
}

var Record = []portmidi.Event{
	{
		Status: CODE,
		Data1:  codes.RECORD,
		Data2:  codes.ON,
	},
	{
		Status: CODE,
		Data1:  codes.RECORD,
		Data2:  codes.OFF,
	},
}

var Forward = []portmidi.Event{
	{
		Status: CODE,
		Data1:  codes.FORWARD,
		Data2:  codes.ON,
	},
	{
		Status: CODE,
		Data1:  codes.FORWARD,
		Data2:  codes.OFF,
	},
}

var Rewind = []portmidi.Event{
	{
		Status: CODE,
		Data1:  codes.REWIND,
		Data2:  codes.ON,
	},
	{
		Status: CODE,
		Data1:  codes.REWIND,
		Data2:  codes.OFF,
	},
}
