package game

// Lap provides the lap information of a driver
type Lap struct {
	LastLapTime       float32
	CurrentLapTime    float32
	BestLapTime       float32
	Sector1Time       float32
	Sector2Time       float32
	LapDistance       float32
	TotalDistance     float32
	SafetyCarDelta    float32
	CarPosition       uint8
	CurrentLapNum     uint8
	PitStatus         uint8
	Sector            uint8
	CurrentLapInvalid uint8
	Penalties         uint8
	GridPosition      uint8
	DriverStatus      uint8
	ResultStatus      uint8
}

// LapData provides the lap information of all drivers
type LapData struct {
	LapData [20]Lap
}
