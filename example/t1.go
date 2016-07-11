package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/nzlov/glog"
	"github.com/nzlov/glog/listener/console"
	"github.com/nzlov/glog/listener/file"
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
	glog.SetLevel(glog.DebugLevel)
	glog.Register(console.New())
	f, err := file.New("l.log")
	if err != nil {
		panic(err)
	}
	glog.Register(f)

	a()
}
func a() {
	go func() {
		i := 1
		for {
			glog.Infoln(i)
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
