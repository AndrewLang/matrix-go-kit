package logging

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLoggerFactory(t *testing.T) {
	factory := NewLoggerFactory()
	length := len(factory.creators)

	assert.Equal(t, 4, length, "There should be 4 creators")
}

func TestConfigureLoggerFactory(t *testing.T) {
	config := defaultLogTargetConfigs()

	factory := NewLoggerFactory().Configure(config)

	configuration := factory.Configuration.GetTarget("Test")

	assert.NotNil(t, configuration)

	assert.Equal(t, "", configuration.Configuration.FileName, "")
	assert.Equal(t, 0, configuration.Configuration.MinLevel, "")
	assert.Equal(t, int64(2097152), configuration.Configuration.FileSize, "")
	assert.Equal(t, true, configuration.Configuration.UseColor, "")
	assert.Equal(t, "Time", configuration.Configuration.LayoutNames[0], "")
	assert.Equal(t, "Level", configuration.Configuration.LayoutNames[1], "")
	assert.Equal(t, "Name", configuration.Configuration.LayoutNames[2], "")
	assert.Equal(t, "Indent", configuration.Configuration.LayoutNames[3], "")
	assert.Equal(t, "Message", configuration.Configuration.LayoutNames[4], "")
}

func TestCreateLogger(t *testing.T) {
	configuration := defaultLogTargetConfigs()
	factory := NewLoggerFactory().Configure(configuration)
	logger, err := factory.Create(ConsoleLoggerName)

	assert.Nil(t, err, "Error should be nil")
	assert.NotNil(t, logger, "Logger should not be nil")

	content := "Logger created from factory"
	logger.Debug("Logger configuration: ", logger.GetConfiguration()).
		Debug(content).
		Info(content).
		Warn(content).
		Error(content).
		Fatal(content)
}

func TestCreateComposeLoggerFromConfig(t *testing.T) {
	configuration := composeLogTargetConfigs()
	factory := NewLoggerFactory().Configure(configuration)
	logger, err := factory.Create(ConsoleLoggerName)

	teardown := setupTestCase(t)
	defer teardown(t, "compose_logger", "Factory Logger")

	assert.Nil(t, err, "Error should be nil")
	assert.NotNil(t, logger, "Logger should not be nil")

	for i := 0; i < 50; i++ {
		content := "Logger created from factory for output multiple target " + strconv.Itoa(i)
		logger.Debug(content).
			Info(content).
			Warn(content).
			Error(content).
			Fatal(content)
	}

	logger.Close()
}

func TestCreateLoggerNotFound(t *testing.T) {
	configuration := defaultLogTargetConfigs()
	factory := NewLoggerFactory().Configure(configuration)
	logger, err := factory.Create("Mock")

	assert.NotNil(t, logger, "Error should not be nil")
	assert.Nil(t, err)
}

func TestCreateLoggerWithoutConfig(t *testing.T) {
	factory := NewLoggerFactory()
	_, err := factory.Create(ConsoleLoggerName)

	assert.NotNil(t, err, "Error should not be nil")
}

func TestRegisterCreator(t *testing.T) {

	factory := NewLoggerFactory()
	factory.RegisterCreator("Mock", func(name string) ILogger {
		return NewConsoleLogger(name)
	})
	length := len(factory.creators)
	assert.Equal(t, 5, length, "There should be 5 creators")
}

func TestConfigureFromFile(t *testing.T) {
	teardown := setupTestCase(t)
	defer teardown(t, "logging.config", "Factory configuration")

	configuration := defaultLogTargetConfigs()
	content := configuration.ToJSON()

	file := "./test/logging.config.json"
	err := writeToFile(file, content)

	assert.Nil(t, err)

	factory := NewLoggerFactory()
	factory.ConfigureFromFile(file)

	logger, err := factory.Create(ConsoleLoggerName)

	assert.NotNil(t, logger)
	assert.Nil(t, err)

	content = "Logger factory loaded configuration from file"
	logger.Debug(content).
		Info(content).
		Warn(content).
		Error(content).
		Fatal(content)

	logger.Close()
}
