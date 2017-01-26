package log

import (
	"bytes"
	"testing"

	"github.com/Sirupsen/logrus"
)

func ExampleNewConsoleLogger() {
	var buffer bytes.Buffer

	log := NewConsoleLogger()
	log.Out = &buffer

	log.WithFields(logrus.Fields{
		"bool_field": true,
	}).Info("I'll be logged with a bool_field field")
}

func TestConsoleLogger(t *testing.T) {
	ExampleNewConsoleLogger()
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
