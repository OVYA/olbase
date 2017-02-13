package formatter

import (
	"testing"

	"github.com/Sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestTextFormatter(t *testing.T) {
	formatter := TextFormatter{
		TextFormatter: logrus.TextFormatter{
			TimestampFormat: "2016-01-02 15:04:05",
			FullTimestamp:   true,
			DisableColors:   true,
		},
		CallerLevel: 1,
	}

	bytes, errChk := formatter.Format(logrus.WithField("test_field", "test"))
	assert.NoError(t, errChk)

	str := string(bytes)
	assert.Contains(t, str, "test_field=test")
	assert.Contains(t, str, `caller="ovya/lib/olbase/log/formatter/text_formater_test.go:20 (TestTextFormatter)"`)
}
