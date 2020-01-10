package logging

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToString(t *testing.T) {
	formatter := Formatter{}
	err := NewException("this is an exception")

	actual := formatter.ToString("this is test", 1, 2, true, err)
	fmt.Println(actual)

	// assert.Equal(t, actual, "hello world", "Join result should be 'hello world'")
}

func TestFormatConsoleStyle(t *testing.T) {
	formatter := NewFormatter()

	actual := formatter.FormatConsoleStyle("1", "4", "31")
	assert.Equal(t, "\033[1;4;31m", actual, `Console style should be \033[1;4;31m `)

	actual = formatter.FormatConsoleStyle("90")
	assert.Equal(t, "\033[90m", actual, "Console style should be \033[90;m")
}
