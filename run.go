package goevent

var  _event_run_loop *EventLoop

func init() {
	_event_run_loop =NewEventLoop()
}

func Run() {
	_event_run_loop.Run()
}

