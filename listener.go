package glog

type Listener interface {
	Name() string
	Event(e Event)
}
