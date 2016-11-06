package goevent

type Delegate struct {
	Handlers []Handler
}

type Source interface{}
type EventArg interface{}

type Handler func(source Source, arg EventArg)

func NewDelegate() *Delegate {
	return &Delegate{make([]Handler, 0, 1<<8)}
}

func (d *Delegate)AddHandler(h Handler) {
	d.Handlers = append(d.Handlers, h)
}

func (d *Delegate)HasHandler() bool {
	if len(d.Handlers) == 0 {
		return false
	} else {
		return true
	}
}

func (d *Delegate)Exec(s Source, a EventArg) {
	if !d.HasHandler() {
		panic("This delegate has any handler")
	} else {
		for _, h := range d.Handlers {
			if h == nil {
				panic("This delegate has a nil handler")
			} else {
				h(s, a)
			}
		}
	}
}
