package observer

type Observer interface {
	Update(event interface{})
}
