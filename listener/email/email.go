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

	mail  *gomail.Dialer
	from  string
	to    []string
	title string
}

func New(smtp string, port int, user, pwd, from string, to []string, title string) (*Email, error) {

	mail := gomail.NewDialer(smtp, port, user, pwd)
	_, err := mail.Dial()
	if err != nil {
		return nil, err
	}

	c := &Email{}
	c.BaseListener = listener.NewBaseListener(c)
	c.mail = mail
	c.from = from
	c.to = to
	c.title = title
	return c, nil
}
func (self *Email) Name() string {
	return "Email"
}

func (self *Email) Event(e glog.Event) {
	if e.Level == glog.PanicLevel {
		msg := gomail.NewMessage()
		msg.SetHeader("From", self.from)
		msg.SetHeader("To", self.to...)
		msg.SetHeader("Subject", self.title)
		if e.Data == nil {
			msg.SetBody("text/html", fmt.Sprintf("[%s][%s] %s", e.Level, e.Time.Format("2006-01-02 15:04:05"), strings.Replace(e.Message, "\n", "<br />", -1)))
		} else {
			msg.SetBody("text/html", fmt.Sprintf("[%s][%s] %s<br /> %s", e.Level, e.Time.Format("2006-01-02 15:04:05"), strings.Replace(e.Message, "\n", "<br />", -1), strings.Replace(fmt.Sprint(e.Data), "\n", "<br />", -1)))
		}
		if err := self.mail.DialAndSend(msg); err != nil {
			panic(err)
		}
	}

}
func (self *Email) Close() {
}
