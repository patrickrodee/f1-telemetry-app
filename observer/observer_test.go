package observer

import (
	"sync"
	"testing"
)

type stuff struct {
	V int
}

func newStuff(v int) *stuff {
	return &stuff{v}
}

func TestObserver(t *testing.T) {
	wg := sync.WaitGroup{}

	o := NewObserver(10)
	defer o.Close()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		cch := make(chan interface{})
		o.Register(cch)

		go func() {
			defer wg.Done()
			defer o.Deregister(cch)
			<-cch
		}()
	}

	o.Send(newStuff(10))
	wg.Wait()
}

func echoer(chin, chout chan interface{}) {
	for m := range chin {
		chout <- m
	}
}

func BenchmarkDirectSend(b *testing.B) {
	chout := make(chan interface{})
	chin := make(chan interface{})
	defer close(chin)

	go echoer(chin, chout)

	for i := 0; i < b.N; i++ {
		chin <- nil
		<-chout
	}
}

func BenchmarkSend(b *testing.B) {
	chout := make(chan interface{})

	o := NewObserver(10)
	defer o.Close()
	o.Register(chout)

	for i := 0; i < b.N; i++ {
		o.Send(nil)
		<-chout
	}
}

func BenchmarkParallelSend(b *testing.B) {
	chout := make(chan interface{})

	o := NewObserver(10)
	defer o.Close()
	o.Register(chout)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			o.Send(nil)
			<-chout
		}
	})
}
