package log

import (
	"bytes"
	"testing"

	"github.com/Sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func ExampleNewConsoleLogger() {
	log := NewConsoleLogger()

	log.WithFields(logrus.Fields{
		"bool_field": true,
	}).Info("I'll be logged with a bool_field field")
}

func TestConsoleLogger(t *testing.T) {
	var buffer bytes.Buffer

	log := NewConsoleLogger()
	log.Out = &buffer

	log.WithFields(logrus.Fields{
		"bool_field": true,
	}).Info("I'll be logged with a bool_field field")

	str := string(buffer.Bytes())

	assert.Contains(t, str, "I'll be logged with a bool_field field")
	assert.Contains(t, str, "bool_field=true")
	assert.Contains(t, str, "time=")
}

func TestNewDefaultLogger(t *testing.T) {

	logger := NewDefaultLogger()
	logger.Info("Log Info with default logger")
}

// func TestErrorLogger(t *testing.T) {
// 	// var buffer bytes.Buffer

// 	log := NewConsoleLogger()
// 	// log.Out = &buffer
// 	err := orror.New("plop")
// 	// fmt.Println(err.Error())
// 	log.Info(err.Error())
// 	log.Errorln(err)

// 	log.WithFields(logrus.Fields{
// 		"backtrace": err.GetStack(),
// 	}).Fatal(err.GetMessage())
// }
