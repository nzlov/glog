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
}

func New(host string) (*Udp, error) {
	u := &Udp{}
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
	return "udp" + self.host
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
