package log

import (
	"ovya.fr/olbase.git/log/formatter"

	"github.com/Sirupsen/logrus"
)

// Logger is the Ovya logger.
// Actuelly this is a wrapper of github.com/Sirupsen/logrus.Logger
// See https://github.com/sirupsen/logrus for the further documentation
type Logger struct {
	logrus.Logger
}

// NewTextLogger is a constructor for a Logrus text based formatter
func NewTextLogger() (log *Logger) {
	log = &Logger{*logrus.New()}

	log.Formatter = &formatter.TextFormatter{
		TextFormatter: logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
		},
	}

	return
}
