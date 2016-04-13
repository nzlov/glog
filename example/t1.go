package main

import (
	"fmt"
	"runtime"

	"github.com/nzlov/glog"
	"github.com/nzlov/glog/listener/console"
)

func panicf() {
	if err := recover(); err != nil {
		errstr := fmt.Sprintf("Runtime error:%v\ntraceback:\n", err)
		i := 1
		for {
			pc, file, line, ok := runtime.Caller(i)
			errstr += fmt.Sprintf("\tstack: %d %v [file:%s][line:%d][func:%s]\n", i, ok, file, line, runtime.FuncForPC(pc).Name())
			i++
			if !ok || i > glog.MAXSTACK {
				break
			}
		}
		fmt.Println(errstr)
	}
}
func main() {
	defer glog.Close()
	//	defer glog.Panicf()
	glog.Register(console.New())
	glog.SetLevel(glog.DebugLevel)

	a()
}

func a() {
	glog.Infoln("aaaaaaaa")
	glog.Infoln("aaaaaaaa1")
	glog.Infoln("aaaaaaaa2")
	glog.Infoln("aaaaaaaa3")
	glog.Infoln("aaaaaaaa4")
	panic("eeee")
}
