package logging

import (
	"fmt"
	"os"

	"golang.org/x/sys/windows"
)

// ConsoleLogger log to console
type ConsoleLogger struct {
	Name             string
	Formatter        Formatter
	indentLevel      int
	layoutNames      []string
	layoutRepository LayoutRepository
	configuration    *LoggerConfiguration
	levelStyles      map[string][]string
}

// NewConsoleLogger create new console logger
func NewConsoleLogger(name string) ILogger {
	logger := &ConsoleLogger{
		Name:             name,
		Formatter:        Formatter{},
		indentLevel:      0,
		layoutNames:      []string{},
		layoutRepository: NewLayoutRepository(),
	}

	/*
	  set console mode to enable virtual terminal processing,
	  otherwise it may not work on some windows
	*/
	stdout := windows.Handle(os.Stdout.Fd())
	var originalMode uint32

	windows.GetConsoleMode(stdout, &originalMode)
	windows.SetConsoleMode(stdout, originalMode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)

	return logger
}

// Configure configure logger
func (logger *ConsoleLogger) Configure(config *LoggerConfiguration) ILogger {
	logger.layoutNames = config.LayoutNames
	logger.configuration = config
	return logger
}

// GetConfiguration get configuration
func (logger *ConsoleLogger) GetConfiguration() *LoggerConfiguration {
	return logger.configuration
}

// StartGroup start a group
func (logger *ConsoleLogger) StartGroup(name string) ILogger {
	logger.indentLevel++
	return logger
}

// EndGroup end a group
func (logger *ConsoleLogger) EndGroup() ILogger {
	logger.indentLevel--
	if logger.indentLevel < 0 {
		logger.indentLevel = 0
	}

	return logger
}

// ResetGroup reset
func (logger *ConsoleLogger) ResetGroup() ILogger {
	logger.indentLevel = 0
	return logger
}

// IsEnable is given level enabled
func (logger *ConsoleLogger) IsEnable(level LogLevel) bool {
	return level.Value >= logger.configuration.MinLevel
}

// WriteMessage writer message and data to string
func (logger *ConsoleLogger) WriteMessage(level string, message string, objects ...interface{}) string {
	layout := logger.layoutRepository.BuildLayout(logger.layoutNames...)

	logMessage := NewMessage(logger.Name, level, logger.indentLevel, message, objects...)
	content := layout.String(logMessage)

	return content
}

// Debug write at debug level
func (logger *ConsoleLogger) Debug(message string, objects ...interface{}) ILogger {

	content := logger.WriteMessage(LevelDebug.Name, message, objects...)

	logger.printMessage(content, LevelDebug)

	return logger
}

// Info write at info level
func (logger *ConsoleLogger) Info(message string, objects ...interface{}) ILogger {

	content := logger.WriteMessage(LevelInfo.Name, message, objects...)

	logger.printMessage(content, LevelInfo)

	return logger
}

// Warn write at warn level
func (logger *ConsoleLogger) Warn(message string, objects ...interface{}) ILogger {

	content := logger.WriteMessage(LevelWarn.Name, message, objects...)

	logger.printMessage(content, LevelWarn)

	return logger
}

// Error write at error level
func (logger *ConsoleLogger) Error(message string, objects ...interface{}) ILogger {

	content := logger.WriteMessage(LevelError.Name, message, objects...)

	logger.printMessage(content, LevelError)

	return logger
}

// Fatal write at fatal level
func (logger *ConsoleLogger) Fatal(message string, objects ...interface{}) ILogger {

	content := logger.WriteMessage(LevelFatal.Name, message, objects...)

	logger.printMessage(content, LevelFatal)

	return logger
}

// Close close logger
func (logger *ConsoleLogger) Close() ILogger {
	return logger
}
func (logger *ConsoleLogger) printMessage(message string, level LogLevel) {
	if logger.configuration.UseColor {
		style := logger.getLevelStyle(level)
		formatter := NewFormatter()
		fmt.Println(formatter.FormatConsoleOutput(message, style...))
	} else {
		fmt.Println(message)
	}
}

func (logger *ConsoleLogger) getLevelStyle(level LogLevel) []string {

	if logger.levelStyles == nil {
		logger.levelStyles = make(map[string][]string)
		logger.levelStyles[LevelDebug.Name] = logger.configuration.DebugStyle.parseLevelStyles()
		logger.levelStyles[LevelInfo.Name] = logger.configuration.InfoStyle.parseLevelStyles()
		logger.levelStyles[LevelWarn.Name] = logger.configuration.WarnStyle.parseLevelStyles()
		logger.levelStyles[LevelError.Name] = logger.configuration.ErrorStyle.parseLevelStyles()
		logger.levelStyles[LevelFatal.Name] = logger.configuration.FatalStyle.parseLevelStyles()
	}

	return logger.levelStyles[level.Name]
}
