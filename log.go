package glog

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type Option struct {
	NumCache int
	NumStack int
	Fc       bool
	Level    Level
}

type Logger struct {
	option    *Option
	listeners map[Level]map[string]Listener
	running   bool

	events chan Event
	done   chan struct{}
	lock   sync.Mutex
}

func New() *Logger {
	return NewWithOption(DefaultOption)
}
func NewWithOption(o *Option) *Logger {
	l := &Logger{
		option:    o,
		listeners: make(map[Level]map[string]Listener),
		events:    make(chan Event, o.NumCache),
		done:      make(chan struct{}),
		running:   true,
		lock:      sync.Mutex{},
	}
	l.listeners[LEVELPANIC] = make(map[string]Listener, 0)
	l.listeners[LEVELERROR] = make(map[string]Listener, 0)
	l.listeners[LEVELWARN] = make(map[string]Listener, 0)
	l.listeners[LEVELINFO] = make(map[string]Listener, 0)
	go l.ev()
	logs = append(logs, l)
	return l
}

func (logger *Logger) AddListener(l ...Listener) {
	for _, ler := range l {
		b := false
		for k, _ := range logger.listeners {
			if k == ler.Option().Level&k {
				logger.listeners[k][ler.ID()] = ler
				b = true
			}
		}
		if b {
			ler.Start()
		}
	}
}

func (logger *Logger) UpdateListener(l ...Listener) {
	for _, ler := range l {
		delete(logger.listeners[LEVELPANIC], ler.ID())
		delete(logger.listeners[LEVELERROR], ler.ID())
		delete(logger.listeners[LEVELWARN], ler.ID())
		delete(logger.listeners[LEVELINFO], ler.ID())
		b := false
		for k, _ := range logger.listeners {
			if k == ler.Option().Level&k {
				logger.listeners[k][ler.ID()] = ler
				b = true
			}
		}
		if b {
			ler.Start()
		}
	}
}

func (logger *Logger) ev() {
	for e := range logger.events {
		if ls, ok := logger.listeners[e.Level]; ok {
			for _, l := range ls {
				l.Notify() <- e
			}
		}
	}
	logger.done <- struct{}{}
}
func (logger *Logger) Close() {
	if !logger.running {
		return
	}
	logger.lock.Lock()
	logger.running = false
	logger.lock.Unlock()
	if err := recover(); err != nil {
		errstr := fmt.Sprintf("Runtime error:%v\ntraceback:\n", err)
		i := 4
		for {
			pc, file, line, ok := runtime.Caller(i)
			if !ok || i > logger.option.NumStack {
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
	}
	close(logger.events)
	<-logger.done

	for _, ls := range logger.listeners {
		for _, l := range ls {
			l.Stop()
		}
	}
}

func (logger *Logger) Go(f interface{}, params ...interface{}) {
	gol(logger, f, params...)
}

func (logger *Logger) Event(e Event) {
	logger.lock.Lock()
	if logger.running {
		if logger.option.Fc {
			e.FuncCall = getCaller(3)
		}
		logger.events <- e
	}
	logger.lock.Unlock()
}

func (logger *Logger) Pause(b bool) {
	for _, ls := range logger.listeners {
		for _, l := range ls {
			l.Pause(b)
		}
	}
}
func (logger *Logger) Panic(args ...interface{}) {
	paincf(logger, fmt.Sprint(args...), 2, nil)
}
func (logger *Logger) Panicf(format string, args ...interface{}) {
	paincf(logger, fmt.Sprintf(format, args...), 2, nil)
}

func (logger *Logger) Panicln(args ...interface{}) {
	paincf(logger, fmt.Sprintln(args...), 2, nil)
}
func (logger *Logger) Error(args ...interface{}) {
	if logger.option.Level&LEVELERROR == LEVELERROR {
		logger.Event(Event{
			Level:   LEVELERROR,
			Message: fmt.Sprint(args...),
			Time:    time.Now(),
			Data:    nil,
		})
	}
}
func (logger *Logger) Errorf(format string, args ...interface{}) {
	if logger.option.Level&LEVELERROR == LEVELERROR {
		logger.Event(Event{
			Level:   LEVELERROR,
			Message: fmt.Sprintf(format, args...),
			Time:    time.Now(),
			Data:    nil,
		})
	}
}
func (logger *Logger) Errorln(args ...interface{}) {
	if logger.option.Level&LEVELERROR == LEVELERROR {
		logger.Event(Event{
			Level:   LEVELERROR,
			Message: fmt.Sprintln(args...),
			Time:    time.Now(),
			Data:    nil,
		})
	}
}

func (logger *Logger) Warn(args ...interface{}) {
	if logger.option.Level&LEVELWARN == LEVELWARN {
		logger.Event(Event{
			Level:   LEVELWARN,
			Message: fmt.Sprint(args...),
			Time:    time.Now(),
			Data:    nil,
		})
	}
}
func (logger *Logger) Warnf(format string, args ...interface{}) {
	if logger.option.Level&LEVELWARN == LEVELWARN {
		logger.Event(Event{
			Level:   LEVELWARN,
			Message: fmt.Sprintf(format, args...),
			Time:    time.Now(),
			Data:    nil,
		})
	}
}
func (logger *Logger) Warnln(args ...interface{}) {
	if logger.option.Level&LEVELWARN == LEVELWARN {
		logger.Event(Event{
			Level:   LEVELWARN,
			Message: fmt.Sprintln(args...),
			Time:    time.Now(),
			Data:    nil,
		})
	}
}

func (logger *Logger) Info(args ...interface{}) {
	if logger.option.Level&LEVELINFO == LEVELINFO {
		logger.Event(Event{
			Level:   LEVELINFO,
			Message: fmt.Sprint(args...),
			Time:    time.Now(),
			Data:    nil,
		})
	}
}
func (logger *Logger) Infof(format string, args ...interface{}) {
	if logger.option.Level&LEVELINFO == LEVELINFO {
		logger.Event(Event{
			Level:   LEVELINFO,
			Message: fmt.Sprintf(format, args...),
			Time:    time.Now(),
			Data:    nil,
		})
	}
}
func (logger *Logger) Infoln(args ...interface{}) {
	if logger.option.Level&LEVELINFO == LEVELINFO {
		logger.Event(Event{
			Level:   LEVELINFO,
			Message: fmt.Sprintln(args...),
			Time:    time.Now(),
			Data:    nil,
		})
	}
}

func (logger *Logger) Option() *Option {
	return logger.option
}
func (logger *Logger) NewField() Field {
	return newField(logger)
}

func (logger *Logger) S(k string, v interface{}) Field {
	return newField(logger).S(k, v)
}
