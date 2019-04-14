package game

import (
	"bytes"
	"encoding/binary"
	"net"
)

// Store contains the parsed information from a UDP stream
type Store struct {
	motion  *MotionData
	session *SessionData
	lap     *LapData
}

// NewStore creates and returns a new store
func NewStore() *Store {
	return &Store{}
}

// ReadMotion returns the store's motion information
func (s Store) ReadMotion() *MotionData {
	return s.motion
}

// WriteMotion writes new motion data to the store
func (s *Store) WriteMotion(b *bytes.Buffer) error {
	d := new(MotionData)
	if err := binary.Read(b, binary.LittleEndian, d); err != nil {
		return err
	}
	s.motion = d
	return nil
}

// WriteSession writes new session data to the store
func (s *Store) WriteSession(b *bytes.Buffer) error {
	d := new(SessionData)
	if err := binary.Read(b, binary.LittleEndian, d); err != nil {
		return err
	}
	s.session = d
	return nil
}

// WriteLap writes new lap data to the store
func (s *Store) WriteLap(b *bytes.Buffer) error {
	d := new(LapData)
	if err := binary.Read(b, binary.LittleEndian, d); err != nil {
		return err
	}
	s.lap = d
	return nil
}

// ReadLap returns the store's lap information
func (s Store) ReadLap() *LapData {
	return s.lap
}

// Start starts up a new game store
func (s Store) Start(addr *net.UDPAddr) error {
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	b := make([]byte, 1341)
	for {
		if _, err := conn.Read(b); err != nil {
			return err
		}

		buf := bytes.NewBuffer(b)
		hb := bytes.NewBuffer(buf.Next(headerSize))
		h := new(header)
		if err := binary.Read(hb, binary.LittleEndian, h); err != nil {
			return err
		}

		if h.ID == 0 {
			if err := s.WriteMotion(buf); err != nil {
				return err
			}
		}

		if h.ID == 1 {
			if err := s.WriteSession(buf); err != nil {
				return err
			}
		}

		if h.ID == 2 {
			if err := s.WriteLap(buf); err != nil {
				return err
			}
		}
	}
}
