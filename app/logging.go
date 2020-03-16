package matrix

import (
	io "github.com/andrewlang/matrix-go-kit/io"
	logging "github.com/andrewlang/matrix-go-kit/logging"
)

// CreateLogger create logger with given name
func CreateLogger(name string) logging.ILogger {
	/*
		If there is logging configuration then load it from file, otherwise use default logging configuration
	*/
	factory := logging.NewLoggerFactory()
	configFile := io.NewFile(LoggingConfigFile)

	if configFile.Exists() {
		factory.ConfigureFromFile(LoggingConfigFile)
	} else {
		config := logging.NewLogTargetConfigurations()

		consoleTarget := logging.NewLogTargetConfiguration("Console",
			logging.ConsoleLoggerName,
			[]string{logging.Time, logging.Level, logging.Name, logging.Indent, logging.Message})
		config.AddTarget(consoleTarget)

		// do not write to log file for now
		// fileTarget := logging.NewLogTargetConfiguration("File", logging.FileLoggerName, []string{logging.Time, logging.Level, logging.Name, logging.Indent, logging.Message})
		// fileTarget.Configuration.FileName = "log.txt"
		// config.AddTarget(fileTarget)

		factory.Configure(config)
	}

	logger, _ := factory.Create(name)
	return logger
}
