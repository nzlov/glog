package glog

import (
	"fmt"
	"time"
)

type Field map[string]interface{}

func NewField() Field {
	return Field{}
}

func (self Field) Set(k string, v interface{}) Field {
	self[k] = v
	return self
}
func (self Field) Get(k string) (interface{}, bool) {
	v, b := self[k]
	return v, b
}
func (self Field) String() string {
	s := "\tShow Field:\n"
	for k, v := range self {
		s = s + fmt.Sprintf("\t\t%v=%+v\n", k, v)
	}
	return s
}
func (self Field) Panic(args ...interface{}) {
	paincf(fmt.Sprint(args...), 2, self)
}

func (self Field) Panicf(format string, args ...interface{}) {
	paincf(fmt.Sprintf(format, args...), 2, self)
}

func (self Field) Panicln(args ...interface{}) {
	paincf(fmt.Sprintln(args...), 2, self)
}
func (self Field) Error(args ...interface{}) {
	if level >= ErrorLevel {
		event(Event{
			Level:   ErrorLevel,
			Message: fmt.Sprint(args...),
			Time:    time.Now(),
			Data:    self,
		})
	}
}
func (self Field) Errorf(format string, args ...interface{}) {
	if level >= ErrorLevel {
		event(Event{
			Level:   ErrorLevel,
			Message: fmt.Sprintf(format, args...),
			Time:    time.Now(),
			Data:    self,
		})
	}
}
func (self Field) Errorln(args ...interface{}) {
	if level >= ErrorLevel {
		event(Event{
			Level:   ErrorLevel,
			Message: fmt.Sprintln(args...),
			Time:    time.Now(),
			Data:    self,
		})
	}
}

func (self Field) Warn(args ...interface{}) {
	if level >= WarnLevel {
		event(Event{
			Level:   WarnLevel,
			Message: fmt.Sprint(args...),
			Time:    time.Now(),
			Data:    self,
		})
	}
}
func (self Field) Warnf(format string, args ...interface{}) {
	if level >= WarnLevel {
		event(Event{
			Level:   WarnLevel,
			Message: fmt.Sprintf(format, args...),
			Time:    time.Now(),
			Data:    self,
		})
	}
}
func (self Field) Warnln(args ...interface{}) {
	if level >= WarnLevel {
		event(Event{
			Level:   WarnLevel,
			Message: fmt.Sprintln(args...),
			Time:    time.Now(),
			Data:    self,
		})
	}
}

func (self Field) Info(args ...interface{}) {
	if level >= InfoLevel {
		event(Event{
			Level:   InfoLevel,
			Message: fmt.Sprint(args...),
			Time:    time.Now(),
			Data:    self,
		})
	}
}
func (self Field) Infof(format string, args ...interface{}) {
	if level >= InfoLevel {
		event(Event{
			Level:   InfoLevel,
			Message: fmt.Sprintf(format, args...),
			Time:    time.Now(),
			Data:    self,
		})
	}
}
func (self Field) Infoln(args ...interface{}) {
	if level >= InfoLevel {
		event(Event{
			Level:   InfoLevel,
			Message: fmt.Sprintln(args...),
			Time:    time.Now(),
			Data:    self,
		})
	}
}

func (self Field) Debug(args ...interface{}) {
	if level >= DebugLevel {
		event(Event{
			Level:   DebugLevel,
			Message: fmt.Sprint(args...),
			Time:    time.Now(),
			Data:    self,
		})
	}
}
func (self Field) Debugf(format string, args ...interface{}) {
	if level >= DebugLevel {
		event(Event{
			Level:   DebugLevel,
			Message: fmt.Sprintf(format, args...),
			Time:    time.Now(),
			Data:    self,
		})
	}
}
func (self Field) Debugln(args ...interface{}) {
	if level >= DebugLevel {
		event(Event{
			Level:   DebugLevel,
			Message: fmt.Sprintln(args...),
			Time:    time.Now(),
			Data:    self,
		})
	}
}

type TagField struct {
	tag   string
	value map[string]interface{}
}

func NewTagField(tag string) *TagField {
	return &TagField{tag: tag, value: make(map[string]interface{})}
}

