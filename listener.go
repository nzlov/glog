package glog

type Listener interface {
	ID() string
	Notify() chan Event
	Start()
	Pause(bool)
	Stop()
	Option() Option
}
