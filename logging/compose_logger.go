package matrix

// ComposeLogger composed logger for internal use
type ComposeLogger struct {
	Loggers       []ILogger
	configuration *LoggerConfiguration
	name          string
}

// NewComposeLogger create new compose logger
func NewComposeLogger(name string) ILogger {
	logger := &ComposeLogger{
		Loggers: make([]ILogger, 0),
		name:    name,
	}
	return logger
}

// CreateComposeLogger create new compose logger
func CreateComposeLogger(name string) *ComposeLogger {
	logger := &ComposeLogger{
		Loggers: make([]ILogger, 0),
		name:    name,
	}
	return logger
}

// AddLogger add a logger
func (logger *ComposeLogger) AddLogger(instance ILogger) *ComposeLogger {

	logger.Loggers = append(logger.Loggers, instance)
	return logger
}

// Configure configure logger
func (logger *ComposeLogger) Configure(config *LoggerConfiguration) ILogger {
	logger.configuration = config
	for _, item := range logger.Loggers {
		item.Configure(config)
	}
	return logger
}

// GetConfiguration get configuration
func (logger *ComposeLogger) GetConfiguration() *LoggerConfiguration {
	return logger.configuration
}

// StartGroup start a new group
func (logger *ComposeLogger) StartGroup(name string) ILogger {
	for _, item := range logger.Loggers {
		item.StartGroup(name)
	}
	return logger
}

// EndGroup end a group
func (logger *ComposeLogger) EndGroup() ILogger {
	for _, item := range logger.Loggers {
		item.EndGroup()
	}
	return logger
}

// ResetGroup reset groups
func (logger *ComposeLogger) ResetGroup() ILogger {
	for _, item := range logger.Loggers {
		item.ResetGroup()
	}
	return logger
}

// IsEnable check whether level is enabled
func (logger *ComposeLogger) IsEnable(level LogLevel) bool {
	return true
}

// Close close logger
func (logger *ComposeLogger) Close() ILogger {
	for _, item := range logger.Loggers {
		item.Close()
	}
	return logger
}

// Debug write message to debug level
func (logger *ComposeLogger) Debug(message string, objects ...interface{}) ILogger {
	for _, item := range logger.Loggers {
		item.Debug(message, objects...)
	}
	return logger
}

// Info write message to info level
func (logger *ComposeLogger) Info(message string, objects ...interface{}) ILogger {
	for _, item := range logger.Loggers {
		item.Info(message, objects...)
	}
	return logger
}

// Warn write message to warn level
func (logger *ComposeLogger) Warn(message string, objects ...interface{}) ILogger {
	for _, item := range logger.Loggers {
		item.Warn(message, objects...)
	}
	return logger
}

// Error write message to error level
func (logger *ComposeLogger) Error(message string, objects ...interface{}) ILogger {
	for _, item := range logger.Loggers {
		item.Error(message, objects...)
	}
	return logger
}

// Fatal write message to fatal level
func (logger *ComposeLogger) Fatal(message string, objects ...interface{}) ILogger {
	for _, item := range logger.Loggers {
		item.Fatal(message, objects...)
	}
	return logger
}
