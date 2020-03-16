package logging

import (
	"fmt"
	"testing"
)

func TestNewException(t *testing.T) {
	exception := NewException("test")

	if exception.Message != "test" {
		t.Errorf("No message of the exception")
	}
}

func TestError(t *testing.T) {
	exception := NewException("test")
	err := exception.Error()

	fmt.Println(err)
}
