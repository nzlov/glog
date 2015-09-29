package glog

import (
	"fmt"
	"time"
)

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
	for isRunning {
		select {
		case e := <-events:
			for _, l := range listeners {
				l.Event(e)
			}
		case <-time.After(time.Millisecond * 500):

		}
	}

	for {
		select {
		case e := <-events:
			for _, l := range listeners {
				l.Event(e)
			}
		case <-time.After(time.Millisecond * 500):
			done <- true
			return

		}
	}
}

func Close() {
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
