package colorconsole

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/nzlov/glog"
	"github.com/nzlov/glog/listener"
)

type ShowType int

const (
	ShowTypeDefault ShowType = iota
	ShowTypeFore
)

var (
	PanicLevelColor = color.New(color.FgRed)
	ErrorLevelColor = color.New(color.FgHiRed)
	DebugLevelColor = color.New(color.FgCyan)
	WarnLevelColor  = color.New(color.FgYellow)
	InfoLevelColor  = color.New(color.FgGreen)
)

type ColorConsole struct {
	*listener.BaseListener

	ty ShowType
	cl *color.Color
}

func New(s ShowType) *ColorConsole {
	return NewWithOption(s, glog.DefaultOption)
}

func NewWithOption(s ShowType, o *glog.Option) *ColorConsole {
	c := &ColorConsole{
		ty: s,
	}
	c.BaseListener = listener.NewBaseListener(c, o)
	return c
}

func (self *ColorConsole) Event(e glog.Event) {
	switch self.ty {
	case ShowTypeFore:
		self.typefg(e)
	default:
		self.typedef(e)
	}
}

func (self *ColorConsole) typefg(e glog.Event) {
	switch e.Level {
	case glog.LEVELPANIC:
		self.cl = PanicLevelColor
	case glog.LEVELERROR:
		self.cl = ErrorLevelColor
	case glog.LEVELWARN:
		self.cl = WarnLevelColor
	case glog.LEVELINFO:
		self.cl = InfoLevelColor
	}
	if e.FuncCall != nil {
		self.cl.Printf("[%s][%s]%s %s", e.Level, e.Time.Format("2006-01-02 15:04:05"), e.FuncCall, e.Message)
	} else {
		self.cl.Printf("[%s][%s] %s", e.Level, e.Time.Format("2006-01-02 15:04:05"), e.Message)
	}
	if e.Data != nil {
		self.cl.Printf("%s", e.Data)
	}
}

func (self *ColorConsole) typedef(e glog.Event) {
	fmt.Print("[")
	switch e.Level {
	case glog.LEVELPANIC:
		PanicLevelColor.Print(e.Level)
	case glog.LEVELERROR:
		ErrorLevelColor.Print(e.Level)
	case glog.LEVELWARN:
		WarnLevelColor.Print(e.Level)
	case glog.LEVELINFO:
		InfoLevelColor.Print(e.Level)
	}
	fmt.Print("]")
	if e.FuncCall != nil {
		fmt.Printf("[%s]%s %s", e.Time.Format("2006-01-02 15:04:05"), e.FuncCall, e.Message)
	} else {
		fmt.Printf("[%s] %s", e.Time.Format("2006-01-02 15:04:05"), e.Message)
	}
	if e.Data != nil {
		fmt.Printf("%s", e.Data)
	}
}
func (self *ColorConsole) Close() {}
