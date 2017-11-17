package listener

import (
	"github.com/nzlov/glog"
)

type listener interface {
	Event(e glog.Event)
	Close()
}

type BaseListener struct {
	id      string
	l       listener
	option  *glog.Option
	running bool
	pause   bool
	notify  chan glog.Event
	done    chan struct{}
}

func NewBaseListener(l listener, o *glog.Option) *BaseListener {
	return &BaseListener{
		id:     glog.GenID(),
		l:      l,
		option: o,
		notify: make(chan glog.Event, o.NumCache),
		done:   make(chan struct{}),
	}
}
func (self *BaseListener) Notify() chan glog.Event {
	return self.notify
}

func (self *BaseListener) event() {
	for e := range self.notify {
		if self.pause {
			self.l.Event(e)
		}
	}
	self.l.Close()
	self.done <- struct{}{}
}
func (self *BaseListener) Stop() {
	if !self.running {
		return
	}
	self.running = false
	close(self.notify)
	<-self.done
}
func (self *BaseListener) Pause(b bool) {
	self.pause = b
}
func (self *BaseListener) ID() string {
	return self.id
}
func (self *BaseListener) Start() {
	self.running = true
	self.pause = true
	go self.event()
}

func (self *BaseListener) SetListener(l listener) {
	self.l = l
}
func (self *BaseListener) Option() glog.Option {
	return *(self.option)
}
