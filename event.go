package goevent

import (
	"strconv"
)



type Event struct {
	*Delegate
}

func DefineOrTakeEvent(eid int)  *Event{
	return &Event{_event_run_loop.GetOrCreateDelegate(eid)}
}

func (e *Event)Fire(s Source,a EventArg) {
	d:=_event_run_loop.GetDelegate(e.EventId)

	if d==nil {
		panic("This delegate with id "+strconv.Itoa(e.EventId)+" is Not existing")
	} else if !d.HasHandler() {
		panic("This event has Not Handler")
	} else {
		_event_run_loop.AddEvent(e.EventId,s,a)
	}
}



