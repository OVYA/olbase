package formatter

import (
	"testing"

	"github.com/Sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestConsoleFormatter(t *testing.T) {
	formatter := ConsoleFormatter{
		logrus.TextFormatter{
			TimestampFormat: "2016-01-02 15:04:05",
			FullTimestamp:   true,
			DisableColors:   true,
		},
	}

	bytes, errChk := formatter.Format(logrus.WithField("test", "test"))
	str := string(bytes)

	assert.Contains(t, str, "test=test")
	assert.NoError(t, errChk)
}
