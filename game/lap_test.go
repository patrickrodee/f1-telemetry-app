package game

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"io"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewLap(t *testing.T) {
	testdata := make([]byte, lapByteActualSize)
	rand.Read(testdata)
	got, _ := newLapData(testdata, 0)
	var want LapData
	buf := bytes.NewBuffer(testdata)
	err := binary.Read(buf, binary.LittleEndian, &want)
	if err != nil && err != io.EOF {
		t.Error(err)
	}
	if diff := cmp.Diff(want, got, withEqualFloat32NaNs); diff != "" {
		t.Errorf("newLapData() mismatch (-want +got):\n%s", diff)
	}
}

func BenchmarkNewLapData(b *testing.B) {
	testdata := make([]byte, lapByteActualSize)
	rand.Read(testdata)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		newLapData(testdata, 0)
	}
}

func BenchmarkNewLapDataWithReflection(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}
	testdata := make([]byte, lapByteActualSize)
	rand.Read(testdata)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		d := new(LapData)
		buf := bytes.NewBuffer(testdata)
		err := binary.Read(buf, binary.LittleEndian, d)
		if err != nil && err != io.EOF {
			b.Error(err)
		}
	}
}
