package game

// Car provides the setup data of a single car
type Car struct {
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

// CarData provides the setup data of all cars
type CarData struct {
	CarData [20]Car
}
