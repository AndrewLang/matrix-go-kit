package app

import (
	logging "github.com/andrewlang/matrix-go-logging"
)

func createTestLogger(name string) logging.ILogger {
	configuration := logging.NewLoggerConfiguration([]string{logging.Time, logging.Level, logging.Name, logging.Indent, logging.Message})
	logger := logging.NewConsoleLogger(name).Configure(configuration)
	return logger
}
