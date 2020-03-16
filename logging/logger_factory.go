package matrix

// LoggerCreator create logger instance
type LoggerCreator func(name string) ILogger

// LoggerFactory logger factory
type LoggerFactory struct {
	Configuration *LogTargetConfigurations
	creators      map[string]LoggerCreator
}

// NewLoggerFactory create new logger factory
func NewLoggerFactory() *LoggerFactory {
	factory := &LoggerFactory{}

	factory.creators = make(map[string]LoggerCreator)

	factory.RegisterCreator(ConsoleLoggerName, func(name string) ILogger {
		return NewConsoleLogger(name)
	}).RegisterCreator(FileLoggerName, func(name string) ILogger {
		return NewFileLogger(name)
	}).RegisterCreator(JSONFileLoggerName, func(name string) ILogger {
		return NewJSONFileLogger(name)
	}).RegisterCreator(ComposeLoggerName, func(name string) ILogger {
		return NewComposeLogger(name)
	})

	return factory
}

// RegisterCreator register a creator to factory
func (factory *LoggerFactory) RegisterCreator(name string, creator LoggerCreator) *LoggerFactory {
	factory.creators[name] = creator
	return factory
}

// Configure Configure factory
func (factory *LoggerFactory) Configure(config *LogTargetConfigurations) *LoggerFactory {
	factory.Configuration = config
	return factory
}

// ConfigureFromFile configure from config file
func (factory *LoggerFactory) ConfigureFromFile(file string) *LoggerFactory {
	if fileExists(file) {
		content := readAllText(file)
		config := NewLogTargetConfigurations()
		config.FromJSON(content)
		factory.Configure(config)
	}
	return factory
}

// Create create a logger instance
func (factory *LoggerFactory) Create(name string) (ILogger, error) {
	if factory.Configuration == nil {
		return nil, NewException("Factory configuration is not configured. ")
	}

	composeLogger := CreateComposeLogger(name)

	for _, target := range factory.Configuration.Targets {
		creator, ok := factory.creators[target.Type]
		if !ok {
			continue
		}

		logger := creator(name)
		logger.Configure(target.Configuration)
		composeLogger.AddLogger(logger)
	}

	return composeLogger, nil
}
