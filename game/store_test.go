package game

import (
	"crypto/rand"
	"testing"
)

func BenchmarkPutAndReceive(b *testing.B) {
	chout := make(chan interface{})

	s := NewStore(100)
	defer s.Motion.Close()
	s.Motion.Register(chout)
	bytes := make([]byte, 1341)
	if _, err := rand.Read(bytes); err != nil {
		b.Error(err)
	}
	bytes[3] = 0
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			s.Put(bytes)
			<-chout
		}
	})
}
