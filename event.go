package goevent

type Event struct {
	*Delegate
}

func DefineOrTakeEvent(eid int) *Event {
	return &Event{_event_run_loop.GetOrCreateDelegate(eid)}
}

func (e *Event) Fire(s Source, a EventArg) {
	_event_run_loop.AddEvent(e.EventId, s, a)
}
