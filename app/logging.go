package app

import (
	io "github.com/andrewlang/matrix-go-kit/io"
	log "github.com/andrewlang/matrix-go-kit/log"
)

// CreateLogger create logger with given name
func CreateLogger(name string) log.ILogger {
	/*
		If there is log configuration then load it from file, otherwise use default log configuration
	*/
	factory := log.NewLoggerFactory()
	configFile := io.NewFile(LoggingConfigFile)

	if configFile.Exists() {
		factory.ConfigureFromFile(LoggingConfigFile)
	} else {
		config := log.NewLogTargetConfigurations()

		consoleTarget := log.NewLogTargetConfiguration("Console",
			log.ConsoleLoggerName,
			[]string{log.Time, log.Level, log.Name, log.Indent, log.Message})
		config.AddTarget(consoleTarget)

		// do not write to log file for now
		// fileTarget := log.NewLogTargetConfiguration("File", log.FileLoggerName, []string{log.Time, log.Level, log.Name, log.Indent, log.Message})
		// fileTarget.Configuration.FileName = "log.txt"
		// config.AddTarget(fileTarget)

		factory.Configure(config)
	}

	logger, _ := factory.Create(name)
	return logger
}
