package glog

type Listener interface {
	Name() string
	Notify() chan Event
	Start()
	Stop()
}
