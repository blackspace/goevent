package goevent

import (
	"reflect"
)


type LoopItem struct {
	_event_id  int
	_source    Source
	_event_arg EventArg
}

type EventLoop struct {
	_channel  chan LoopItem
	_delegates []*Delegate
}

func NewEventLoop() *EventLoop {
	return &EventLoop{_channel:make(chan LoopItem,1<<10), _delegates:make([]*Delegate,0,1<<8)}
}


func (el *EventLoop)GetDelegate(eid int) *Delegate {
	var d *Delegate

	for i:=0;i<len(el._delegates);i++ {
		if el._delegates[i].EventId == eid {
			d=el._delegates[i]
			break
		}
	}


	if d==nil {
		return nil
	} else {
		return d
	}
}


func (el *EventLoop)CreateNewDelegate() *Delegate {
	for id := 0; id < 100000; id++ {
		d := el.GetDelegate(id)

		if d == nil {
			d := NewDelegate(id)
			el._delegates = append(el._delegates, d)
			return d
		}
	}

	return nil
}


func (el *EventLoop)AddEvent(eid int,s Source,a EventArg) {
	el._channel<-LoopItem{_event_id:eid, _source:s, _event_arg:a}
}

func (el *EventLoop)Run() {
	for {
		e:=<-el._channel

		var d *Delegate =el.GetDelegate(e._event_id)

		if d!=nil {
			d.Exec(e._source,e._event_arg)
		} else {
			panic("Can't find a handler for the event of "+reflect.TypeOf(e).String()+".")
		}
	}
}


