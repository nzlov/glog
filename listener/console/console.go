package console

import (
	"fmt"
	"github.com/nzlov/glog"
)

type Console struct {
}

func New() *Console {
	return &Console{}
}
func (self *Console) Name() string {
	return "Console"
}
func (self *Console) Event(e glog.Event) {
	fmt.Printf("[%s][%s]%s %s", e.Level, e.Time.Format("2006-01-02 15:04:05"), e.Data, e.Message)
}
func (self *Console) Close() {
}
