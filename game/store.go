package game

import (
	"net"

	"github.com/patrickrodee/f1-telemetry-app/observer"
)

// Store contains the parsed information from a UDP stream
type Store struct {
	Motion      observer.Observer
	Session     observer.Observer
	Lap         observer.Observer
	Event       observer.Observer
	Participant observer.Observer
	Setup       observer.Observer
	Telemetry   observer.Observer
	Status      observer.Observer
}

// NewStore creates and returns a new store
func NewStore(buflen int) *Store {
	return &Store{
		Motion:      observer.NewObserver(buflen),
		Session:     observer.NewObserver(buflen),
		Lap:         observer.NewObserver(buflen),
		Event:       observer.NewObserver(buflen),
		Participant: observer.NewObserver(buflen),
		Setup:       observer.NewObserver(buflen),
		Telemetry:   observer.NewObserver(buflen),
		Status:      observer.NewObserver(buflen),
	}
}

// Put updates the store
func (s *Store) Put(b []byte) {
	h, next := newHeader(b, 0)
	switch h.ID {
	case 0:
		s.putMotion(b, next)
	case 1:
		s.putSession(b, next)
	case 2:
		s.putLap(b, next)
	case 3:
		s.putEvent(b, next)
	case 4:
		s.putParticipant(b, next)
	case 5:
		s.putSetup(b, next)
	case 6:
		s.putTelemetry(b, next)
	case 7:
		s.putStatus(b, next)
	}
}

func (s *Store) putMotion(b []byte, next int) {
	s.Motion.Send(newMotionData(b, next))
}

func (s *Store) putSession(b []byte, next int) {
	s.Session.Send(newSessionData(b, next))
}

func (s *Store) putLap(b []byte, next int) {
	s.Lap.Send(newLapData(b, next))
}

func (s *Store) putEvent(b []byte, next int) {
	s.Event.Send(newEventData(b, next))
}

func (s *Store) putParticipant(b []byte, next int) {
	s.Participant.Send(newParticipantData(b, next))
}

func (s *Store) putSetup(b []byte, next int) {
	s.Setup.Send(newSetupData(b, next))
}

func (s *Store) putTelemetry(b []byte, next int) {
	s.Telemetry.Send(newTelemetryData(b, next))
}

func (s *Store) putStatus(b []byte, next int) {
	s.Status.Send(newStatusData(b, next))
}

// Start starts up a new game store
func (s Store) Start(conn *net.UDPConn) error {
	defer conn.Close()
	b := make([]byte, 1341)
	for {
		if _, err := conn.Read(b); err != nil {
			return err
		}

		s.Put(b)
	}
}
