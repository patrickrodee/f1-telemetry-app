package game

const (
	sessionByteSize       = 147
	sessionByteActualSize = sessionByteSize - headerSize
)

var (
	zoneFlags = map[int8]string{
		-1: "Unknown",
		0:  "None",
		1:  "Green",
		2:  "Blue",
		3:  "Yellow",
		4:  "Red",
	}

	weather = map[uint8]string{
		0: "Clear",
		1: "Light Clouds",
		2: "Overcast",
		3: "Light Rain",
		4: "Heavy Rain",
		5: "Storm",
	}

	sessionType = map[uint8]string{
		0:  "Unknown",
		1:  "FP1",
		2:  "FP2",
		3:  "FP3",
		4:  "Short FP",
		5:  "Q1",
		6:  "Q2",
		7:  "Q3",
		8:  "Short Q",
		9:  "OSQ",
		10: "R",
		11: "R2",
		12: "Time Trial",
	}

	tracks = map[int8]string{
		-1: "Unknown",
		0:  "Melbourne",
		1:  "Paul Ricard",
		2:  "Shanghai",
		3:  "Sakhir",
		4:  "Catalunya",
		5:  "Monaco",
		6:  "Montreal",
		7:  "Silverstone",
		8:  "Hockenheim",
		9:  "Hungaroring",
		10: "Spa",
		11: "Monza",
		12: "Singapore",
		13: "Suzuka",
		14: "Abu Dhabi",
		15: "Texas",
		16: "Brazil",
		17: "Austria",
		18: "Sochi",
		19: "Mexico",
		20: "Baku",
		21: "Sakhir (Short)",
		22: "Silverstone (Short)",
		23: "Texas (Short)",
		24: "Suzuka (Short)",
	}

	eras = map[uint8]string{
		0: "Modern",
		1: "Classic",
	}

	safetyCar = map[uint8]string{
		0: "No Safety Car",
		1: "Full Safety Car",
		2: "Virtual Safety Car",
	}
)

// MarshalZone provides zone data
type MarshalZone struct {
	ZoneStart float32
	ZoneFlag  int8
}

func newMarshalZone(b []byte, next int) (MarshalZone, int) {
	var m MarshalZone
	m.ZoneStart, next = bsfloat32(b, next)
	m.ZoneFlag, next = bsint8(b, next)
	return m, next
}

// CurrentFlag returns the current flag in human-readable form
func (m MarshalZone) CurrentFlag() string {
	return zoneFlags[m.ZoneFlag]
}

// Percentage returns the percentage in human-readable form
func (m MarshalZone) Percentage(end float32) float32 {
	return end - m.ZoneStart
}

// SessionData provides all the data about the current racing session
type SessionData struct {
	Weather             uint8
	TrackTemperature    int8
	AirTemperature      int8
	TotalLaps           uint8
	TrackLength         uint16
	SessionType         uint8
	TrackID             int8
	Era                 uint8
	SessionTimeLeft     uint16
	SessionDuration     uint16
	PitSpeedLimit       uint8
	GamePaused          uint8
	IsSpectating        uint8
	SpectatorCarIndex   uint8
	SliProNativeSupport uint8
	NumMarshalZones     uint8
	MarshalZones        [21]MarshalZone
	SafetyCarStatus     uint8
	NetworkGame         uint8
}

func newSessionData(b []byte, next int) SessionData {
	var s SessionData
	s.Weather, next = bsuint8(b, next)
	s.TrackTemperature, next = bsint8(b, next)
	s.AirTemperature, next = bsint8(b, next)
	s.TotalLaps, next = bsuint8(b, next)
	s.TrackLength, next = bsuint16(b, next)
	s.SessionType, next = bsuint8(b, next)
	s.TrackID, next = bsint8(b, next)
	s.Era, next = bsuint8(b, next)
	s.SessionTimeLeft, next = bsuint16(b, next)
	s.SessionDuration, next = bsuint16(b, next)
	s.PitSpeedLimit, next = bsuint8(b, next)
	s.GamePaused, next = bsuint8(b, next)
	s.IsSpectating, next = bsuint8(b, next)
	s.SpectatorCarIndex, next = bsuint8(b, next)
	s.SliProNativeSupport, next = bsuint8(b, next)
	s.NumMarshalZones, next = bsuint8(b, next)
	for i := range s.MarshalZones {
		s.MarshalZones[i], next = newMarshalZone(b, next)
	}
	s.SafetyCarStatus, next = bsuint8(b, next)
	s.NetworkGame, next = bsuint8(b, next)
	return s
}

// CurrentWeather provides the weather in human-readable form
func (s SessionData) CurrentWeather() string {
	return weather[s.Weather]
}

// CurrentSessionType provides the session type in human-readable form
func (s SessionData) CurrentSessionType() string {
	return sessionType[s.SessionType]
}

// CurrentTrack provides the name of the current track
func (s SessionData) CurrentTrack() string {
	return tracks[s.TrackID]
}

// CurrentEra provides the era of the current session
func (s SessionData) CurrentEra() string {
	return eras[s.Era]
}

// ZoneFlags provides the flags of the current session
func (s SessionData) ZoneFlags() []string {
	flags := make([]string, s.NumMarshalZones)
	for i := range flags {
		flags[i] = s.MarshalZones[i].CurrentFlag()
	}
	return flags
}

// ZonePercentages provides the percentages of the current session
func (s SessionData) ZonePercentages() []float32 {
	zones := make([]float32, s.NumMarshalZones)
	for i := range zones {
		zone := s.MarshalZones[i]
		if i == cap(zones)-1 {
			zones[i] = zone.Percentage(1)
			continue
		}
		zones[i] = zone.Percentage(s.MarshalZones[i+1].ZoneStart)
	}
	return zones
}

// SafetyCar provides the human-readable status of the safety car
func (s SessionData) SafetyCar() string {
	return safetyCar[s.SafetyCarStatus]
}
