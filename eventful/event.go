package eventful

import "gioui.org/io/event"

type DumbQueue struct{}

func (e DumbQueue) Events(t event.Tag) []event.Event {
	return []event.Event{}
}

func DummyQueue() event.Queue {
	return DumbQueue{}
}
