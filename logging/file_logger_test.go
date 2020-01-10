package logging

import (
	// "fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var (
	TestFolder      string = "./test"
	MessageContent  string = "Log messages to file from test environmentmance test"
	CleanTestFolder bool   = true
	MessageCount    int    = 100000
)

func setupTestCase(t *testing.T) func(t *testing.T, fileName string, message string) {
	createFolder(TestFolder)

	return func(t *testing.T, fileName string, message string) {
		if CleanTestFolder {
			configuration := NewLoggerConfiguration([]string{Time, Level, Name, Indent, Message})
			logger := NewConsoleLogger("Test_Env").Configure(configuration)

			dirRead, _ := os.Open(TestFolder)
			files, _ := dirRead.Readdir(0)

			for index := range files {
				file := files[index]
				name := file.Name()
				if strings.HasPrefix(name, fileName) {
					err := os.Remove(filepath.Join(TestFolder, name))

					logger.Warn("Delete file: ", name, " Error: ", err)
				}
			}

			logger.Warn("Delete test folder result: ", message)
		}
	}
}

func TestFileLoggerDebugResetIndent(t *testing.T) {
	configuration := NewLoggerConfiguration([]string{Time, Level, Name, Indent, Message})
	configuration.FileName = "./test/logging_test.txt"
	exception := NewException("Argument null exception")
	logger := NewFileLogger("Test.txt").Configure(configuration)

	teardown := setupTestCase(t)
	defer teardown(t, "logging_test", "File Logger")

	logger.StartGroup("Start Indent").StartGroup("Indent 2")

	for i := 0; i < MessageCount; i++ {
		logger.Info(MessageContent, i, i*2, true, false, exception)
	}
	logger.EndGroup()

	for i := 0; i < 10; i++ {
		logger.Debug(MessageContent, i, i*2, true, false, exception)
	}
	logger.EndGroup().Close()

}

func TestFileLoggerDebugSlowLogging(t *testing.T) {
	configuration := NewLoggerConfiguration([]string{Time, Level, Name, Indent, Message})
	configuration.FileName = "./test/logging_slow.txt"

	exception := NewException("Argument null exception")
	logger := NewFileLogger("SlowMessageGenerator").Configure(configuration)

	teardown := setupTestCase(t)
	defer teardown(t, "logging_slow", "File logger with slow messages")

	logger.StartGroup("Start Indent").StartGroup("Indent 2")

	for i := 0; i < 95; i++ {
		logger.Info(MessageContent, i, i*2, true, false, exception)
	}
	logger.EndGroup().Close()
}
