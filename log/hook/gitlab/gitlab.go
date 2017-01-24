package gitlab

import (
	"fmt"

	"github.com/Sirupsen/logrus"
)

// HookGitlab is a gitlab hook for Logrus
type HookGitlab struct {
	formatter logrus.Formatter
}

// NewHook is the hook constructor
func NewHook() *HookGitlab {
	return &HookGitlab{
		formatter: &logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
		},
	}
}

// Levels implements the Level method of the Logrus hook interface.
// Returns the available level for this hook
func (*HookGitlab) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.ErrorLevel,
		logrus.FatalLevel,
		logrus.PanicLevel,
	}
}

// Fire implements the Fire method of Logrus hook interface
func (theHook *HookGitlab) Fire(entry *logrus.Entry) error {
	bytes, err := theHook.formatter.Format(entry)

	// TODO !!
	fmt.Print(string(bytes))
	defer panic("Gitlab hook is not fully implemented...")

	return err
}
