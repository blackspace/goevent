package goevent

import (
	"fmt"
)

type LoopItem struct {
	_event_id  int
	_source    Source
	_event_arg EventArg
}

type EventLoop struct {
	_channel   chan LoopItem
	_delegates []*Delegate
}

func NewEventLoop() *EventLoop {
	return &EventLoop{_channel: make(chan LoopItem, 1<<10), _delegates: make([]*Delegate, 0, 1<<8)}
}

func (el *EventLoop) GetDelegate(eid int) *Delegate {
	var d *Delegate

	for i := 0; i < len(el._delegates); i++ {
		if el._delegates[i].Id == eid {
			d = el._delegates[i]
			break
		}
	}

	if d == nil {
		return nil
	} else {
		return d
	}
}

const MaxEventCount = 100000

func (el *EventLoop) CreateNewDelegate() *Delegate {

	for id := 0; id < MaxEventCount; id++ {
		d := el.GetDelegate(id)

		if d == nil {
			d := NewDelegate(id)
			el._delegates = append(el._delegates, d)
			return d
		}
	}

	panic(fmt.Sprintf("Too many Event , MaxEventCount is %v\n", MaxEventCount))
}

func (el *EventLoop) GetOrCreateDelegate(eid int) *Delegate {
	d := el.GetDelegate(eid)

	if d == nil {
		d = NewDelegate(eid)
		el._delegates = append(el._delegates, d)
		return d
	} else {
		return d
	}
}

func (el *EventLoop) AddEvent(eid int, s Source, a EventArg) {
	el._channel <- LoopItem{_event_id: eid, _source: s, _event_arg: a}
}

func (el *EventLoop) Run() {
	for {
		e := <-el._channel

		var d *Delegate = el.GetDelegate(e._event_id)

		if d == nil {
			panic(fmt.Sprintf("Can't find a delegate for the event id %v\n", e._event_id))
		} else if !d.HasHandler() {
			panic(fmt.Sprintf("This delegate of event id %v has Not a handler", e._event_id))
		} else {
			d.Exec(e._source, e._event_arg)
		}
	}
}
