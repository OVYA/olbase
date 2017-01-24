package formatter

import (
	"fmt"
	"strings"
	"testing"

	"github.com/Sirupsen/logrus"
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

	fmt.Println(str)

	if !strings.Contains(str, "test=test") {
		t.Error(str + " does not contain the string 'test=test'")
	}

	if errChk != nil {
		t.Error(errChk)
	}

}
