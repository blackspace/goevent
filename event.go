package goevent

type Event struct {
	*Delegate
}

func NewEvent()  *Event{
	return &Event{_event_run_loop.CreateNewDelegate()}
}

func (e *Event)Fire(s Source,a EventArg) {
	_event_run_loop.AddEvent(e.EventId,s,a)
}



