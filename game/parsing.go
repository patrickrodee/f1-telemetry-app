package game

import (
	"encoding/binary"
	"math"
)

const (
	b8 = 1 << iota
	b16
	b32
	b64
)

func bsuint8(b []byte, start int) (uint8, int) {
	end := start + b8
	return uint8(b[start]), end
}

func bsuint16(b []byte, start int) (uint16, int) {
	end := start + b16
	return binary.LittleEndian.Uint16(b[start:end]), end
}

func bsuint32(b []byte, start int) (uint32, int) {
	end := start + b32
	return binary.LittleEndian.Uint32(b[start:end]), end
}

func bsuint64(b []byte, start int) (uint64, int) {
	end := start + b64
	return binary.LittleEndian.Uint64(b[start:end]), end
}

func bsfloat32(b []byte, start int) (float32, int) {
	v, end := bsuint32(b, start)
	return math.Float32frombits(v), end
}

func bsfloat64(b []byte, start int) (float64, int) {
	v, end := bsuint64(b, start)
	return math.Float64frombits(v), end
}

func bsint8(b []byte, start int) (int8, int) {
	v, end := bsuint8(b, start)
	return int8(v), end
}

func bsint16(b []byte, start int) (int16, int) {
	v, end := bsuint16(b, start)
	return int16(v), end
}

func bsbytes(b []byte, start, size int) ([]byte, int) {
	end := start + size
	return b[start:end], end
}
