package app

import (
	"testing"

	"github.com/andrewlang/matrix-go-kit/io"
	"github.com/stretchr/testify/assert"
)

func TestNewConfiguration(t *testing.T) {
	config := NewConfiguration()

	assert.NotNil(t, config)
}

func TestToJSON(t *testing.T) {
	config := NewConfiguration()
	content := config.ToJSON()

	assert.True(t, len(content) > 0)

	logger := createTestLogger("Configuration")
	logger.Info("Configuration to JSON").Info(content)
}
func TestFromJSON(t *testing.T) {
	config := NewConfiguration()

}

func TestToFromFile(t *testing.T) {
	file := io.NewFile("default.config")
	config := NewConfiguration()

	config.ToFile(file.Path())

	actual := NewConfiguration()
	actual.FromFile(file.Path())

	file.Delete()
}
