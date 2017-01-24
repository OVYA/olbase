package log

import "testing"

func TestConsoleLogger(t *testing.T) {
	log := NewConsoleLogger()
	log.Info("I'll be logged with common and other field")
	t.Skip("ConsoleLogger.info does not panic")
}
