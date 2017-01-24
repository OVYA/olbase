package log

import (
	"ovya/olbase/log/formatter"

	"github.com/Sirupsen/logrus"
)

// NewConsoleLogger is a constructor for a Logrus text based formatter
func NewConsoleLogger() (log *logrus.Logger) {
	log = logrus.New()
	log.Formatter = &formatter.ConsoleFormatter{
		logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
		},
	}

	return
}
