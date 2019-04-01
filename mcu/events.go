package mcu

import (
	"sort"
	"sync"

	"github.com/markbates/control/mcu/transport"
	"github.com/markbates/portmidi"
)

var Events = func() EventMap {
	e := EventMap{}
	e.Store("transport.Play", transport.Play)
	e.Store("transport.Stop", transport.Stop)
	e.Store("transport.Record", transport.Record)
	e.Store("transport.Forward", transport.Forward)
	e.Store("transport.Rewind", transport.Rewind)
	return e
}()

// EventMap wraps sync.Map and uses the following types:
// key:   string
// value: string
type EventMap struct {
	data sync.Map
}

// Delete the key from the map
func (m *EventMap) Delete(key string) {
	m.data.Delete(key)
}

// Load the key from the map.
// Returns string or bool.
// A false return indicates either the key was not found
// or the value is not of type string
func (m *EventMap) Load(key string) ([]portmidi.Event, bool) {
	i, ok := m.data.Load(key)
	if !ok {
		return []portmidi.Event{}, false
	}
	s, ok := i.([]portmidi.Event)
	return s, ok
}

// LoadOrStore will return an existing key or
// store the value if not already in the map
func (m *EventMap) LoadOrStore(key string, value string) ([]portmidi.Event, bool) {
	i, _ := m.data.LoadOrStore(key, value)
	s, ok := i.([]portmidi.Event)
	return s, ok
}

// Range over the string values in the map
func (m *EventMap) Range(f func(key string, value []portmidi.Event) bool) {
	m.data.Range(func(k, v interface{}) bool {
		key, ok := k.(string)
		if !ok {
			return false
		}
		value, ok := v.([]portmidi.Event)
		if !ok {
			return false
		}
		return f(key, value)
	})
}

// Store a string in the map
func (m *EventMap) Store(key string, value []portmidi.Event) {
	m.data.Store(key, value)
}

// Keys returns a list of keys in the map
func (m *EventMap) Keys() []string {
	var keys []string
	m.Range(func(key string, value []portmidi.Event) bool {
		keys = append(keys, key)
		return true
	})
	sort.Strings(keys)
	return keys
}
