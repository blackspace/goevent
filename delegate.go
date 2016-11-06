package goevent

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

func (d *Delegate)Exec(s Source,a EventArg) {
	for _,h:=range d.Handlers {
		h(s,a)
	}
}
