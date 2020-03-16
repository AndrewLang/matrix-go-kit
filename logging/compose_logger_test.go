package matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func createComposeLogger() ComposeLogger {

	logger := ComposeLogger{
		Loggers: make([]ILogger, 0),
	}

	return logger
}

func createConfiguredComposeLogger() ComposeLogger {
	config := defaultLogginConfig()
	logger := ComposeLogger{
		Loggers: make([]ILogger, 0),
	}

	logger.AddLogger(NewConsoleLogger("Compose Console"))

	logger.Configure(config)

	return logger
}

func TestNewComposeLogger(t *testing.T) {
	logger := NewComposeLogger("Testing")

	assert.NotNil(t, logger, "Logger should not be nil")
}

func TestComposeConfigure(t *testing.T) {
	config := defaultLogginConfig()
	logger := createComposeLogger()
	logger.Configure(config)
}

func TestComposeGetConfiguration(t *testing.T) {
	logger := createConfiguredComposeLogger()
	config := logger.GetConfiguration()

	assert.NotNil(t, config)
}

func TestComposeStartGroup(t *testing.T) {
	logger := createConfiguredComposeLogger()
	logger.StartGroup("Test")

}

func TestComposeEndGroup(t *testing.T) {
	logger := createConfiguredComposeLogger()
	logger.EndGroup()
}

func TestComposeResetGroup(t *testing.T) {
	logger := createConfiguredComposeLogger()
	logger.ResetGroup()
}

func TestComposeIsEnabled(t *testing.T) {
	logger := createConfiguredComposeLogger()
	actual := logger.IsEnable(LevelInfo)

	assert.True(t, actual)
}

func TestComposeClose(t *testing.T) {
	logger := createConfiguredComposeLogger()
	logger.Close()
}

func TestDebugCompose(t *testing.T) {
	logger := createConfiguredComposeLogger()
	logger.Debug("Message for compose logger")
}

func TestInfoCompose(t *testing.T) {
	logger := createConfiguredComposeLogger()
	logger.Info("Message for compose logger")
}

func TestWarnCompose(t *testing.T) {
	logger := createConfiguredComposeLogger()
	logger.Warn("Message for compose logger")
}

func TestErrorCompose(t *testing.T) {
	logger := createConfiguredComposeLogger()
	logger.Error("Message for compose logger")
}

func TestFatalCompose(t *testing.T) {
	logger := createConfiguredComposeLogger()
	logger.Fatal("Message for compose logger")
}
