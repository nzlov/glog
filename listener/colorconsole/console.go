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
	c := &ColorConsole{
		ty: s,
	}
	c.BaseListener = listener.NewBaseListener(c)
	return c
}
func (self *ColorConsole) Name() string {
	return "ColorConsole"
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
	case glog.PanicLevel:
		self.cl = PanicLevelColor
	case glog.ErrorLevel:
		self.cl = ErrorLevelColor
	case glog.DebugLevel:
		self.cl = DebugLevelColor
	case glog.WarnLevel:
		self.cl = WarnLevelColor
	case glog.InfoLevel:
		self.cl = InfoLevelColor
	}
	self.cl.Printf("[%s][%s] %s", e.Level, e.Time.Format("2006-01-02 15:04:05"), e.Message)
	if e.Data != nil {
		self.cl.Printf("%s", e.Data)
	}
}

func (self *ColorConsole) typedef(e glog.Event) {
	fmt.Print("[")
	switch e.Level {
	case glog.PanicLevel:
		PanicLevelColor.Print(e.Level)
	case glog.ErrorLevel:
		ErrorLevelColor.Print(e.Level)
	case glog.DebugLevel:
		DebugLevelColor.Print(e.Level)
	case glog.WarnLevel:
		WarnLevelColor.Print(e.Level)
	case glog.InfoLevel:
		InfoLevelColor.Print(e.Level)
	}
	fmt.Print("]")
	fmt.Printf("[%s] %s", e.Time.Format("2006-01-02 15:04:05"), e.Message)
	if e.Data != nil {
		fmt.Printf("%s", e.Data)
	}
}
func (self *ColorConsole) Close() {}
