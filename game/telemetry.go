package game

const (
	telemetryByteSize       = 1085
	telemetryByteActualSize = telemetryByteSize - headerSize
)

// WheelTemperature provides the temperature data for all wheels of a single car
type WheelTemperature struct {
	RearLeft   uint16
	RearRight  uint16
	FrontLeft  uint16
	FrontRight uint16
}

func newWheelTemperature(b []byte, next int) (WheelTemperature, int) {
	var w WheelTemperature
	w.RearLeft, next = bsuint16(b, next)
	w.RearRight, next = bsuint16(b, next)
	w.FrontLeft, next = bsuint16(b, next)
	w.FrontRight, next = bsuint16(b, next)
	return w, next
}

// Telemetry provides the data for a single car
type Telemetry struct {
	Speed                   uint16
	Throttle                uint8
	Steer                   int8
	Brake                   uint8
	Clutch                  uint8
	Gear                    int8
	EngineRPM               uint16
	DRS                     uint8
	RevLightsPercentage     uint8
	BrakesTemperature       WheelTemperature
	TyresSurfaceTemperature WheelTemperature
	TyresInnerTemperature   WheelTemperature
	EngineTemperature       uint16
	TyresPressure           WheelData
}

func newTelemetry(b []byte, next int) (Telemetry, int) {
	var t Telemetry
	t.Speed, next = bsuint16(b, next)
	t.Throttle, next = bsuint8(b, next)
	t.Steer, next = bsint8(b, next)
	t.Brake, next = bsuint8(b, next)
	t.Clutch, next = bsuint8(b, next)
	t.Gear, next = bsint8(b, next)
	t.EngineRPM, next = bsuint16(b, next)
	t.DRS, next = bsuint8(b, next)
	t.RevLightsPercentage, next = bsuint8(b, next)
	t.BrakesTemperature, next = newWheelTemperature(b, next)
	t.TyresSurfaceTemperature, next = newWheelTemperature(b, next)
	t.TyresInnerTemperature, next = newWheelTemperature(b, next)
	t.EngineTemperature, next = bsuint16(b, next)
	t.TyresPressure, next = newWheelData(b, next)
	return t, next
}

// TelemetryData provides telemetry for all cars plus the human driver's controller usage
type TelemetryData struct {
	Cars         [20]Telemetry
	ButtonStatus uint32
}

func newTelemetryData(b []byte, next int) TelemetryData {
	var t TelemetryData
	for i := range t.Cars {
		t.Cars[i], next = newTelemetry(b, next)
	}
	t.ButtonStatus, next = bsuint32(b, next)
	return t
}
