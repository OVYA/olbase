package log

import (
	"ovya/lib/olbase/log/formatter"

	"github.com/Sirupsen/logrus"
)

// Logger is the Ovya logger.
// Actuelly this is a wrapper of github.com/Sirupsen/logrus.Logger
// See https://github.com/sirupsen/logrus for the further documentation
type Logger struct {
	logrus.Logger
}

// NewDefaultLogger
func NewDefaultLogger() *Logger {

	logger := &Logger{*logrus.New()}
	logger.Formatter = new(formatter.DefaultFormatter)

	return logger
}

// NewConsoleLogger is a constructor for a Logrus text based formatter
func NewConsoleLogger() (log *Logger) {
	log = &Logger{*logrus.New()}

	log.Formatter = &formatter.ConsoleFormatter{
		logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
		},
	}

	return
}
