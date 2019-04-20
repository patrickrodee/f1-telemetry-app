package game

import (
	"net"

	"github.com/patrickrodee/f1-telemetry-app/observer"
)

// Store contains the parsed information from a UDP stream
type Store struct {
	Motion  observer.Observer
	Session observer.Observer
	Lap     observer.Observer
}

// Put updates the store
func (s *Store) Put(b []byte) error {
	h, next := newHeader(b, 0)
	if h.ID == 0 {
		s.putMotion(b, next)
	}

	// if h.ID == 1 {
	// 	return s.putSession(buf)
	// }

	// if h.ID == 2 {
	// 	return s.putLap(buf)
	// }

	return nil
}

func (s *Store) putMotion(b []byte, next int) {
	m, _ := newMotionData(b, next)
	s.Motion.Send(m)
}

// func (s *Store) putSession(buf *bytes.Buffer) error {
// 	d := new(MotionData)
// 	err := binary.Read(buf, binary.LittleEndian, d)
// 	if err != nil && err != io.EOF {
// 		return err
// 	}
// 	s.Session.Send(d)
// 	return nil
// }

// func (s *Store) putLap(buf *bytes.Buffer) error {
// 	d := new(LapData)
// 	err := binary.Read(buf, binary.LittleEndian, d)
// 	if err != nil && err != io.EOF {
// 		return err
// 	}
// 	s.Lap.Send(d)
// 	return nil
// }

// NewStore creates and returns a new store
func NewStore(buflen int) *Store {
	return &Store{
		Motion:  observer.NewObserver(buflen),
		Session: observer.NewObserver(buflen),
		Lap:     observer.NewObserver(buflen),
	}
}

// Start starts up a new game store
func (s Store) Start(conn *net.UDPConn) error {
	defer conn.Close()
	b := make([]byte, 1341)
	for {
		if _, err := conn.Read(b); err != nil {
			return err
		}

		if err := s.Put(b); err != nil {
			return err
		}
	}
}
