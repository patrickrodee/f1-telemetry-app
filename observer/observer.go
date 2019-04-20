package observer

type observer struct {
	input      chan interface{}
	register   chan chan<- interface{}
	deregister chan chan<- interface{}
	observers  map[chan<- interface{}]bool
}

// Observer provides methods for observing data changes
type Observer interface {
	Register(ch chan<- interface{})
	Deregister(ch chan<- interface{})
	Send(v interface{})
	Close() error
}

// NewObserver returns a new Observer instance
func NewObserver(buflen int) Observer {
	o := &observer{
		input:      make(chan interface{}, buflen),
		register:   make(chan chan<- interface{}),
		deregister: make(chan chan<- interface{}),
		observers:  make(map[chan<- interface{}]bool),
	}
	go o.run()
	return o
}

func (o *observer) Register(ch chan<- interface{}) {
	o.register <- ch
}

func (o *observer) Deregister(ch chan<- interface{}) {
	o.deregister <- ch
}

func (o *observer) Send(v interface{}) {
	if o != nil {
		o.input <- v
	}
}

func (o *observer) Close() error {
	close(o.register)
	return nil
}

func (o *observer) run() {
	for {
		select {
		case v := <-o.input:
			o.deliver(v)
		case ch, ok := <-o.register:
			if !ok {
				return
			}
			o.observers[ch] = true
		case ch := <-o.deregister:
			close(ch)
			delete(o.observers, ch)
		}
	}
}

func (o observer) deliver(v interface{}) {
	for ch := range o.observers {
		ch <- v
	}
}
