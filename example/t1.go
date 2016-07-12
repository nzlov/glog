package main

import (
	"time"

	"github.com/nzlov/glog"
	"github.com/nzlov/glog/listener/console"
	"github.com/nzlov/glog/listener/file"
)

func main() {
	defer glog.Close()
	glog.SetLevel(glog.DebugLevel)
	f, err := file.New("l.log")
	if err != nil {
		panic(err)
	}
	glog.Register(f)
	glog.Register(console.New())

	a()
}
func a() {
	go func() {
		i := 1
		for {
			glog.Infoln(i)
			glog.NewFiled().
				Set("k", "v").
				Set("k1", "v1").
				Infoln(i)
			glog.NewTagFiled("tag").
				Set("k", "v").Infoln(i)
			i++
			time.Sleep(time.Millisecond)
		}
	}()
	time.Sleep(time.Second)
	a1()
}
func a1() {
	panic("eeee")
}
