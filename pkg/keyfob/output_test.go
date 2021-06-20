package keyfob

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/acarl005/stripansi"
	"github.com/stretchr/testify/assert"
)

func TestOutputs(t *testing.T) {
	b := bytes.NewBufferString("")

	kf := New()
	kf.Out = b
	kf.Err = b

	testPhrase := "I'm a little hamster."

	t.Run("outputs string", func(t *testing.T) {
		kf.Output(testPhrase)

		out, err := ioutil.ReadAll(b)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, testPhrase+"\n", string(out))
	})

	t.Run("outputs header", func(t *testing.T) {
		kf.OutputHeading(testPhrase)

		out, err := ioutil.ReadAll(b)
		if err != nil {
			t.Fatal(err)
		}

		assert.Contains(t, stripansi.Strip(string(out)), testPhrase)
	})
}
