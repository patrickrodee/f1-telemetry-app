package game

const (
	motionByteSize       = 1341
	motionByteActualSize = motionByteSize - headerSize
	wheelDataSize        = 16
	carMotionCount       = 20
	wheelDataCount       = 5
)

// CarMotion provides the motion data for a single car
type CarMotion struct {
	WorldPositionX     float32
	WorldPositionY     float32
	WorldPositionZ     float32
	WorldVelocityX     float32
	WorldVelocityY     float32
	WorldVelocityZ     float32
	WorldForwardDirX   int16
	WorldForwardDirY   int16
	WorldForwardDirZ   int16
	WorldRightDirX     int16
	WorldRightDirY     int16
	WorldRightDirZ     int16
	GForceLateral      float32
	GForceLongitudinal float32
	GForceVertical     float32
	Yaw                float32
	Pitch              float32
	Roll               float32
}

func newCarMotion(b []byte, next int) (CarMotion, int) {
	var m CarMotion
	m.WorldPositionX, next = bsfloat32(b, next)
	m.WorldPositionY, next = bsfloat32(b, next)
	m.WorldPositionZ, next = bsfloat32(b, next)
	m.WorldVelocityX, next = bsfloat32(b, next)
	m.WorldVelocityY, next = bsfloat32(b, next)
	m.WorldVelocityZ, next = bsfloat32(b, next)
	m.WorldForwardDirX, next = bsint16(b, next)
	m.WorldForwardDirY, next = bsint16(b, next)
	m.WorldForwardDirZ, next = bsint16(b, next)
	m.WorldRightDirX, next = bsint16(b, next)
	m.WorldRightDirY, next = bsint16(b, next)
	m.WorldRightDirZ, next = bsint16(b, next)
	m.GForceLateral, next = bsfloat32(b, next)
	m.GForceLongitudinal, next = bsfloat32(b, next)
	m.GForceVertical, next = bsfloat32(b, next)
	m.Yaw, next = bsfloat32(b, next)
	m.Pitch, next = bsfloat32(b, next)
	m.Roll, next = bsfloat32(b, next)
	return m, next
}

// WheelData provides information for all wheels of a single car
type WheelData struct {
	RearLeft   float32
	RearRight  float32
	FrontLeft  float32
	FrontRight float32
}

func newWheelData(b []byte, next int) (WheelData, int) {
	var w WheelData
	w.RearLeft, next = bsfloat32(b, next)
	w.RearRight, next = bsfloat32(b, next)
	w.FrontLeft, next = bsfloat32(b, next)
	w.FrontRight, next = bsfloat32(b, next)
	return w, next
}

// MotionData provides motion data for all cars plus additional data for the driver's car
type MotionData struct {
	CarMotion              [carMotionCount]CarMotion
	SuspensionVelocity     WheelData
	SuspensionPosition     WheelData
	SuspensionAcceleration WheelData
	WheelSpeed             WheelData
	WheelSlip              WheelData
	LocalVelocityX         float32
	LocalVelocityY         float32
	LocalVelocityZ         float32
	AngularVelocityX       float32
	AngularVelocityY       float32
	AngularVelocityZ       float32
	AngularAccelerationX   float32
	AngularAccelerationY   float32
	AngularAccelerationZ   float32
	FrontWheelsAngle       float32
}

func newMotionData(b []byte, next int) (MotionData, int) {
	var m MotionData
	for i := 0; i < carMotionCount; i++ {
		m.CarMotion[i], next = newCarMotion(b, next)
	}
	m.SuspensionVelocity, next = newWheelData(b, next)
	m.SuspensionPosition, next = newWheelData(b, next)
	m.SuspensionAcceleration, next = newWheelData(b, next)
	m.WheelSpeed, next = newWheelData(b, next)
	m.WheelSlip, next = newWheelData(b, next)
	m.LocalVelocityX, next = bsfloat32(b, next)
	m.LocalVelocityY, next = bsfloat32(b, next)
	m.LocalVelocityZ, next = bsfloat32(b, next)
	m.AngularVelocityX, next = bsfloat32(b, next)
	m.AngularVelocityY, next = bsfloat32(b, next)
	m.AngularVelocityZ, next = bsfloat32(b, next)
	m.AngularAccelerationX, next = bsfloat32(b, next)
	m.AngularAccelerationY, next = bsfloat32(b, next)
	m.AngularAccelerationZ, next = bsfloat32(b, next)
	m.FrontWheelsAngle, next = bsfloat32(b, next)
	return m, next
}
