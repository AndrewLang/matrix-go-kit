package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJoin(t *testing.T) {
	actual := Join("hello ", "world")

	assert.Equal(t, actual, "hello world", "Join result should be 'hello world'")
}

func TestJoinWith(t *testing.T) {
	actual := JoinWith(" ", "hello", "world")
	assert.Equal(t, actual, "hello world", "Join result should be 'hello world'")

	actual = JoinWith(",", "90")
	assert.Equal(t, "90", actual, "Join result should be '90'")

	array := []string{"90"}
	actual = JoinWith(",", array...)
	assert.Equal(t, "90", actual, "Join result should be '90'")

	actual = JoinWith(",", "90", "80")
	assert.Equal(t, "90,80", actual, "Join result should be '90,80'")
}

func TestJoinStrings(t *testing.T) {
	actual := JoinStrings([]string{"hello", "world"}, " ")
	assert.Equal(t, actual, "hello world", "Join result should be 'hello world'")

	actual = JoinStrings([]string{"90"}, ",")
	assert.Equal(t, "90", actual, "Join result should be '90'")

	actual = JoinStrings([]string{"90", "80"}, ",")
	assert.Equal(t, "90,80", actual, "Join result should be '90,80'")
}

func TestIsNumber(t *testing.T) {
	assert.Equal(t, IsNumber("123"), true, "123 is number")
	assert.Equal(t, IsNumber("123b"), false, "123b is not number")
}

func TestToNumber(t *testing.T) {
	value, err := ToNumber("123")
	assert.Equal(t, value, 123, "123 is number")

	value, err = ToNumber("123b")
	assert.NotNil(t, err, "123b is not number")
}
