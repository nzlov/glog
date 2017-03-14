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
	c := &Console{}
	c.BaseListener = listener.NewBaseListener(c)
	return c
}
func (self *Console) Name() string {
	return "Console"
}

func (self *Console) Event(e glog.Event) {
	if e.Data == nil {
		fmt.Printf("[%s][%s] %s", e.Level, e.Time.Format("2006-01-02 15:04:05"), e.Message)
	} else {
		fmt.Printf("[%s][%s] %s", e.Level, e.Time.Format("2006-01-02 15:04:05"), e.Message)
		fmt.Printf("%s", e.Data)
	}
}
func (self *Console) Close() {}
