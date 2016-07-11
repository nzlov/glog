package file

import (
	"fmt"
	"os"

	"github.com/nzlov/glog"
)

type File struct {
	notify chan glog.Event
	quit   chan bool
	name   string
	f      *os.File
}

func New(name string) (*File, error) {
	f := &File{
		notify: make(chan glog.Event, 10),
		quit:   make(chan bool),
	}
	f.name = name
	var err error
	f.f, err = os.Create(name)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (self *File) Name() string {
	return self.name
}

func (self *File) Notify() chan glog.Event {
	return self.notify
}

func (self *File) event() {
	for {
		e, ok := <-self.notify
		if !ok {
			break
		}
		fmt.Fprintf(self.f, "[%s][%s]%s %s", e.Level, e.Time.Format("2006-01-02 15:04:05"), e.Data, e.Message)
	}
	self.f.Close()
	self.quit <- true
}
func (self *File) Stop() chan bool {
	close(self.notify)
	return self.quit
}

func (self *File) Start() {
	go self.event()
}
