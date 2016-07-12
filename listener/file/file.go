package file

import (
	"fmt"
	"os"

	"github.com/nzlov/glog"
	"github.com/nzlov/glog/listener"
)

type File struct {
	*listener.BaseListener

	name string
	f    *os.File
}

func New(name string) (*File, error) {
	f := &File{
		name: name,
	}
	var err error
	f.f, err = os.Create(name)
	if err != nil {
		return nil, err
	}
	f.BaseListener = listener.NewBaseListener(f)
	return f, nil
}

func (self *File) Name() string {
	return self.name
}

func (self *File) Event(e glog.Event) {
	if e.Data == nil {
		fmt.Fprintf(self.f, "[%s][%s] %s", e.Level, e.Time.Format("2006-01-02 15:04:05"), e.Message)

	} else {
		fmt.Fprintf(self.f, "[%s][%s]%s %s", e.Level, e.Time.Format("2006-01-02 15:04:05"), e.Data, e.Message)
	}
}

func (self *File) Close() {
	self.f.Close()
}
