package email

import (
	"fmt"

	"strings"

	"github.com/go-gomail/gomail"
	"github.com/nzlov/glog"
	"github.com/nzlov/glog/listener"
)

type Email struct {
	*listener.BaseListener

	l     glog.Level
	mail  *gomail.Dialer
	from  string
	to    []string
	title string
}

func New(l glog.Level, smtp string, port int, user, pwd, from string, to []string, title string) (*Email, error) {
	return NewWithOption(l, smtp, port, user, pwd, from, to, title, glog.DefaultOption)
}
func NewWithOption(l glog.Level, smtp string, port int, user, pwd, from string, to []string, title string, o *glog.Option) (*Email, error) {

	mail := gomail.NewDialer(smtp, port, user, pwd)
	_, err := mail.Dial()
	if err != nil {
		return nil, err
	}

	c := &Email{}
	c.BaseListener = listener.NewBaseListener(c, o)
	c.l = l
	c.mail = mail
	c.from = from
	c.to = to
	c.title = title
	return c, nil
}

func (self *Email) Event(e glog.Event) {
	if e.Level >= self.l {
		msg := gomail.NewMessage()
		msg.SetHeader("From", self.from)
		msg.SetHeader("To", self.to...)
		msg.SetHeader("Subject", self.title)
		if e.FuncCall != nil {
			if e.Data == nil {
				msg.SetBody("text/html", fmt.Sprintf("[%s][%s]%s %s", e.Level, e.Time.Format("2006-01-02 15:04:05"), e.FuncCall, strings.Replace(e.Message, "\n", "<br />", -1)))
			} else {
				msg.SetBody("text/html", fmt.Sprintf("[%s][%s]%s %s<br /> %s", e.Level, e.Time.Format("2006-01-02 15:04:05"), e.FuncCall, strings.Replace(e.Message, "\n", "<br />", -1), strings.Replace(fmt.Sprint(e.Data), "\n", "<br />", -1)))
			}
		} else {
			if e.Data == nil {
				msg.SetBody("text/html", fmt.Sprintf("[%s][%s] %s", e.Level, e.Time.Format("2006-01-02 15:04:05"), strings.Replace(e.Message, "\n", "<br />", -1)))
			} else {
				msg.SetBody("text/html", fmt.Sprintf("[%s][%s] %s<br /> %s", e.Level, e.Time.Format("2006-01-02 15:04:05"), strings.Replace(e.Message, "\n", "<br />", -1), strings.Replace(fmt.Sprint(e.Data), "\n", "<br />", -1)))
			}
		}
		if err := self.mail.DialAndSend(msg); err != nil {
			panic(err)
		}

	}

}
func (self *Email) Close() {
}
