package main

import (
	. "github.com/blackspace/goevent"
	"log"
	"time"
)

type MyTick struct {
	*time.Ticker
	TickEvent *Event
}

func NewMyTick() *MyTick {
	return &MyTick{Ticker: time.NewTicker(time.Second), TickEvent: NewEvent()}
}

func main() {
	t := NewMyTick()

	t.TickEvent.AddHandler(func(s Source, a EventArg) {
		log.Println(a.(time.Time))
	})

	for {
		_ = <-t.Ticker.C
		t.TickEvent.Fire(t, time.Now())
	}
}

