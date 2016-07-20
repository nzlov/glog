package main

import (
	"time"

	"fmt"

	"github.com/nzlov/glog"
	"github.com/nzlov/glog/listener/colorconsole"
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
	glog.Register(colorconsole.New(colorconsole.ShowTypeDefault))

	a()
}
func a() {
	go func() {
		i := 1
		for {
			glog.Errorln(i)
			glog.Debugln(i)
			glog.NewField().
				Set("k", "v").
				Set("k1", "v1").
				Warnln(i)
			glog.NewTagField("tag").
				Set("k", "v").Infoln(i)
			i++
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(time.Second * 2)
	go glog.Go(a1, 1, 0)
	time.Sleep(time.Second * 5)
}
func a1(x, y int) {
	fmt.Printf("%d / %d = %d\n", x, y, x/y)
}
