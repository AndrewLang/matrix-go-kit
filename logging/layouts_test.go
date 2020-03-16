package logging

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLogTimeLayout(t *testing.T) {
	layout := NewLogTimeLayout()

	fmt.Println("layout date", layout.date)
	assert.NotEmpty(t, layout.date, "layout date is empty")
}

func TestLogTimeLayoutValue(t *testing.T) {
	message := NewMessage("LoggingComponent", LevelDebug.Name, 0, "This is a test message ", []interface{}{1, 2, 3})
	layout := NewLogTimeLayout()
	now := layout.String(message)

	fmt.Println("layout value", now)
	assert.NotEmpty(t, now, "layout date is empty")
}

func TestLogLevelLayout(t *testing.T) {
	message := NewMessage("LoggingComponent", LevelDebug.Name, 0, "This is a test message ", []interface{}{1, 2, 3})

	layout := NewLogLevelLayout()
	actual := layout.String(message)

	fmt.Println("Log level value", actual)
	assert.NotEmpty(t, actual, "layout date is empty")
}

func TestLogNameLayout(t *testing.T) {
	message := NewMessage("LoggingComponent", LevelDebug.Name, 0, "This is a test message ", []interface{}{1, 2, 3})
	layout := NewLogNameLayout()
	actual := layout.String(message)

	fmt.Println("Log name value", actual)
	assert.NotEmpty(t, actual, "layout date is empty")
}

func TestLogIndentLayout(t *testing.T) {
	message := NewMessage("LoggingComponent", LevelDebug.Name, 1, "This is a test message ", []interface{}{1, 2, 3})
	layout := NewLogIndentLayout()
	actual := layout.String(message)

	fmt.Println("Log indent value", actual)
	assert.NotEmpty(t, actual, "layout date is empty")
}

func TestLogMessageLayout(t *testing.T) {
	message := NewMessage("LoggingComponent", LevelDebug.Name, 0, "This is a test message ", []interface{}{1, 2, 3})
	layout := NewLogMessageLayout()
	actual := layout.String(message)

	fmt.Println("Log message value", actual)
	assert.NotEmpty(t, actual, "layout date is empty")
}

func TestComposeLayout(t *testing.T) {

	message := NewMessage("LoggingComponent", LevelDebug.Name, 1, "This is a test message ", []interface{}{1, 2, 3}...)

	timeLayout := NewLogTimeLayout()
	levelLayout := NewLogLevelLayout()
	nameLayout := NewLogNameLayout()
	indentLayout := NewLogIndentLayout()
	messageLayout := NewLogMessageLayout()
	composeLayout := NewComposeLayout()

	composeLayout.AddLayouts(timeLayout, levelLayout, nameLayout, indentLayout, messageLayout)

	assert.Equal(t, 5, len(composeLayout.Layouts), "Layout count should be 5")

	actual := composeLayout.String(message)

	fmt.Println("Compose layout result: ", actual)

}

func TestLayoutRepository(t *testing.T) {
	repository := NewLayoutRepository()

	assert.NotEqual(t, repository.GetLayout("Name"), nil, "Name layout should not be nil")
	assert.NotEqual(t, repository.GetLayout("NameLayout"), nil, "Name layout should not be nil")
}

func TestBuildLayout(t *testing.T) {
	repository := NewLayoutRepository()

	layout := repository.BuildLayout(Time, Level, Name, Indent, Message)
	message := NewMessage("LoggingComponent", LevelDebug.Name, 1, "This is a test message ", 1, 2, 3)

	actual := layout.String(message)

	fmt.Println("Compose layout result with repository: ", actual)
}
