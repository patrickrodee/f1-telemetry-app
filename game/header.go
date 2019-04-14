package game

const headerSize = 21

type header struct {
	Format    uint16
	Version   uint8
	ID        uint8
	UUID      uint64
	Timestamp float32
	Frame     uint32
	Index     uint8
}
