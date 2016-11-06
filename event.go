package goevent

type Event struct {
	*Delegate
}

func NewEvent() *Event {
	return &Event{NewDelegate()}
}

func (e *Event)Fire(s Source, a EventArg) {
	e.Exec(s, a)
}
