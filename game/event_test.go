package game

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"io"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewEventData(t *testing.T) {
	testdata := make([]byte, eventByteActualSize)
	rand.Read(testdata)
	got := newEventData(testdata, 0)
	var want EventData
	buf := bytes.NewBuffer(testdata)
	err := binary.Read(buf, binary.LittleEndian, &want)
	if err != nil && err != io.EOF {
		t.Error(err)
	}
	if diff := cmp.Diff(want, got, withEqualFloat32NaNs); diff != "" {
		t.Errorf("newEventData() mismatch (-want +got):\n%s", diff)
	}
}

func BenchmarkNewEventData(b *testing.B) {
	testdata := make([]byte, eventByteActualSize)
	rand.Read(testdata)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		newEventData(testdata, 0)
	}
}

func BenchmarkNewEventDataWithReflection(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}
	testdata := make([]byte, eventByteActualSize)
	rand.Read(testdata)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		d := new(EventData)
		buf := bytes.NewBuffer(testdata)
		err := binary.Read(buf, binary.LittleEndian, d)
		if err != nil && err != io.EOF {
			b.Error(err)
		}
	}
}
