package game

const (
	nameByteSize              = 48
	participantByteSize       = 1082
	participantByteActualSize = participantByteSize - headerSize
)

var teamNames = map[uint8]string{
	0:  "Mercedes",
	1:  "Ferrari",
	2:  "Red Bull",
	3:  "Williams",
	4:  "Force India",
	5:  "Renault",
	6:  "Toro Rosso",
	7:  "Haas",
	8:  "McLaren",
	9:  "Sauber",
	10: "McLaren 1988",
	11: "McLaren 1991",
	12: "Williams 1992",
	13: "Ferrari 1995",
	14: "Williams 1996",
	15: "McLaren 1998",
	16: "Ferrari 2002",
	17: "Ferrari 2004",
	18: "Renault 2006",
	19: "Ferrari 2007",
	20: "McLaren 2008",
	21: "Red Bull 2010",
	22: "Ferrari 1976",
	34: "McLaren 1976",
	35: "Lotus 1972",
	36: "Ferrari 1979",
	37: "McLaren 1982",
	38: "Williams 2003",
	39: "Brawn 2009",
	40: "Lotus 1978",
}

// Participant provides details about a single driver
type Participant struct {
	AIControlled uint8
	DriverID     uint8
	TeamID       uint8
	RaceNumber   uint8
	Nationality  uint8
	Name         [48]byte
}

func newParticipant(b []byte, next int) (Participant, int) {
	var p Participant
	p.AIControlled, next = bsuint8(b, next)
	p.DriverID, next = bsuint8(b, next)
	p.TeamID, next = bsuint8(b, next)
	p.RaceNumber, next = bsuint8(b, next)
	p.Nationality, next = bsuint8(b, next)
	end := next + nameByteSize
	copy(p.Name[:], b[next:end])
	return p, end
}

// IsHuman provides the humanity of the driver
func (p Participant) IsHuman() bool {
	return p.AIControlled == 0
}

// TeamName provides the human-readable name of the driver's team
func (p Participant) TeamName() string {
	return teamNames[p.TeamID]
}

// DriverName provides the human-readable name of the driver
func (p Participant) DriverName() string {
	return string(p.Name[:])
}

// ParticipantData provides the data for all drivers
type ParticipantData struct {
	NumCars      uint8
	Participants [20]Participant
}

func newParticipantData(b []byte, next int) ParticipantData {
	var p ParticipantData
	p.NumCars, next = bsuint8(b, next)
	for i := range p.Participants {
		p.Participants[i], next = newParticipant(b, next)
	}
	return p
}

// ActiveParticipants provides only the subset of active drivers
func (p ParticipantData) ActiveParticipants() []Participant {
	return p.Participants[:p.NumCars]
}
