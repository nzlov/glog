package file

import (
	"fmt"
	"os"

	"github.com/nzlov/glog"
)

type File struct {
	name string
	f    os.File
}

func (self *File) Name() string {
	return self.name
}
func (self *File) Event(e glog.Event) {
	fmt.Fprintf(self.f, "[%s][%s]%s %s", e.Level, e.Time.Format("2006-01-02 15:04:05"), e.Data, e.Message)
}
