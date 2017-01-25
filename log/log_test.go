package log

import (
	"bytes"
	"testing"
)

func TestConsoleLogger(t *testing.T) {
	var buffer bytes.Buffer

	log := NewConsoleLogger()
	log.Out = &buffer
	log.Info("I'll be logged with common and other field")
}
