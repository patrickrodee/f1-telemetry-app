package game

const (
	setupByteSize       = 841
	setupByteActualSize = setupByteSize - headerSize
)

// Setup provides the setup data of a single car
type Setup struct {
	FrontWing             uint8
	RearWing              uint8
	OnThrottle            uint8
	OffThrottle           uint8
	FrontCamber           float32
	RearCamber            float32
	FrontToe              float32
	RearToe               float32
	FrontSuspension       uint8
	RearSuspension        uint8
	FrontAntiRollBar      uint8
	RearAntiRollBar       uint8
	FrontSuspensionHeight uint8
	RearSuspensionHeight  uint8
	BrakePressure         uint8
	BrakeBias             uint8
	FrontTyrePressure     float32
	RearTyrePressure      float32
	Ballast               uint8
	FuelLoad              float32
}

func newSetup(b []byte, next int) (Setup, int) {
	var s Setup
	s.FrontWing, next = bsuint8(b, next)
	s.RearWing, next = bsuint8(b, next)
	s.OnThrottle, next = bsuint8(b, next)
	s.OffThrottle, next = bsuint8(b, next)
	s.FrontCamber, next = bsfloat32(b, next)
	s.RearCamber, next = bsfloat32(b, next)
	s.FrontToe, next = bsfloat32(b, next)
	s.RearToe, next = bsfloat32(b, next)
	s.FrontSuspension, next = bsuint8(b, next)
	s.RearSuspension, next = bsuint8(b, next)
	s.FrontAntiRollBar, next = bsuint8(b, next)
	s.RearAntiRollBar, next = bsuint8(b, next)
	s.FrontSuspensionHeight, next = bsuint8(b, next)
	s.RearSuspensionHeight, next = bsuint8(b, next)
	s.BrakePressure, next = bsuint8(b, next)
	s.BrakeBias, next = bsuint8(b, next)
	s.FrontTyrePressure, next = bsfloat32(b, next)
	s.RearTyrePressure, next = bsfloat32(b, next)
	s.Ballast, next = bsuint8(b, next)
	s.FuelLoad, next = bsfloat32(b, next)
	return s, next
}

// SetupData provides the setup data of all cars
type SetupData struct {
	Setups [20]Setup
}

func newSetupData(b []byte, next int) SetupData {
	var d SetupData
	for i := range d.Setups {
		d.Setups[i], next = newSetup(b, next)
	}
	return d
}