func (self *TagField) Set(k string, v interface{}) *TagField {
	self.value[k] = v
	return self
}
func (self *TagField) Get(k string) (interface{}, bool) {
	v, b := self.value[k]
	return v, b
}
func (self *TagField) GetTag() string {
	return self.tag
}
func (self *TagField) SetTag(tag string) *TagField {
	self.tag = tag
	return self
}
func (self *TagField) String() string {
	s := "\t[" + self.tag + "]Show Field:\n"
	for k, v := range self.value {
		s = s + fmt.Sprintf("\t\t%v=%+v\n", k, v)
	}
	return s
}
func (self *TagField) Panic(args ...interface{}) {
	paincf(fmt.Sprint(args...), 2, self)
}
func (self *TagField) Panicf(format string, args ...interface{}) {
	paincf(fmt.Sprintf(format, args...), 2, self)
}
func (self *TagField) Panicln(args ...interface{}) {
	paincf(fmt.Sprintln(args...), 2, self)
}
func (self *TagField) Error(args ...interface{}) {
	if level >= ErrorLevel {
		event(Event{
			Level:   ErrorLevel,
			Message: fmt.Sprint(args...),
			Time:    time.Now(),
			Data:    self,
		})
	}
}
func (self *TagField) Errorf(format string, args ...interface{}) {
	if level >= ErrorLevel {
		event(Event{
			Level:   ErrorLevel,
			Message: fmt.Sprintf(format, args...),
			Time:    time.Now(),
			Data:    self,
		})
	}
}
func (self *TagField) Errorln(args ...interface{}) {
	if level >= ErrorLevel {
		event(Event{
			Level:   ErrorLevel,
			Message: fmt.Sprintln(args...),
			Time:    time.Now(),
			Data:    self,
		})
	}
}

func (self *TagField) Warn(args ...interface{}) {
	if level >= WarnLevel {
		event(Event{
			Level:   WarnLevel,
			Message: fmt.Sprint(args...),
			Time:    time.Now(),
			Data:    self,
		})
	}
}
func (self *TagField) Warnf(format string, args ...interface{}) {
	if level >= WarnLevel {
		event(Event{
			Level:   WarnLevel,
			Message: fmt.Sprintf(format, args...),
			Time:    time.Now(),
			Data:    self,
		})
	}
}
func (self *TagField) Warnln(args ...interface{}) {
	if level >= WarnLevel {
		event(Event{
			Level:   WarnLevel,
			Message: fmt.Sprintln(args...),
			Time:    time.Now(),
			Data:    self,
		})
	}
}

func (self *TagField) Info(args ...interface{}) {
	if level >= InfoLevel {
		event(Event{
			Level:   InfoLevel,
			Message: fmt.Sprint(args...),
			Time:    time.Now(),
			Data:    self,
		})
	}
}
func (self *TagField) Infof(format string, args ...interface{}) {
	if level >= InfoLevel {
		event(Event{
			Level:   InfoLevel,
			Message: fmt.Sprintf(format, args...),
			Time:    time.Now(),
			Data:    self,
		})
	}
}
func (self *TagField) Infoln(args ...interface{}) {
	if level >= InfoLevel {
		event(Event{
			Level:   InfoLevel,
			Message: fmt.Sprintln(args...),
			Time:    time.Now(),
			Data:    self,
		})
	}
}

func (self *TagField) Debug(args ...interface{}) {
	if level >= DebugLevel {
		event(Event{
			Level:   DebugLevel,
			Message: fmt.Sprint(args...),
			Time:    time.Now(),
			Data:    self,
		})
	}
}
func (self *TagField) Debugf(format string, args ...interface{}) {
	if level >= DebugLevel {
		event(Event{
			Level:   DebugLevel,
			Message: fmt.Sprintf(format, args...),
			Time:    time.Now(),
			Data:    self,
		})
	}
}
func (self *TagField) Debugln(args ...interface{}) {
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
	PanicLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
)

func (level Level) String() string {
	switch level {
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case WarnLevel:
		return "WARNING"
	case ErrorLevel:
		return "ERROR"
	case PanicLevel:
		return "PANIC"
	}

	return "UNKNOWN"
}

func ParseLevel(lvl string) Level {
	switch lvl {
	case "PANIC", "panic":
		return PanicLevel
	case "ERROR", "error":
		return ErrorLevel
	case "WARNING", "warning":
		return WarnLevel
	case "INFO", "info":
		return InfoLevel
	case "DEBUG", "debug":
		return DebugLevel
	default:
		return UnknownLevel
	}
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
