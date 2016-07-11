package udp

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/nzlov/glog"
)

type Udp struct {
	host string
	udp  *net.UDPConn

	notify chan glog.Event
	quit   chan bool
}

func New(host string) (*Udp, error) {
	u := &Udp{
		notify: make(chan glog.Event, 10),
		quit:   make(chan bool),
	}
	addr, err := net.ResolveUDPAddr("udp4", host)
	if err != nil {
		return nil, err
	}
	udp, err := net.DialUDP("udp4", nil, addr)
	if err != nil {
		return nil, err
	}
	u.udp = udp
	u.host = host
	return u, nil
}

func (self *Udp) Name() string {
	return "udp:" + self.host
}
func (self *Udp) Notify() chan glog.Event {
	return self.notify
}

func (self *Udp) event() {
	for {
		e, ok := <-self.notify
		if !ok {
			break
		}
		b, err := json.Marshal(e)
		if err != nil {
			fmt.Println(err)
			return
		}
		self.udp.Write(b)
	}
	self.udp.Close()
	self.quit <- true
}
func (self *Udp) Stop() chan bool {
	close(self.notify)
	return self.quit
}

func (self *Udp) Start() {
	go self.event()
}
