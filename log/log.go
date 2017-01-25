package log

import (
	"ovya/olbase/log/formatter"

	"github.com/Sirupsen/logrus"
)

// Logger is the Ovya logger, wrapping logrus.Logger
// to ovoid app logrus hard dependency
type Logger struct {
	*logrus.Logger
}

// NewConsoleLogger is a constructor for a Logrus text based formatter
func NewConsoleLogger() (log *Logger) {
	log = &Logger{}
	log.Logger = logrus.New()

	log.Formatter = &formatter.ConsoleFormatter{
		logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
		},
	}

	return
}
