package game

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"log"
	"testing"
)

func motionWorker(s *Store) {
	ch := make(chan interface{})
	s.Motion.Register(ch)
	go func() {
		defer s.Motion.Deregister(ch)
		for v := range ch {
			if _, ok := v.(MotionData); !ok {
				panic(fmt.Errorf("motion worker couldn't coerce interface to motiondata"))
			}
		}
	}()
}

func lapWorker(s *Store) {
	ch := make(chan interface{})
	s.Lap.Register(ch)
	go func() {
		defer s.Lap.Deregister(ch)
		for v := range ch {
			if _, ok := v.(LapData); !ok {
				panic(fmt.Errorf("lap worker couldn't coerce interface to lapdata"))
			}
		}
	}()
}

func testMotionData(t *testing.T, m MotionData) []byte {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, m); err != nil {
		t.Error(err)
	}
	return buf.Bytes()
}

func TestStorePutMotion(t *testing.T) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	size := 10
	s := NewStore(size)
	ch1 := make(chan interface{})
	ch2 := make(chan interface{})
	s.Motion.Register(ch1)
	s.Motion.Register(ch2)
	defer s.Motion.Deregister(ch1)
	defer s.Motion.Deregister(ch2)
	token := make([]byte, 1341)
	for i := 0; i < size; i++ {
		if _, err := rand.Read(token); err != nil {
			t.Error(err)
		}
		token[3] = 0
		if err := s.Put(token); err != nil {
			t.Error(err)
		}
		var v1 interface{}
		var v2 interface{}
	Loop:
		for {
			select {
			case v1 = <-ch1:
			case v2 = <-ch2:
			default:
				if v1 != nil && v2 != nil {
					break Loop
				}
			}
		}
		if _, ok := v1.(MotionData); !ok {
			t.Errorf("motion worker couldn't coerce interface to motiondata")
		}
		if _, ok := v2.(MotionData); !ok {
			t.Errorf("motion worker couldn't coerce interface to motiondata")
		}
	}
}
