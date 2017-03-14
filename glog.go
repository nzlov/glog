package glog

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
	"sync"
	"time"
)

var QuitWait sync.WaitGroup

var MAXSTACK int = 99
var LOGCHANSIZE int = 20

var listeners map[string]Listener
var level Level
var events chan Event
var isRunning bool
var done chan bool

func init() {
	listeners = make(map[string]Listener)
	level = InfoLevel
	events = make(chan Event, LOGCHANSIZE)
	done = make(chan bool)
	isRunning = true
	go le()
}

func Register(l Listener) {
	l.Start()
	listeners[l.Name()] = l
	QuitWait.Add(1)
}
func SetLevel(l Level) {
	level = l
}

func event(e Event) {
	if isRunning {
		events <- e
	}
}

func le() {
	for e := range events {
		for _, l := range listeners {
			l.Notify() <- e
		}
	}
	done <- true
}

func Close() {
	if !isRunning {
		return
	}
	if err := recover(); err != nil {
		errstr := fmt.Sprintf("Runtime error:%v\ntraceback:\n", err)
		i := 3
		for {
			pc, file, line, ok := runtime.Caller(i)
			if !ok || i > MAXSTACK {
				break
			}
			errstr += fmt.Sprintf("\tstack: %d %v [file:%s][line:%d][func:%s]\n", i-2, ok, file, line, runtime.FuncForPC(pc).Name())
			i++
		}
		event(Event{
			Level:   PanicLevel,
			Message: errstr,
			Time:    time.Now(),
			Data:    nil,
		})
	}
	isRunning = false
	close(events)
	<-done
	for _, l := range listeners {
		l.Stop()
	}
	QuitWait.Wait()
}

func exit() {
	Close()
	os.Exit(1)
}

func Panic(args ...interface{}) {
	paincf(fmt.Sprint(args...), 2, nil)
}
func Panicf(format string, args ...interface{}) {
	paincf(fmt.Sprintf(format, args...), 2, nil)
}

func Panicln(args ...interface{}) {
	paincf(fmt.Sprintln(args...), 2, nil)
}
func paincf(s string, c int, data interface{}) {
	errstr := fmt.Sprintf("Runtime error:%v\nTraceback:\n", s)
	i := c
	for {
		pc, file, line, ok := runtime.Caller(i)
		if !ok || i > MAXSTACK {
			break
		}
		errstr += fmt.Sprintf("\tstack: %d [file:%s][line:%d][func:%s]\n", i-c, file, line, runtime.FuncForPC(pc).Name())
		i++
	}
	event(Event{
		Level:   PanicLevel,
		Message: errstr,
		Time:    time.Now(),
		Data:    data,
	})
	exit()
}

func Go(f interface{}, params ...interface{}) {
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
					Panicf("params[%d] type %v don't is Func params[%d] type %v\n", i, pv.Kind(), i, ft.In(i).Kind())
				}
			}
			defer func() {
				if err := recover(); err != nil {
					errstr := fmt.Sprintf("Runtime error:%v\ntraceback:\n", err)
					i := 3
					for {
						pc, file, line, ok := runtime.Caller(i)
						if !ok || i > MAXSTACK {
							break
						}
						errstr += fmt.Sprintf("\tstack: %d %v [file:%s][line:%d][func:%s]\n", i-2, ok, file, line, runtime.FuncForPC(pc).Name())
						i++
					}
					event(Event{
						Level:   PanicLevel,
						Message: errstr,
						Time:    time.Now(),
						Data:    nil,
					})
					exit()
				}
			}()
			fv.Call(in)
		} else {
			Panicln("params len don't == Func params")
		}
	} else {
		Panicln("f don't is Func")
	}
}

func Error(args ...interface{}) {
	if level >= ErrorLevel {
		event(Event{
			Level:   ErrorLevel,
			Message: fmt.Sprint(args...),
			Time:    time.Now(),
			Data:    nil,
		})
	}
}
func Errorf(format string, args ...interface{}) {
	if level >= ErrorLevel {
		event(Event{
			Level:   ErrorLevel,
			Message: fmt.Sprintf(format, args...),
			Time:    time.Now(),
			Data:    nil,
		})
	}
}
func Errorln(args ...interface{}) {
	if level >= ErrorLevel {
		event(Event{
			Level:   ErrorLevel,
			Message: fmt.Sprintln(args...),
			Time:    time.Now(),
			Data:    nil,
		})
	}
}

func Warn(args ...interface{}) {
	if level >= WarnLevel {
		event(Event{
			Level:   WarnLevel,
			Message: fmt.Sprint(args...),
			Time:    time.Now(),
			Data:    nil,
		})
	}
}
func Warnf(format string, args ...interface{}) {
	if level >= WarnLevel {
		event(Event{
			Level:   WarnLevel,
			Message: fmt.Sprintf(format, args...),
			Time:    time.Now(),
			Data:    nil,
		})
	}
}
func Warnln(args ...interface{}) {
	if level >= WarnLevel {
		event(Event{
			Level:   WarnLevel,
			Message: fmt.Sprintln(args...),
			Time:    time.Now(),
			Data:    nil,
		})
	}
}

func Info(args ...interface{}) {
	if level >= InfoLevel {
		event(Event{
			Level:   InfoLevel,
			Message: fmt.Sprint(args...),
			Time:    time.Now(),
			Data:    nil,
		})
	}
}
func Infof(format string, args ...interface{}) {
	if level >= InfoLevel {
		event(Event{
			Level:   level,
			Message: fmt.Sprintf(format, args...),
			Time:    time.Now(),
			Data:    nil,
		})
	}
}
func Infoln(args ...interface{}) {
	if level >= InfoLevel {
		event(Event{
			Level:   InfoLevel,
			Message: fmt.Sprintln(args...),
			Time:    time.Now(),
			Data:    nil,
		})
	}
}

func Debug(args ...interface{}) {
	if level >= DebugLevel {
		event(Event{
			Level:   level,
			Message: fmt.Sprint(args...),
			Time:    time.Now(),
			Data:    nil,
		})
	}
}
func Debugf(format string, args ...interface{}) {
	if level >= DebugLevel {
		event(Event{
			Level:   level,
			Message: fmt.Sprintf(format, args...),
			Time:    time.Now(),
			Data:    nil,
		})
	}
}
func Debugln(args ...interface{}) {
	if level >= DebugLevel {
		event(Event{
			Level:   level,
			Message: fmt.Sprintln(args...),
			Time:    time.Now(),
			Data:    nil,
		})
	}
}
