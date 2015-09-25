package file

import (
	"fmt"
	"os"

	"github.com/nzlov/glog"
)

type File struct {
	name string
	f    *os.File
}

func New(name string) (*File, error) {
	f := &File{}
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
func (self *File) Event(e glog.Event) {
	fmt.Fprintf(self.f, "[%s][%s]%s %s", e.Level, e.Time.Format("2006-01-02 15:04:05"), e.Data, e.Message)
}
func (self *File) Close() {
	self.f.Close()
}
