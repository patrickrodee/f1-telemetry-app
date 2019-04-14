package game

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

// WheelData provides information for all wheels of a single car
type WheelData struct {
	RearLeft   float32
	RearRight  float32
	FrontLeft  float32
	RrontRight float32
}

// MotionData provides motion data for all cars plus additional data for the driver's car
type MotionData struct {
	CarMotion              [20]CarMotion
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
