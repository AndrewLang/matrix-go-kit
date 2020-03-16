package logging

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateFileName(t *testing.T) {
	fileName := "1000-batch.txt"
	configuration := NewLoggerConfiguration([]string{Time, Level, Name, Indent, Message})
	configuration.FileName = fileName

	logger := NewFileLogger("SlowMessageGenerator").Configure(configuration)

	actual := generateFileName(logger.GetConfiguration().FileName)

	assert.Equal(t, "1000-batch_1.txt", actual, "New log file name is wrong")
}
