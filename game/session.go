package game

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
