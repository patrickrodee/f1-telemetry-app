package game

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"io"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewStatusData(t *testing.T) {
	testdata := make([]byte, statusByteActualSize)
	rand.Read(testdata)
	got := newStatusData(testdata, 0)
	var want StatusData
	buf := bytes.NewBuffer(testdata)
	err := binary.Read(buf, binary.LittleEndian, &want)
	if err != nil && err != io.EOF {
		t.Error(err)
	}
	if diff := cmp.Diff(want, got, withEqualFloat32NaNs); diff != "" {
		t.Errorf("newStatusData() mismatch (-want +got):\n%s", diff)
	}
}

func BenchmarkNewStatusData(b *testing.B) {
	testdata := make([]byte, statusByteActualSize)
	rand.Read(testdata)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		newStatusData(testdata, 0)
	}
}

func BenchmarkNewStatusDataWithReflection(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}
	testdata := make([]byte, statusByteActualSize)
	rand.Read(testdata)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		d := new(StatusData)
		buf := bytes.NewBuffer(testdata)
		err := binary.Read(buf, binary.LittleEndian, d)
		if err != nil && err != io.EOF {
			b.Error(err)
		}
	}
}
