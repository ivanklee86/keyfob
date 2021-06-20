package keyfob

import (
	"bytes"
	"testing"
)

func TestKeyfobHappyPath(t *testing.T) {
	b := bytes.NewBufferString("")

	kf := New()
	kf.Out = b
	kf.Err = b
}
