package observer

type Subject interface {
	Register(observer Observer)
	Remove(observer Observer)
	Notify(event interface{})
}
