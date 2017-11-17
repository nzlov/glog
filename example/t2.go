package main

import (
	"fmt"
	"time"

	"github.com/nzlov/glog"
	"github.com/nzlov/glog/listener/colorconsole"
)

var log *glog.Logger

func main() {
	log = glog.New()
	defer func() {
		fmt.Println("1111")
		log.Close()
	}()
	log.AddListener(colorconsole.New(colorconsole.ShowTypeDefault))
	//log.AddListener(console.New())

	a()
}
func a() {
	go func() {
		for i := 0; i < 3; i++ {
			log.Errorln(i)
			log.NewField().
				S("k", "v").
				S("k1", "v1").
				Warnln(i)
			log.NewField().
				S("k", "v").
				Infoln(i)
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(time.Second * 2)
	go log.Go(a1, 1, 0)
	time.Sleep(time.Second * 5)
}
func a1(x, y int) {
	fmt.Printf("%d / %d = %d\n", x, y, x/y)
	log.Panicln("ooo")
}
