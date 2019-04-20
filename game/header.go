package game

const headerSize = 21

// Header provides detail about the incoming data
type Header struct {
	Format    uint16
	Version   uint8
	ID        uint8
	UUID      uint64
	Timestamp float32
	Frame     uint32
	Index     uint8
}

func newHeader(b []byte, next int) (Header, int) {
	var h Header
	h.Format, next = bsuint16(b, next)
	h.Version, next = bsuint8(b, next)
	h.ID, next = bsuint8(b, next)
	h.UUID, next = bsuint64(b, next)
	h.Timestamp, next = bsfloat32(b, next)
	h.Frame, next = bsuint32(b, next)
	h.Index, next = bsuint8(b, next)
	return h, next
}
