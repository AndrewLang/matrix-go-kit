package matrix

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

// LogJSONMessage reresent a json message
type LogJSONMessage struct {
	Datetime string `json:"datetime"`
	Name     string `json:"name"`
	Level    string `json:"level"`
	Message  string `json:"message"`
}

// JSONFileLogger json logger
type JSONFileLogger struct {
	Name             string
	Formatter        Formatter
	layoutNames      []string
	fileName         string
	indentLevel      int
	fileSize         int64
	layoutRepository LayoutRepository
	mutex            sync.Mutex
	file             *os.File
	timer            *time.Timer
	messages         []*LogJSONMessage
	configuration    *LoggerConfiguration
}

// NewJSONFileLogger create new file logger
func NewJSONFileLogger(name string) ILogger {
	logger := &JSONFileLogger{
		Name:             name,
		Formatter:        Formatter{},
		layoutNames:      []string{},
		fileName:         "",
		indentLevel:      0,
		fileSize:         DefaultLogFileSize,
		layoutRepository: NewLayoutRepository(),
		messages:         make([]*LogJSONMessage, 0),
	}
	runtime.SetFinalizer(logger, closeJSONLogger)

	logger.timer = time.AfterFunc(3*time.Second, func() {
		fmt.Println("Flush message to file by timer")
		logger.writeFile()
	})

	return logger
}

// Configure configure logger
func (logger *JSONFileLogger) Configure(config *LoggerConfiguration) ILogger {
	logger.layoutNames = config.LayoutNames
	logger.fileName = config.FileName
	logger.fileSize = config.FileSize
	logger.configuration = config
	return logger
}

// GetConfiguration get configuration
func (logger *JSONFileLogger) GetConfiguration() *LoggerConfiguration {
	return logger.configuration
}

// StartGroup start a group
func (logger *JSONFileLogger) StartGroup(name string) ILogger {
	logger.indentLevel++
	return logger
}

// EndGroup end a group
func (logger *JSONFileLogger) EndGroup() ILogger {
	logger.indentLevel--
	if logger.indentLevel < 0 {
		logger.indentLevel = 0
	}
	return logger
}

// ResetGroup reset
func (logger *JSONFileLogger) ResetGroup() ILogger {
	logger.indentLevel = 0
	return logger
}

// IsEnable is given level enabled
func (logger *JSONFileLogger) IsEnable(level LogLevel) bool {
	return level.Value >= logger.configuration.MinLevel
}

// WriteMessage writer message and data to string
func (logger *JSONFileLogger) WriteMessage(level string, message string, objects ...interface{}) *LogJSONMessage {
	layout := logger.layoutRepository.BuildLayout(logger.layoutNames...)

	logMessage := NewMessage(logger.Name, level, logger.indentLevel, message, objects...)
	jsonMessage := &LogJSONMessage{
		Datetime: logger.Formatter.FormatNow(),
		Name:     logger.Name,
		Level:    level,
		Message:  layout.String(logMessage),
	}

	return jsonMessage
}

// Debug write at debug level
func (logger *JSONFileLogger) Debug(message string, objects ...interface{}) ILogger {

	if logger.IsEnable(LevelDebug) {
		content := logger.WriteMessage(LevelDebug.Name, message, objects...)

		logger.printMessage(content)
	}

	return logger
}

// Info write at info level
func (logger *JSONFileLogger) Info(message string, objects ...interface{}) ILogger {

	if logger.IsEnable(LevelInfo) {
		content := logger.WriteMessage(LevelInfo.Name, message, objects...)

		logger.printMessage(content)
	}

	return logger
}

// Warn write at warn level
func (logger *JSONFileLogger) Warn(message string, objects ...interface{}) ILogger {

	if logger.IsEnable(LevelWarn) {
		content := logger.WriteMessage(LevelWarn.Name, message, objects...)

		logger.printMessage(content)
	}

	return logger
}

// Error write at error level
func (logger *JSONFileLogger) Error(message string, objects ...interface{}) ILogger {

	if logger.IsEnable(LevelError) {
		content := logger.WriteMessage(LevelError.Name, message, objects...)

		logger.printMessage(content)
	}

	return logger
}

// Fatal write at fatal level
func (logger *JSONFileLogger) Fatal(message string, objects ...interface{}) ILogger {

	if logger.IsEnable(LevelFatal) {
		content := logger.WriteMessage(LevelFatal.Name, message, objects...)

		logger.printMessage(content)
	}

	return logger
}

func (logger *JSONFileLogger) writeFile() {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	size, sizeErr := getFileSize(logger.fileName)
	if sizeErr != nil {
		return
	}
	if size >= logger.fileSize {
		logger.file = nil
		logger.fileName = generateFileName(logger.configuration.FileName)
		logger.initialize()
	}

	content, err := json.MarshalIndent(logger.messages, "", "\t")
	if err != nil {
		fmt.Println("Convert messages to json error", err)
		return
	}

	_, err = logger.file.Write(content)
	if err != nil {
		fmt.Println("Write to file error: ", err)
		return
	}
}

func (logger *JSONFileLogger) initialize() {
	if logger.file == nil {
		file, err := os.OpenFile(logger.fileName, os.O_CREATE|os.O_WRONLY, 0644)
		if err == nil {
			logger.file = file
		} else {
			fmt.Println("Create file writer error: ", err, logger.fileName)
		}
	}
}

// Close close logger
func (logger *JSONFileLogger) Close() ILogger {
	closeJSONLogger(logger)
	return logger
}

// closeJSONLogger close logger resources
func closeJSONLogger(logger *JSONFileLogger) {
	logger.writeFile()

	if logger.file != nil {
		logger.file.Sync()
		logger.file.Close()
		logger.file = nil
	}

	logger.timer.Stop()
	logger.messages = nil
	fmt.Println("Json File logger closed")
}

func (logger *JSONFileLogger) printMessage(message *LogJSONMessage) {

	logger.initialize()

	logger.messages = append(logger.messages, message)

	if len(logger.messages) > 100 {
		logger.writeFile()
	}
}
