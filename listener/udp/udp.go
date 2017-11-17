package udp

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/nzlov/glog"
	"github.com/nzlov/glog/listener"
)

type Udp struct {
	*listener.BaseListener

	host string
	udp  *net.UDPConn
}

func New(host string) (*Udp, error) {
	return NewWithOption(host, glog.DefaultOption)
}
func NewWithOption(host string, o *glog.Option) (*Udp, error) {
	addr, err := net.ResolveUDPAddr("udp4", host)
	if err != nil {
		return nil, err
	}
	udp, err := net.DialUDP("udp4", nil, addr)
	if err != nil {
		return nil, err
	}
	u := &Udp{
		host: host,
		udp:  udp,
	}
	u.BaseListener = listener.NewBaseListener(u, o)
	return u, nil
}

func (self *Udp) Event(e glog.Event) {
	b, err := json.Marshal(e)
	if err != nil {
		fmt.Println(err)
		return
	}
	self.udp.Write(b)
}
func (self *Udp) Close() {
	self.udp.Close()
}
