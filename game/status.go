package game

const (
	statusByteSize       = 1061
	statusByteActualSize = statusByteSize - headerSize
)

// TyreWear provides the wear data for all wheels of a car
type TyreWear struct {
	RearLeft   uint8
	RearRight  uint8
	FrontLeft  uint8
	FrontRight uint8
}

func newTyreWear(b []byte, n int) (TyreWear, int) {
	var t TyreWear
	t.RearLeft, n = bsuint8(b, n)
	t.RearRight, n = bsuint8(b, n)
	t.FrontLeft, n = bsuint8(b, n)
	t.FrontRight, n = bsuint8(b, n)
	return t, n
}

// CarStatus provides that status data for a car
type CarStatus struct {
	TractionControl         uint8
	AntiLockBrakes          uint8
	FuelMix                 uint8
	FrontBrakeBias          uint8
	PitLimiterStatus        uint8
	FuelInTank              float32
	FuelCapacity            float32
	MaxRPM                  uint16
	IdleRPM                 uint16
	MaxGears                uint8
	DRSAllowed              uint8
	TyresWear               TyreWear
	TyreCompound            uint8
	TyresDamage             TyreWear
	FrontLeftWingDamage     uint8
	FrontRightWingDamage    uint8
	RearWingDamage          uint8
	EngineDamage            uint8
	GearBoxDamage           uint8
	ExhaustDamage           uint8
	VehicleFIAFlags         int8
	ERSStoreEnergy          float32
	ERSDeployMode           uint8
	ERSHarvestedThisLapMGUK float32
	ERSHarvestedThisLapMGUH float32
	ERSDeployedThisLap      float32
}

func newCarStatus(b []byte, n int) (CarStatus, int) {
	var c CarStatus
	c.TractionControl, n = bsuint8(b, n)
	c.AntiLockBrakes, n = bsuint8(b, n)
	c.FuelMix, n = bsuint8(b, n)
	c.FrontBrakeBias, n = bsuint8(b, n)
	c.PitLimiterStatus, n = bsuint8(b, n)
	c.FuelInTank, n = bsfloat32(b, n)
	c.FuelCapacity, n = bsfloat32(b, n)
	c.MaxRPM, n = bsuint16(b, n)
	c.IdleRPM, n = bsuint16(b, n)
	c.MaxGears, n = bsuint8(b, n)
	c.DRSAllowed, n = bsuint8(b, n)
	c.TyresWear, n = newTyreWear(b, n)
	c.TyreCompound, n = bsuint8(b, n)
	c.TyresDamage, n = newTyreWear(b, n)
	c.FrontLeftWingDamage, n = bsuint8(b, n)
	c.FrontRightWingDamage, n = bsuint8(b, n)
	c.RearWingDamage, n = bsuint8(b, n)
	c.EngineDamage, n = bsuint8(b, n)
	c.GearBoxDamage, n = bsuint8(b, n)
	c.ExhaustDamage, n = bsuint8(b, n)
	c.VehicleFIAFlags, n = bsint8(b, n)
	c.ERSStoreEnergy, n = bsfloat32(b, n)
	c.ERSDeployMode, n = bsuint8(b, n)
	c.ERSHarvestedThisLapMGUK, n = bsfloat32(b, n)
	c.ERSHarvestedThisLapMGUH, n = bsfloat32(b, n)
	c.ERSDeployedThisLap, n = bsfloat32(b, n)
	return c, n
}

// StatusData provides the status of all cars
type StatusData struct {
	Cars [20]CarStatus
}

func newStatusData(b []byte, n int) StatusData {
	var s StatusData
	for i := range s.Cars {
		s.Cars[i], n = newCarStatus(b, n)
	}
	return s
}
