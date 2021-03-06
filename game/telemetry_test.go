package game

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"io"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewTelemetryData(t *testing.T) {
	testdata := make([]byte, telemetryByteActualSize)
	rand.Read(testdata)
	got := newTelemetryData(testdata, 0)
	var want TelemetryData
	buf := bytes.NewBuffer(testdata)
	err := binary.Read(buf, binary.LittleEndian, &want)
	if err != nil && err != io.EOF {
		t.Error(err)
	}
	if diff := cmp.Diff(want, got, withEqualFloat32NaNs); diff != "" {
		t.Errorf("newTelemetryData() mismatch (-want +got):\n%s", diff)
	}
}

func BenchmarkNewTelemetryData(b *testing.B) {
	testdata := make([]byte, telemetryByteActualSize)
	rand.Read(testdata)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		newTelemetryData(testdata, 0)
	}
}

func BenchmarkNewTelemetryDataWithReflection(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}
	testdata := make([]byte, telemetryByteActualSize)
	rand.Read(testdata)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		d := new(TelemetryData)
		buf := bytes.NewBuffer(testdata)
		err := binary.Read(buf, binary.LittleEndian, d)
		if err != nil && err != io.EOF {
			b.Error(err)
		}
	}
}
