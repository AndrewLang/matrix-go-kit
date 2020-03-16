package app

import (
	log "github.com/andrewlang/matrix-go-kit/log"
)

func createTestLogger(name string) log.ILogger {
	configuration := log.NewLoggerConfiguration([]string{log.Time, log.Level, log.Name, log.Indent, log.Message})
	logger := log.NewConsoleLogger(name).Configure(configuration)
	return logger
}
