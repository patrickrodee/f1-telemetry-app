package game

const (
	lapByteSize       = 841
	lapByteActualSize = lapByteSize - headerSize
)

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

func newLap(b []byte, next int) (Lap, int) {
	var l Lap
	l.LastLapTime, next = bsfloat32(b, next)
	l.CurrentLapTime, next = bsfloat32(b, next)
	l.BestLapTime, next = bsfloat32(b, next)
	l.Sector1Time, next = bsfloat32(b, next)
	l.Sector2Time, next = bsfloat32(b, next)
	l.LapDistance, next = bsfloat32(b, next)
	l.TotalDistance, next = bsfloat32(b, next)
	l.SafetyCarDelta, next = bsfloat32(b, next)
	l.CarPosition, next = bsuint8(b, next)
	l.CurrentLapNum, next = bsuint8(b, next)
	l.PitStatus, next = bsuint8(b, next)
	l.Sector, next = bsuint8(b, next)
	l.CurrentLapInvalid, next = bsuint8(b, next)
	l.Penalties, next = bsuint8(b, next)
	l.GridPosition, next = bsuint8(b, next)
	l.DriverStatus, next = bsuint8(b, next)
	l.ResultStatus, next = bsuint8(b, next)
	return l, next
}

// LapData provides the lap information of all drivers
type LapData struct {
	Laps [20]Lap
}

func newLapData(b []byte, next int) (LapData, int) {
	var l LapData
	for i := range l.Laps {
		l.Laps[i], next = newLap(b, next)
	}
	return l, next
}
