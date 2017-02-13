package formatter

import (
	"fmt"
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/facebookgo/stack"
)

// TextFormatter is a formatter for the console.
type TextFormatter struct {
	logrus.TextFormatter
	// CallerLevel is the level used by runtime.Caller(level) to
	// capture the file and line number where the log is done.
	// Default value is 4 for the need of our library, so, can not be set to 0.
	CallerLevel int
}

// Format implements our custom Formatter that add the field `caller`.
func (formatter *TextFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	level := 4

	if formatter.CallerLevel != 0 {
		level = formatter.CallerLevel
	}

	frame := stack.Caller(level)
	entry.Data["caller"] = fmt.Sprintf("%s:%s (%s)", frame.File, strconv.Itoa(frame.Line), frame.Name)

	return formatter.TextFormatter.Format(entry)
}
