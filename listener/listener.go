package listener

import (
	"github.com/nzlov/glog"
)

type listener interface {
	Event(e glog.Event)
	Close()
}

type BaseListener struct {
	l      listener
	notify chan glog.Event
	quit   chan bool
}

func NewBaseListener(l listener) *BaseListener {
	return &BaseListener{
		l:      l,
		notify: make(chan glog.Event, 10),
		quit:   make(chan bool),
	}
}
func (self *BaseListener) Notify() chan glog.Event {
	return self.notify
}

func (self *BaseListener) event() {
	for {
		e, ok := <-self.notify
		if !ok {
			break
		}
		self.l.Event(e)
	}
	self.l.Close()
	self.quit <- true
}
func (self *BaseListener) Stop() chan bool {
	close(self.notify)
	return self.quit
}

func (self *BaseListener) Start() {
	go self.event()
}

func (self *BaseListener) SetListener(l listener) {
	self.l = l
}
