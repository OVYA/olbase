package log

import (
	"bytes"
	"testing"

	"github.com/Sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func ExampleNewTextLogger() {
	log := NewTextLogger()

	log.WithFields(logrus.Fields{
		"bool_field": true,
	}).Info("I'll be logged with a bool_field field")
}

func TestConsoleLogger(t *testing.T) {
	var buffer bytes.Buffer

	log := NewTextLogger()
	log.Out = &buffer

	log.WithFields(logrus.Fields{
		"bool_field": true,
	}).Info("I'll be logged with a bool_field field")

	str := string(buffer.Bytes())

	assert.Contains(t, str, "I'll be logged with a bool_field field")
	assert.Contains(t, str, "bool_field=true")
	assert.Contains(t, str, "time=")
}
