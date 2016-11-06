package main

import (
	. "github.com/blackspace/goevent"
	"log"
	"time"
)

const (
	TickEvent = iota
)

type MyTick struct {
	*time.Ticker
	TickEvent *Event
}

func NewMyTick() *MyTick {
	return &MyTick{Ticker: time.NewTicker(time.Second), TickEvent: DefineOrTakeEvent(TickEvent)}
}

func main() {
	t := NewMyTick()
	go func() {
		for {
			_ = <-t.Ticker.C
			t.TickEvent.Fire(t, time.Now())
		}
	}()

	t.TickEvent.AddHandler(func(s Source, a EventArg) {
		log.Println(a.(time.Time))
	})

	Run()
}
