package glog

import (
	"fmt"
	"runtime"
	"time"
)

var MAXSTACK int = 10

var listeners map[string]Listener
var level Level
var events chan Event
var isRunning bool
var done chan bool

func init() {
	listeners = make(map[string]Listener)
	level = InfoLevel
	events = make(chan Event, 99)
	done = make(chan bool)
	isRunning = true
	go le()
}

func Register(l Listener) {
	listeners[l.Name()] = l
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
	d := time.Millisecond * 10
	t := time.NewTimer(d)
	defer t.Stop()
	for isRunning {
		t.Reset(d)
		select {
		case e := <-events:
			for _, l := range listeners {
				l.Event(e)
			}
		case <-t.C:

		}
	}

	for {
		t.Reset(time.Millisecond)
		select {
		case e := <-events:
			for _, l := range listeners {
				l.Event(e)
			}
		case <-t.C:
			done <- true
			return

		}
	}
}

func Close() {
	if !isRunning {
		return
	}
	if err := recover(); err != nil {
		errstr := fmt.Sprintf("Runtime error:%v\ntraceback:\n", err)
		i := 1
		for {
			pc, file, line, ok := runtime.Caller(i)
			if !ok || i > MAXSTACK {
				break
			}
			errstr += fmt.Sprintf("\tstack: %d %v [file:%s][line:%d][func:%s]\n", i, ok, file, line, runtime.FuncForPC(pc).Name())
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
	<-done
	close(events)

	for _, l := range listeners {
		l.Close()
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
