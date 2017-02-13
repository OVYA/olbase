package logruscaller

import (
	"fmt"
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/facebookgo/stack"
)

// CallerHook is the hook struct. Read the doc.
type CallerHook struct {
	Level int
}

// Fire takes the entry that the hook is fired for. `entry.Data[]` contains
// the fields for the entry. See the Fields section of the README.
func (hook *CallerHook) Fire(entry *logrus.Entry) error {
	fields := hook.caller()
	entry.Data["caller"] = fmt.Sprintf("%s:%s (%s)", fields["file"], fields["line"], fields["func"])

	// for k, v := range fields {
	// 	entry.Data[k] = v
	// }

	return nil
}

// Levels returns a slice of `Levels` the hook is fired for.
func (hook *CallerHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	}
}

func (hook *CallerHook) caller() map[string]string {
	var level = 6

	if hook.Level != 0 {
		level = hook.Level
	}

	frame := stack.Caller(level)

	fields := map[string]string{
		"file": frame.File,
		"func": frame.Name,
		"line": strconv.Itoa(frame.Line),
	}

	return fields
	// if _, file, line, ok := runtime.Caller(6); ok {
	// 	return strings.Join([]string{filepath.Base(file), strconv.Itoa(line)}, ":")
	// }
	// not sure what the convention should be here
	// return ""
}
