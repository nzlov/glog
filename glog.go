package glog

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
	"time"
)

var DefaultOption = &Option{
	NumCache: 20,
	NumStack: 99,
	Level:    LEVELALL,
}

type logger interface {
	Close()
	Event(Event)
	Option() *Option
}

var logs = []logger{}

func getCaller(i int) *FuncCall {
	pc, file, line, ok := runtime.Caller(i)
	if !ok {
		return nil
	}
	fs := strings.Split(file, "/")
	// fcs := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	return &FuncCall{
		File: fs[len(fs)-2] + "/" + fs[len(fs)-1],
		Line: line,
		// Func: fcs[len(fcs)-1],
		Func: runtime.FuncForPC(pc).Name(),
	}
}

func paincf(logger logger, s string, c int, data interface{}) {
	errstr := fmt.Sprintf("Panic : %v\nTraceback:\n", s)
	i := c
	for {
		pc, file, line, ok := runtime.Caller(i)
		if !ok || i > logger.Option().NumStack {
			break
		}
		errstr += fmt.Sprintf("\tstack: %d [file:%s][line:%d][func:%s]\n", i-c+1, file, line, runtime.FuncForPC(pc).Name())
		i++
	}
	logger.Event(Event{
		Level:   LEVELPANIC,
		Message: errstr,
		Time:    time.Now(),
		Data:    data,
	})
	exit()
}

func gol(logger logger, f interface{}, params ...interface{}) {
	fv := reflect.ValueOf(f)
	ft := reflect.TypeOf(f)
	if fv.Kind() == reflect.Func {
		if ft.NumIn() == len(params) {
			in := make([]reflect.Value, len(params))
			for i, p := range params {
				pv := reflect.ValueOf(p)
				if pv.Kind() == ft.In(i).Kind() {
					in[i] = pv
				} else {
					paincf(logger, fmt.Sprintf("params[%d] type %v don't is Func params[%d] type %v\n", i, pv.Kind(), i, ft.In(i).Kind()), 2, nil)
				}
			}
			defer func() {
				if err := recover(); err != nil {
					errstr := fmt.Sprintf("Panic : %v\ntraceback:\n", err)
					i := 4
					for {
						pc, file, line, ok := runtime.Caller(i)
						if !ok || i > logger.Option().NumStack {
							break
						}
						errstr += fmt.Sprintf("\tstack: %d %v [file:%s][line:%d][func:%s]\n", i-3, ok, file, line, runtime.FuncForPC(pc).Name())
						i++
					}
					logger.Event(Event{
						Level:   LEVELPANIC,
						Message: errstr,
						Time:    time.Now(),
						Data:    nil,
					})
					exit()
				}
			}()
			fv.Call(in)
		} else {
			paincf(logger, "params len don't == Func params", 2, nil)
		}
	} else {
		paincf(logger, "f don't is Func", 2, nil)
	}
}

func exit() {
	Close()
	os.Exit(1)
}

func Close() {
	for _, l := range logs {
		l.Close()
	}
}

var uuid int

func GenID() string {
	uuid++
	return fmt.Sprint("glog", uuid)
}
