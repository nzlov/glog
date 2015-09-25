package glog

import (
	"fmt"
	"time"
)

type Filed map[string]interface{}

func NewFiled() Filed {
	return Filed{}
}

func (self Filed) Set(k string, v interface{}) {
	self[k] = v
}
func (self Filed) Get(k string) (interface{}, bool) {
	v, b := self[k]
	return v, b
}
func (self Filed) String() string {
	s := ""
	for k, v := range self {
		s = s + fmt.Sprintf(" %s=%s", k, v)
	}
	s = s + ""
	return s
}

func (self Filed) Error(args ...interface{}) {
	if level >= ErrorLevel {
		event(Event{
			Level:   ErrorLevel,
			Message: fmt.Sprint(args...),
			Time:    time.Now(),
			Data:    self,
		})
	}
}
func (self Filed) Errorf(format string, args ...interface{}) {
	if level >= ErrorLevel {
		event(Event{
			Level:   ErrorLevel,
			Message: fmt.Sprintf(format, args...),
			Time:    time.Now(),
			Data:    self,
		})
	}
}
func (self Filed) Errorln(args ...interface{}) {
	if level >= ErrorLevel {
		event(Event{
			Level:   ErrorLevel,
			Message: fmt.Sprintln(args...),
			Time:    time.Now(),
			Data:    self,
		})
	}
}

func (self Filed) Warn(args ...interface{}) {
	if level >= WarnLevel {
		event(Event{
			Level:   WarnLevel,
			Message: fmt.Sprint(args...),
			Time:    time.Now(),
			Data:    self,
		})
	}
}
func (self Filed) Warnf(format string, args ...interface{}) {
	if level >= WarnLevel {
		event(Event{
			Level:   WarnLevel,
			Message: fmt.Sprintf(format, args...),
			Time:    time.Now(),
			Data:    self,
		})
	}
}
func (self Filed) Warnln(args ...interface{}) {
	if level >= WarnLevel {
		event(Event{
			Level:   WarnLevel,
			Message: fmt.Sprintln(args...),
			Time:    time.Now(),
			Data:    self,
		})
	}
}

func (self Filed) Info(args ...interface{}) {
	if level >= InfoLevel {
		event(Event{
			Level:   InfoLevel,
			Message: fmt.Sprint(args...),
			Time:    time.Now(),
			Data:    self,
		})
	}
}
func (self Filed) Infof(format string, args ...interface{}) {
	if level >= InfoLevel {
		event(Event{
			Level:   InfoLevel,
			Message: fmt.Sprintf(format, args...),
			Time:    time.Now(),
			Data:    self,
		})
	}
}
func (self Filed) Infoln(args ...interface{}) {
	if level >= InfoLevel {
		event(Event{
			Level:   InfoLevel,
			Message: fmt.Sprintln(args...),
			Time:    time.Now(),
			Data:    self,
		})
	}
}

func (self Filed) Debug(args ...interface{}) {
	if level >= DebugLevel {
		event(Event{
			Level:   DebugLevel,
			Message: fmt.Sprint(args...),
			Time:    time.Now(),
			Data:    self,
		})
	}
}
func (self Filed) Debugf(format string, args ...interface{}) {
	if level >= DebugLevel {
		event(Event{
			Level:   DebugLevel,
			Message: fmt.Sprintf(format, args...),
			Time:    time.Now(),
			Data:    self,
		})
	}
}
func (self Filed) Debugln(args ...interface{}) {
	if level >= DebugLevel {
		event(Event{
			Level:   DebugLevel,
			Message: fmt.Sprintln(args...),
			Time:    time.Now(),
			Data:    self,
		})
	}
}

type Level uint8

const (
	UnknownLevel Level = iota
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
)

func (level Level) String() string {
	switch level {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warning"
	case ErrorLevel:
		return "error"
	}

	return "unknown"
}

func ParseLevel(lvl string) (Level, error) {
	switch lvl {
	case "error":
		return ErrorLevel, nil
	case "warn", "warning":
		return WarnLevel, nil
	case "info":
		return InfoLevel, nil
	case "debug":
		return DebugLevel, nil
	}

	var l Level
	return l, fmt.Errorf("not a valid logrus Level: %q", lvl)
}

type Event struct {
	Level   Level
	Message string
	Time    time.Time
	Data    Filed
}
