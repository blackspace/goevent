package goevent

import "strconv"

type Delegate struct {
	EventId  int
	Handlers []Handler
}

type Source interface {}
type EventArg interface {}

type Handler func(source Source,arg EventArg)

func NewDelegate(id int) *Delegate {
	return &Delegate{id,make([]Handler,0,1<<8)}
}

func (d *Delegate)AddHandler(h Handler) {
	d.Handlers =append(d.Handlers,h)
}

func (d *Delegate)HasHandler() bool {
	if len(d.Handlers)==0 {
		return false
	} else {
		return true
	}
}

func (d *Delegate)Exec(s Source,a EventArg) {
	for _,h:=range d.Handlers {
		if h!=nil {
			h(s,a)
		} else {
			panic("this delegate of EventId "+strconv.Itoa(d.EventId)+" has a nil handler")
		}

	}
}
