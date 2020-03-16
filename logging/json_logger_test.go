package matrix

import (
	"testing"
)

// TestJsonFileLoggerDebug test
func TestJsonFileLoggerDebug(t *testing.T) {
	configuration := NewLoggerConfiguration([]string{Message})
	configuration.FileName = "./test/logging_test.json"
	configuration.FileSize = OneMega * 5

	exception := NewException("Argument null exception")
	logger := NewJSONFileLogger("JsonLogger").Configure(configuration)

	teardown := setupTestCase(t)
	defer teardown(t, "logging_test", "Delte json test")

	logger.StartGroup("Start Indent").StartGroup("Indent 2")

	for i := 0; i < 500; i++ {
		logger.Info(MessageContent, i, i*2, true, false, exception)
	}
	logger.EndGroup().EndGroup().Close()
}

//
func TestJsonFileLoggerLevel(t *testing.T) {
	configuration := NewLoggerConfiguration([]string{Message})
	configuration.FileName = "./test/logging_warn.json"
	configuration.FileSize = OneMega * 5
	configuration.MinLevel = LevelWarn.Value
	exception := NewException("Argument null exception")
	logger := NewJSONFileLogger("JsonLogger").Configure(configuration)

	teardown := setupTestCase(t)
	defer teardown(t, "logging_warn", "Delete Json files for level test")

	logger.StartGroup("Start Indent").StartGroup("Indent 2")

	for i := 0; i < 500; i++ {
		logger.Info(MessageContent, i, i*2, true, false, exception)
	}

	for i := 0; i < 50; i++ {
		logger.Warn(MessageContent, i, i*2, true, false, exception)
	}

	for i := 0; i < 50; i++ {
		logger.Error(MessageContent, i, i*2, true, false, exception)
	}

	for i := 0; i < 50; i++ {
		logger.Fatal(MessageContent, i, i*2, true, false, exception)
	}

	logger.EndGroup().EndGroup().Close()
}
