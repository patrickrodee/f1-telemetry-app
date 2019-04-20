package game

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"io"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewHeader(t *testing.T) {
	testdata := make([]byte, headerSize)
	rand.Read(testdata)
	got, _ := newHeader(testdata, 0)
	var want Header
	buf := bytes.NewBuffer(testdata)
	err := binary.Read(buf, binary.LittleEndian, &want)
	if err != nil && err != io.EOF {
		t.Error(err)
	}
	if diff := cmp.Diff(want, got, withEqualFloat32NaNs); diff != "" {
		t.Errorf("newHeader() mismatch (-want +got):\n%s", diff)
	}
}
