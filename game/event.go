package game

const (
	eventByteSize       = 25
	eventByteActualSize = eventByteSize - headerSize
	start               = "SSTA"
	end                 = "SEND"
)

// EventData provides the events of the race
type EventData struct {
	Code [eventByteActualSize]byte
}

func newEventData(b []byte, next int) EventData {
	var e EventData
	end := next + eventByteActualSize
	copy(e.Code[:], b[next:end])
	return e
}

// Start returns if the event is a start event
func (e EventData) Start() bool {
	return string(e.Code[:]) == start
}

// End returns if the event is a end event
func (e EventData) End() bool {
	return string(e.Code[:]) == end
}
