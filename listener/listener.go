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
}

func NewBaseListener(l listener) *BaseListener {
	return &BaseListener{
		l:      l,
		notify: make(chan glog.Event, glog.LOGCHANSIZE),
	}
}
func (self *BaseListener) Notify() chan glog.Event {
	return self.notify
}

func (self *BaseListener) event() {
	for e := range self.notify {
		self.l.Event(e)
	}
	self.l.Close()
	glog.QuitWait.Done()
}
func (self *BaseListener) Stop() {
	close(self.notify)
}

func (self *BaseListener) Start() {
	go self.event()
}

func (self *BaseListener) SetListener(l listener) {
	self.l = l
}
