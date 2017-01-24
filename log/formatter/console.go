package formatter

import (
	"github.com/Sirupsen/logrus"
)

// ConsoleFormatter is a formatter for the console.
type ConsoleFormatter struct {
	logrus.TextFormatter
}
