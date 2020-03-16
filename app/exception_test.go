package app

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewException(t *testing.T) {
	exception := NewException("test")

	assert.Equal(t, "test", exception.Message)
	assert.Nil(t, exception.InnerException)
}

func TestExceptionError(t *testing.T) {
	exception := NewException("test")
	err := exception.Error()

	fmt.Println(err)
}

func TestNewExceptionWithInner(t *testing.T) {
	innerErr := NewException("inner barsoom is too hot.")
	err := NewExceptionWithInner("Basoom not normal", innerErr)

	assert.NotNil(t, err.InnerException)
	assert.Equal(t, "Basoom not normal", err.Message)
}
