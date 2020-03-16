package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseLevelStyles(t *testing.T) {
	style := LogLevelStyle{ColorDefaultText.Name, ColorRed.Value, "1 , 4 "}

	actual := style.parseLevelStyles()

	assert.Equal(t, 4, len(actual), "Style length should 4")
	assert.Equal(t, "39", actual[0], "Style length should be 4")
	assert.Equal(t, "48;5;31", actual[1], "Style length should be 4")
	assert.Equal(t, "1", actual[2], "Style length should be 1")
	assert.Equal(t, "4", actual[3], "Style length should be 4")
}
