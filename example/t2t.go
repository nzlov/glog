package main

import (
	"runtime"
	"time"

	"github.com/nzlov/glog"
	"github.com/nzlov/glog/listener/console"
)

var log *glog.Logger

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() * 2)
	log = glog.New()
	defer func() {
		log.Close()
	}()
	//log.AddListener(colorconsole.New(colorconsole.ShowTypeFore))
	log.AddListener(console.New())
	for i := 0; i < 10; i++ {
		go func(gid int) {
			n := 0
			for {
				log.Infoln(gid, n)
				n++
			}
		}(i)
	}
	time.Sleep(time.Second * 10)
}
