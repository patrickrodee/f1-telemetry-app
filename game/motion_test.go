package game

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"io"
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var oneCarData = []byte{
	// -- WorldPosition --
	// float32: 8.25
	0, 0, 4, 65,
	// float32: 33
	0, 0, 4, 66,
	// float32: -0.1289139
	1, 2, 4, 190,
	// -- WorldVelocity --
	// float32: 8.25
	0, 0, 4, 65,
	// float32: 33
	0, 0, 4, 66,
	// float32: 2112.1252
	1, 2, 4, 69,
	// -- WorldForwardDir --
	// int16: 262
	6, 1,
	// int16: -9210
	6, 220,
	// int16: 476
	220, 1,
	// -- WorldRightDir --
	// int16: 262
	6, 1,
	// int16: -9210
	6, 220,
	// int16: 476
	220, 1,
	// -- GForce --
	// float32: -0.1289139
	1, 2, 4, 190,
	// float32: 2112.1252
	1, 2, 4, 69,
	// float32: 8.25
	0, 0, 4, 65,
	// -- Yaw, Pitch, and Roll --
	// float32: 0.008057118
	1, 2, 4, 60,
	// float32: 0.032228474
	1, 2, 4, 61,
	// float32: 0.5186768
	1, 200, 4, 63,
}

var oneWheelData = []byte{
	// -- RearLeft --
	// float32: 0.007873536
	1, 0, 1, 60,
	// -- RearRight --
	// float32: 0.011535883
	0, 1, 61, 60,
	// -- FrontLeft --
	// float32: 0.007873536
	1, 0, 1, 60,
	// -- FrontRight --
	// float32: 0.011535883
	0, 1, 61, 60,
}

// 47.309803
var f32 = []byte{61, 61, 61, 61}

var withEqualFloat32NaNs = cmp.Comparer(func(x, y float32) bool {
	return (math.IsNaN(float64(x)) && math.IsNaN(float64(y))) || x == y
})

func initData() []byte {
	var testData []byte
	for i := 0; i < carMotionCount; i++ {
		testData = append(testData, oneCarData...)
	}
	for i := 0; i < wheelDataCount; i++ {
		testData = append(testData, oneWheelData...)
	}
	for i := 0; i < 10; i++ {
		testData = append(testData, f32...)
	}
	return testData
}

func TestNewMotionData(t *testing.T) {
	testdata := make([]byte, motionByteActualSize)
	rand.Read(testdata)
	got := newMotionData(testdata, 0)
	var want MotionData
	buf := bytes.NewBuffer(testdata)
	err := binary.Read(buf, binary.LittleEndian, &want)
	if err != nil && err != io.EOF {
		t.Error(err)
	}
	if diff := cmp.Diff(want, got, withEqualFloat32NaNs); diff != "" {
		t.Errorf("newMotionData() mismatch (-want +got):\n%s", diff)
	}
}

func BenchmarkNewMotionData(b *testing.B) {
	testdata := make([]byte, motionByteActualSize)
	rand.Read(testdata)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		newMotionData(testdata, 0)
	}
}

func BenchmarkNewMotionDataWithReflection(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}
	testdata := make([]byte, motionByteActualSize)
	rand.Read(testdata)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		d := new(MotionData)
		buf := bytes.NewBuffer(testdata)
		err := binary.Read(buf, binary.LittleEndian, d)
		if err != nil && err != io.EOF {
			b.Error(err)
		}
	}
}
