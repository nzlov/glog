package console

import (
	"fmt"

	"github.com/nzlov/glog"
	"github.com/nzlov/glog/listener"
)

type Console struct {
	*listener.BaseListener
}

func New() *Console {
	return NewWithOption(glog.DefaultOption)
}
func NewWithOption(o *glog.Option) *Console {
	c := &Console{}
	c.BaseListener = listener.NewBaseListener(c, o)
	return c
}

func (self *Console) Event(e glog.Event) {
	if e.FuncCall != nil {
		fmt.Printf("[%s][%s]%s %s", e.Level, e.Time.Format("2006-01-02 15:04:05"), e.FuncCall, e.Message)
	} else {
		fmt.Printf("[%s][%s] %s", e.Level, e.Time.Format("2006-01-02 15:04:05"), e.Message)
	}
	if e.Data != nil {
		fmt.Printf("%s", e.Data)
	}
}
func (self *Console) Close() {}
