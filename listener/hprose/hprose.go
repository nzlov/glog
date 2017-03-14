package hprose

import (
	"github.com/hprose/hprose-golang/rpc"

	"github.com/nzlov/glog"
	"github.com/nzlov/glog/listener"
)

type Log struct {
	Log func(glog.Event)
}

type Hprose struct {
	*listener.BaseListener

	host   string
	client rpc.Client
	log    *Log
}

func New(host string) *Hprose {
	c := rpc.NewClient(host)
	l := &Log{}
	c.UseService(l)
	u := &Hprose{
		host:   host,
		client: c,
		log:    l,
	}
	u.BaseListener = listener.NewBaseListener(u)
	return u
}

func (self *Hprose) Name() string {
	return "hprose:" + self.host
}
func (self *Hprose) Event(e glog.Event) {
	self.log.Log(e)
}
func (self *Hprose) Close() {
	self.client.Close()
}
