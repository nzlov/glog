package console

import (
	"fmt"

	"github.com/nzlov/glog"
)

type Console struct {
	notify chan glog.Event
	quit   chan bool
}

func New() *Console {
	return &Console{
		notify: make(chan glog.Event, 10),
		quit:   make(chan bool),
	}
}
func (self *Console) Name() string {
	return "Console"
}

func (self *Console) Notify() chan glog.Event {
	return self.notify
}

func (self *Console) event() {
	for {
		e, ok := <-self.notify
		if !ok {
			break
		}
		fmt.Printf("[%s][%s]%s %s", e.Level, e.Time.Format("2006-01-02 15:04:05"), e.Data, e.Message)
	}
	self.quit <- true
}
func (self *Console) Stop() chan bool {
	close(self.notify)
	return self.quit
}

func (self *Console) Start() {
	go self.event()
}
