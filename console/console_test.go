package console

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAskYN(t *testing.T) {
	assert := assert.New(t)

	w := new(bytes.Buffer)
	output := bufio.NewWriterSize(w, 1024)

	tests := map[string]func(bool, ...interface{}) bool{
		"y\n":    assert.True,
		"Y\n":    assert.True,
		"q\ny\n": assert.True,
		"q\nY\n": assert.True,
		"n\n":    assert.False,
		"N\n":    assert.False,
		"q\nn\n": assert.False,
		"q\nN\n": assert.False,
	}

	for key, fnc := range tests {
		input := bytes.NewBufferString(key)
		fnc(askYN("Sure ?", output, input))
	}
}
