package formatter
import (
	"github.com/Sirupsen/logrus"
	"github.com/facebookgo/stack"
)
// ~ type DefaultFormatter ---------------------------------------------------------------------------------------------
var textFormatter = &logrus.TextFormatter{
	TimestampFormat: "2006-01-02 15:04:05",
	FullTimestamp:   true,
}
type DefaultFormatter struct {
	CallerLevel *int
}
func (this *DefaultFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	frame := stack.Caller(4)
	if this.CallerLevel != nil {
		frame = stack.Caller(*this.CallerLevel)
	}
	fields := logrus.Fields{
		"file": frame.File,
		"func": frame.Name,
		"line": frame.Line,
	}
	for k, v := range fields {
		entry.Data[k] = v
	}
	return textFormatter.Format(entry)
}