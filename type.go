package glog

import (
	"fmt"
	"time"
)

type Field struct {
	logger logger
	value  map[string]interface{}
}

func newField(logger logger) Field {
	return Field{
		logger: logger,
		value:  make(map[string]interface{}),
	}
}

func (self Field) S(k string, v interface{}) Field {
	self.value[k] = v
	return self
}
func (self Field) G(k string) (interface{}, bool) {
	v, b := self.value[k]
	return v, b
}
func (self Field) String() string {
	s := "\tShow Field:\n"
	for k, v := range self.value {
		s = s + fmt.Sprintf("\t\t%v=%+v\n", k, v)
	}
	return s
}
func (self Field) Panic(args ...interface{}) {
	paincf(self.logger, fmt.Sprint(args...), 2, self)
}

func (self Field) Panicf(format string, args ...interface{}) {
	paincf(self.logger, fmt.Sprintf(format, args...), 2, self)
}

func (self Field) Panicln(args ...interface{}) {
	paincf(self.logger, fmt.Sprintln(args...), 2, self)
}
func (self Field) Error(args ...interface{}) {
	if self.logger.Option().Level&LEVELERROR == LEVELERROR {
		self.logger.Event(Event{
			Level:   LEVELERROR,
			Message: fmt.Sprint(args...),
			Time:    time.Now(),
			Data:    self,
		})
	}
}
func (self Field) Errorf(format string, args ...interface{}) {
	if self.logger.Option().Level&LEVELERROR == LEVELERROR {
		self.logger.Event(Event{
			Level:   LEVELERROR,
			Message: fmt.Sprintf(format, args...),
			Time:    time.Now(),
			Data:    self,
		})
	}
}
func (self Field) Errorln(args ...interface{}) {
	if self.logger.Option().Level&LEVELERROR == LEVELERROR {
		self.logger.Event(Event{
			Level:   LEVELERROR,
			Message: fmt.Sprintln(args...),
			Time:    time.Now(),
			Data:    self,
		})
	}
}

func (self Field) Warn(args ...interface{}) {
	if self.logger.Option().Level&LEVELWARN == LEVELWARN {
		self.logger.Event(Event{
			Level:   LEVELWARN,
			Message: fmt.Sprint(args...),
			Time:    time.Now(),
			Data:    self,
		})
	}
}
func (self Field) Warnf(format string, args ...interface{}) {
	if self.logger.Option().Level&LEVELWARN == LEVELWARN {
		self.logger.Event(Event{
			Level:   LEVELWARN,
			Message: fmt.Sprintf(format, args...),
			Time:    time.Now(),
			Data:    self,
		})
	}
}
func (self Field) Warnln(args ...interface{}) {
	if self.logger.Option().Level&LEVELWARN == LEVELWARN {
		self.logger.Event(Event{
			Level:   LEVELWARN,
			Message: fmt.Sprintln(args...),
			Time:    time.Now(),
			Data:    self,
		})
	}
}

func (self Field) Info(args ...interface{}) {
	if self.logger.Option().Level&LEVELINFO == LEVELINFO {
		self.logger.Event(Event{
			Level:   LEVELINFO,
			Message: fmt.Sprint(args...),
			Time:    time.Now(),
			Data:    self,
		})
	}
}
func (self Field) Infof(format string, args ...interface{}) {
	if self.logger.Option().Level&LEVELINFO == LEVELINFO {
		self.logger.Event(Event{
			Level:   LEVELINFO,
			Message: fmt.Sprintf(format, args...),
			Time:    time.Now(),
			Data:    self,
		})
	}
}
func (self Field) Infoln(args ...interface{}) {
	if self.logger.Option().Level&LEVELINFO == LEVELINFO {
		self.logger.Event(Event{
			Level:   LEVELINFO,
			Message: fmt.Sprintln(args...),
			Time:    time.Now(),
			Data:    self,
		})
	}
}

func (self Field) Debug(args ...interface{}) {
	if self.logger.Option().Level&LEVELDEBUG == LEVELDEBUG {
		self.logger.Event(Event{
			Level:   LEVELDEBUG,
			Message: fmt.Sprint(args...),
			Time:    time.Now(),
			Data:    self,
		})
	}
}
func (self Field) Debugf(format string, args ...interface{}) {
	if self.logger.Option().Level&LEVELDEBUG == LEVELDEBUG {
		self.logger.Event(Event{
			Level:   LEVELDEBUG,
			Message: fmt.Sprintf(format, args...),
			Time:    time.Now(),
			Data:    self,
		})
	}
}
func (self Field) Debugln(args ...interface{}) {
	if self.logger.Option().Level&LEVELDEBUG == LEVELDEBUG {
		self.logger.Event(Event{
			Level:   LEVELDEBUG,
			Message: fmt.Sprintln(args...),
			Time:    time.Now(),
			Data:    self,
		})
	}
}

type Level uint8

const (
	LEVELUNKNOWN Level = 0
	LEVELPANIC   Level = 1 << iota
	LEVELERROR
	LEVELWARN
	LEVELINFO
	LEVELDEBUG
	LEVELALL Level = LEVELPANIC | LEVELERROR | LEVELWARN | LEVELINFO | LEVELDEBUG
)

func (level Level) String() string {
	switch level {
	case LEVELDEBUG:
		return "D"
	case LEVELINFO:
		return "I"
	case LEVELWARN:
		return "W"
	case LEVELERROR:
		return "E"
	case LEVELPANIC:
		return "P"
	}

	return "U"
}

type Event struct {
	Level    Level
	Message  string
	Time     time.Time
	Data     interface{}
	FuncCall *FuncCall
}

type FuncCall struct {
	File string
	Line int
	Func string
}

func (fc FuncCall) String() string {

	return fmt.Sprintf("[%s:%d][%s]", fc.File, fc.Line, fc.Func)
}
